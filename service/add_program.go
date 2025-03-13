package service

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	schedulerexecutor "github.com/nightsilvertech/clockwerk/executors"
	schedulerhttpmethod "github.com/nightsilvertech/clockwerk/executors/http"
	pbScheduler "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-promo-auth/constants"
	programconst "gitlab.com/nbdgocean6/nobita-promo-program/constants"
	"gitlab.com/nbdgocean6/nobita-promo-program/gvars"

	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	pbwhitelist "gitlab.com/nbdgocean6/nobita-promo-whitelist/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-util/er"
	"gitlab.com/nbdgocean6/nobita-util/jwt"
	"gitlab.com/nbdgocean6/nobita-util/lgr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"time"
)

func CreateContext(ctx context.Context, token string) context.Context {
	md := metadata.MD{}
	md.Set("authorization", "Bearer "+token)
	return metadata.NewIncomingContext(ctx, md)
}

func (s service) bulkInsertCustCSV(ctx context.Context, req *pb.AddProgramReq, programID string) error {
	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return er.Ebl(codes.Internal, "failed to read jwt", err)
	}

	_, err = s.repo.Microservice.WhitelistSvc.AddCustomerWhitelistBulk(CreateContext(ctx, jwtToken.Bearer), &pbwhitelist.AddCustomerWhitelistBulkReq{
		ProgramId: programID,
		CreatedBy: req.CreatedBy,
		Url:       req.CustomerCsvUrl,
	})
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return err
	}
	return nil
}

func (s service) bulkInsertMerchCSV(ctx context.Context, req *pb.AddProgramReq, programID string) error {
	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return er.Ebl(codes.Internal, "failed to read jwt", err)
	}

	_, err = s.repo.Microservice.WhitelistSvc.AddMerchantWhitelistBulk(CreateContext(ctx, jwtToken.Bearer), &pbwhitelist.AddMerchantWhitelistBulkReq{
		ProgramId: programID,
		CreatedBy: req.CreatedBy,
		Url:       req.MerchantCsvUrl,
	})
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return err
	}
	return nil
}

func (s service) statusUpdateScheduler(ctx context.Context, req *pb.AddProgramReq, programID string) error {
	username := s.vault.Get("/scheduler_basic_auth:username")
	password := s.vault.Get("/scheduler_basic_auth:password")
	systemUsername := s.vault.Get("/basic_auth_system:username")
	systemPassword := s.vault.Get("/basic_auth_system:password")

	if req.Status == 0 {
		// start at update program to active
		startAtTime := time.Unix(req.StartAt/1000, 0)
		_, err := s.repo.Scheduler.AddScheduler(ctx, &pbScheduler.Scheduler{
			Username: username,
			Password: password,
			Name:     fmt.Sprintf("update program id %s status to active", programID),
			Url:      fmt.Sprintf(programconst.UpdateProgramToActiveURL+"%s/1", programID),
			Executor: schedulerexecutor.HTTP,
			Method:   schedulerhttpmethod.MethodPut,
			Disabled: false,
			Persist:  false,
			Spec:     fmt.Sprintf("0 0 %d %d *", startAtTime.Day(), int(startAtTime.Month())),
			Headers: []string{
				"Content-Type|application/json",
				fmt.Sprintf("Authorization|Basic %s:%s", systemUsername, systemPassword),
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for activate program", err)
		}

		// end at update program to inactive
		endAtTime := time.Unix(req.EndAt/1000, 0)
		_, err = s.repo.Scheduler.AddScheduler(ctx, &pbScheduler.Scheduler{
			Username: username,
			Password: password,
			Name:     fmt.Sprintf("update program id %s status to inactivate", programID),
			Url:      fmt.Sprintf(programconst.UpdateProgramToActiveURL+"%s/0", programID),
			Executor: schedulerexecutor.HTTP,
			Method:   schedulerhttpmethod.MethodPut,
			Disabled: false,
			Persist:  false,
			Spec:     fmt.Sprintf("0 0 %d %d *", endAtTime.Day(), int(endAtTime.Month())),
			Headers: []string{
				"Content-Type|application/json",
				fmt.Sprintf("Authorization|Basic %s:%s", systemUsername, systemPassword),
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for inactivate program", err)
		}
	} else {
		// end at update program to inactive
		endAtTime := time.Unix(req.EndAt/1000, 0)
		_, err := s.repo.Scheduler.AddScheduler(ctx, &pbScheduler.Scheduler{
			Username: username,
			Password: password,
			Name:     fmt.Sprintf("update program id %s status to inactivate", programID),
			Url:      fmt.Sprintf(programconst.UpdateProgramToActiveURL+"%s/0", programID),
			Executor: schedulerexecutor.HTTP,
			Method:   schedulerhttpmethod.MethodPut,
			Disabled: false,
			Persist:  false,
			Spec:     fmt.Sprintf("0 0 %d %d *", endAtTime.Day(), int(endAtTime.Month())),
			Headers: []string{
				"Content-Type|application/json",
				fmt.Sprintf("Authorization|Basic %s:%s", systemUsername, systemPassword),
			},
		})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return er.Ebl(codes.Internal, "failed to create scheduler for inactivate program", err)
		}
	}
	return nil
}

func (s service) AddProgram(ctx context.Context, req *pb.AddProgramReq) (res *pb.ProgramRes, err error) {
	const funcName = `AddProgram`
	const maxProgramPerPeriod = 3
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to generate jwt", err)
	}

	jwtData, err := jwt.
		NewJWT(jwtToken.Bearer, jwtToken.Secret).
		ExtractKeys([]string{constants.UserID})
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to extract jwt data", err)
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	currentUnix := now.Unix() * 1000
	if currentUnix >= req.StartAt && currentUnix <= req.EndAt {
		req.Status = 1
	} else {
		req.Status = 0
	}

	var response pb.ProgramRes
	req.CreatedBy = jwtData[constants.UserID]

	count, err := s.repo.ReadWriter.ReadProgramsByBetweenStartEndDate(ctx, req.TopicId, req.StartAt, req.EndAt)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.AlreadyExists, "failed to check running program", err)
	}
	if count >= maxProgramPerPeriod {
		return nil, er.Ebl(codes.ResourceExhausted, "Error add program: has reach maximum running concurrently", nil)
	}

	programID, isSuccess, err := s.repo.ReadWriter.WriteProgram(ctx, req)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.AlreadyExists, "failed to add program already existed", err)
	}

	if !isSuccess || len(programID) == 0 {
		level.Error(gvars.Log).Log(lgr.LogErr, "is not success")
		return nil, er.Ebl(codes.AlreadyExists, "failed to add program already existed", err)
	}

	err = s.statusUpdateScheduler(ctx, req, programID)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.AlreadyExists, "failed to create schedulers for program", err)
	}

	var errBulkInsertCustCSV, errBulkInsertMerchCSV error

	if len(req.CustomerCsvUrl) != 0 {
		req.RefreshProgramQuotaDaily = 1
		errBulkInsertCustCSV = s.bulkInsertCustCSV(ctx, req, programID)
	} else {
		req.RefreshProgramQuotaDaily = 0
	}

	if len(req.MerchantCsvUrl) != 0 {
		errBulkInsertMerchCSV = s.bulkInsertMerchCSV(ctx, req, programID)
	}

	if errBulkInsertCustCSV != nil || errBulkInsertMerchCSV != nil {
		_, err := s.repo.ReadWriter.DeleteProgram(ctx, &pb.DeleteProgramReq{Id: programID})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return nil, er.Ebl(codes.NotFound, "failed to delete data because the data is not existed", err)
		}

		level.Error(gvars.Log).Log(lgr.LogErr, errBulkInsertCustCSV, lgr.LogErr, errBulkInsertMerchCSV)
		return nil, er.Ebl(codes.Internal, "customer and merchant whitelist incompatible", nil)
	}

	response.Success = true
	response.Messages = "new program has been add"

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", &response))

	return &response, nil
}

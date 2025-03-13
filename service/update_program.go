package service

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	"gitlab.com/nbdgocean6/nobita-promo-auth/constants"
	"gitlab.com/nbdgocean6/nobita-promo-program/gvars"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-util/er"
	"gitlab.com/nbdgocean6/nobita-util/jwt"
	"gitlab.com/nbdgocean6/nobita-util/lgr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"time"
)

func (s service) EditProgram(ctx context.Context, req *pb.EditProgramReq) (res *pb.ProgramRes, err error) {
	const funcName = `EditProgram`
	const maxProgramPerPeriod = 3
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "Failed to generate jwt", err)
	}

	jwtData, err := jwt.
		NewJWT(jwtToken.Bearer, jwtToken.Secret).
		ExtractKeys([]string{constants.UserID})
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "Failed to extract jwt data", err)
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	currentUnix := now.Unix() * 1000
	req.Status = 0
	if currentUnix >= req.StartAt && currentUnix <= req.EndAt {
		req.Status = 1
	}

	var response pb.ProgramRes
	req.UpdatedBy = jwtData[constants.UserID]

	count, err := s.repo.ReadWriter.ReadProgramsByBetweenStartEndDate(ctx, req.TopicId, req.StartAt, req.EndAt)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.AlreadyExists, "failed to check running program", err)
	}
	if count >= maxProgramPerPeriod {
		return nil, er.Ebl(codes.ResourceExhausted, "Error add program: has reach maximum running concurrently", nil)
	}

	isSuccess, err := s.repo.ReadWriter.UpdateProgram(ctx, req)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.NotFound, "failed to update data not existed", err)
	}

	if !isSuccess {
		level.Error(gvars.Log).Log(lgr.LogErr, "is not success")
		return nil, er.Ebl(codes.NotFound, "failed to update data not existed", err)
	}

	response.Success = true
	response.Messages = "program has been updated"

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", &response))

	return &response, nil
}

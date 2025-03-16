package service

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/oceaninov/naeco-promo-auth/constants"
	"github.com/oceaninov/naeco-promo-program/gvars"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-util/er"
	"github.com/oceaninov/naeco-promo-util/jwt"
	"github.com/oceaninov/naeco-promo-util/lgr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"time"
)

func (s service) EditProgram(ctx context.Context, reqUpdate *pb.EditProgramReq) (res *pb.ProgramRes, err error) {
	const funcName = `EditProgram`
	const maxProgramPerPeriod = 3
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", reqUpdate))

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

	req := &pb.AddProgramReq{
		ChannelId:                reqUpdate.ChannelId,
		TopicId:                  reqUpdate.TopicId,
		Description:              reqUpdate.Description,
		MemoUrl:                  reqUpdate.MemoUrl,
		StartAt:                  reqUpdate.StartAt,
		EndAt:                    reqUpdate.EndAt,
		AllocatedAmount:          reqUpdate.AllocatedAmount,
		AvailableAllocatedAmount: reqUpdate.AvailableAllocatedAmount,
		EligibilityCheck:         reqUpdate.EligibilityCheck,
		Status:                   reqUpdate.Status,
		CreatedBy:                jwtData[constants.UserID],
		UpdatedBy:                jwtData[constants.UserID],
		SourceOfFund:             reqUpdate.SourceOfFund,
		DiscountCalculation:      reqUpdate.DiscountCalculation,
		AllocatedQuota:           reqUpdate.AllocatedQuota,
		AvailableAllocatedQuota:  reqUpdate.AvailableAllocatedQuota,
		DiscountPercent:          reqUpdate.DiscountPercent,
		DiscountAmount:           reqUpdate.DiscountAmount,
		MerchantCsvUrl:           reqUpdate.MerchantCsvUrl,
		CustomerCsvUrl:           reqUpdate.CustomerCsvUrl,
		RefreshProgramQuotaDaily: reqUpdate.RefreshProgramQuotaDaily,
		OnBoardingDateStart:      reqUpdate.OnBoardingDateStart,
		OnBoardingDateTo:         reqUpdate.OnBoardingDateTo,
		RangeTrxAmountMinimum:    reqUpdate.RangeTrxAmountMinimum,
		RangeTrxAmountMaximum:    reqUpdate.RangeTrxAmountMaximum,
		HistoryGroupId:           reqUpdate.HistoryGroupId,
		ProgramChannels:          reqUpdate.ProgramChannels,
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	currentUnix := now.Unix() * 1000
	if currentUnix >= req.StartAt && currentUnix <= req.EndAt {
		req.Status = 1
	} else {
		currentDate := time.Unix(currentUnix/1000, 0).Format("20060102")
		endDate := time.Unix(req.EndAt/1000, 0).Format("20060102")
		if currentDate == endDate {
			req.Status = 1
		} else {
			req.Status = 0
		}
	}

	err = s.repo.ReadWriter.ChangeDeprecatedState(ctx, reqUpdate.Id, true)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to update deprecate state previous program", err)
	}

	var response pb.ProgramRes
	req.CreatedBy = jwtData[constants.UserID]

	count, err := s.repo.ReadWriter.ReadProgramsByBetweenStartEndDate(ctx, req.TopicId, req.Status, req.StartAt, req.EndAt)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.AlreadyExists, "failed to check running program", err)
	}
	if count >= maxProgramPerPeriod {
		return nil, er.Ebl(codes.ResourceExhausted, "Program has reach maximum running concurrently", nil)
	}

	balance, err := s.repo.API.CheckBalance(req.SourceOfFund)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, err.Error(), err)
	}
	if !(balance >= req.AllocatedAmount) {
		return nil, er.Ebl(codes.ResourceExhausted, "Insufficient balance for the source of found", nil)
	}
	_ = s.repo.Cache.RedisSetSourceOfFundBalance(ctx, req.SourceOfFund, balance)

	programID, isSuccess, err := s.repo.ReadWriter.WriteProgram(ctx, req)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.AlreadyExists, "failed to add program already existed", err)
	}

	if !isSuccess || len(programID) == 0 {
		level.Error(gvars.Log).Log(lgr.LogErr, "is not success")
		return nil, er.Ebl(codes.AlreadyExists, "failed to add program already existed", err)
	}

	err = s.repo.ReadWriter.WriteProgramChannels(ctx, programID, req.ProgramChannels)
	if err != nil {
		_, err := s.repo.ReadWriter.DeleteProgram(ctx, &pb.DeleteProgramReq{Id: programID})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return nil, er.Ebl(codes.NotFound, "failed to delete data because the data is not existed", err)
		}

		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to create program channels for program", err)
	}

	err = s.statusUpdateScheduler(req, programID)
	if err != nil {
		_, err := s.repo.ReadWriter.DeleteProgram(ctx, &pb.DeleteProgramReq{Id: programID})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return nil, er.Ebl(codes.NotFound, "failed to delete data because the data is not existed", err)
		}

		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to create schedulers for program", err)
	}

	err = s.refreshDailyQuotaScheduler(req, programID)
	if err != nil {
		_, err := s.repo.ReadWriter.DeleteProgram(ctx, &pb.DeleteProgramReq{Id: programID})
		if err != nil {
			level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
			return nil, er.Ebl(codes.NotFound, "failed to delete data because the data is not existed", err)
		}

		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to create schedulers for program", err)
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

	response.NewProgramId = programID
	response.Success = true
	response.Messages = "new program has been edited"

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", &response))

	return &response, nil
}

//func (s service) EditProgram(ctx context.Context, req *pb.EditProgramReq) (res *pb.ProgramRes, err error) {
//	const funcName = `EditProgram`
//	const maxProgramPerPeriod = 3
//	_, span := s.tracer.StartSpan(ctx, funcName)
//	defer span.End()
//
//	execTime := time.Now()
//
//	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))
//
//	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))
//
//	jwtToken, err := jwt.Bearer(ctx, s.vault.Get("/jwt:secret"))
//	if err != nil {
//		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
//		return nil, er.Ebl(codes.Internal, "Failed to generate jwt", err)
//	}
//
//	jwtData, err := jwt.
//		NewJWT(jwtToken.Bearer, jwtToken.Secret).
//		ExtractKeys([]string{constants.UserID})
//	if err != nil {
//		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
//		return nil, er.Ebl(codes.Internal, "Failed to extract jwt data", err)
//	}
//
//	loc, _ := time.LoadLocation("Asia/Jakarta")
//	now := time.Now().In(loc)
//
//	currentUnix := now.Unix() * 1000
//	req.Status = 0
//	if currentUnix >= req.StartAt && currentUnix <= req.EndAt {
//		req.Status = 1
//	}
//
//	var response pb.ProgramRes
//	req.UpdatedBy = jwtData[constants.UserID]
//
//	count, err := s.repo.ReadWriter.ReadProgramsByBetweenStartEndDate(ctx, req.TopicId, req.StartAt, req.EndAt)
//	if err != nil {
//		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
//		return nil, er.Ebl(codes.AlreadyExists, "failed to check running program", err)
//	}
//	if count >= maxProgramPerPeriod {
//		return nil, er.Ebl(codes.ResourceExhausted, "Error add program: has reach maximum running concurrently", nil)
//	}
//
//	isSuccess, err := s.repo.ReadWriter.UpdateProgram(ctx, req)
//	if err != nil {
//		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
//		return nil, er.Ebl(codes.NotFound, "failed to update data not existed", err)
//	}
//
//	if !isSuccess {
//		level.Error(gvars.Log).Log(lgr.LogErr, "is not success")
//		return nil, er.Ebl(codes.NotFound, "failed to update data not existed", err)
//	}
//
//	response.Success = true
//	response.Messages = "program has been updated"
//
//	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))
//
//	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", &response))
//
//	return &response, nil
//}

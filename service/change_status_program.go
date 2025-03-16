package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/oceaninov/naeco-promo-program/gvars"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-util/er"
	"github.com/oceaninov/naeco-promo-util/lgr"
	"google.golang.org/grpc/codes"
	"time"
)

func (s service) ChangeStatusProgram(ctx context.Context, req *pb.ProgramStatus) (res *pb.Response, err error) {
	const funcName = `ChangeStatusProgram`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	var response pb.Response
	err = s.repo.ReadWriter.ChangeProgramStatus(ctx, req.Id, req.Status)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		response.Messages = "failed to update program status"
		response.Success = false
		return &response, er.Ebl(codes.AlreadyExists, "failed to update program status", err)
	}

	response.Messages = "program status has been change"
	response.Success = true

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", response))

	return &response, nil
}

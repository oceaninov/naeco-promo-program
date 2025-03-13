package service

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	"gitlab.com/nbdgocean6/nobita-promo-program/gvars"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-util/er"
	"gitlab.com/nbdgocean6/nobita-util/lgr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"time"
)

func (s service) DeleteProgram(ctx context.Context, req *pb.DeleteProgramReq) (res *pb.ProgramRes, err error) {
	const funcName = `DeleteProgram`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	var response pb.ProgramRes
	isSuccess, err := s.repo.ReadWriter.DeleteProgram(ctx, req)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.NotFound, "failed to delete data not existed", err)
	}

	if !isSuccess {
		level.Error(gvars.Log).Log(lgr.LogErr, "is not success")
		return nil, er.Ebl(codes.NotFound, "failed to delete data not existed", err)
	}

	response.Success = true
	response.Messages = "program has been deleted"

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", &response))

	return &response, nil
}

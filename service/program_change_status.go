package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"gitlab.com/nbdgocean6/nobita-promo-program/gvars"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-util/er"
	"gitlab.com/nbdgocean6/nobita-util/lgr"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (s service) ProgramChangeStatus(ctx context.Context, req *pb.ProgramStatus) (res *emptypb.Empty, err error) {
	const funcName = `ProgramChangeStatus`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	err = s.repo.ReadWriter.ChangeProgramStatus(ctx, req.Id, req.Status)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return res, er.Ebl(codes.AlreadyExists, "failed to update program status", err)
	}

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", res))

	return res, nil
}

package service

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	"gitlab.com/nbdgocean6/nobita-promo-program/gvars"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-util/er"
	"github.com/oceaninov/naeco-promo-util/lgr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"time"
)

func (s service) ProgramStatusUpdate(ctx context.Context, req *pb.ProgramStatusUpdateRes) (res *pb.ProgramStatusUpdateRes, err error) {
	const funcName = `ProgramStatusUpdate`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	isSuccess, err := s.repo.ReadWriter.UpdateAllProgramByTodayDateToActive(ctx)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return res, er.Ebl(codes.Internal, "Failed Update to Active", err)
	}

	if !isSuccess {
		level.Error(gvars.Log).Log(lgr.LogErr, "Failed to Update Status")
		return res, er.Ebl(codes.Internal, "Failed Update to Active", err)
	}

	isSuccess, err = s.repo.ReadWriter.UpdateAllProgramByTodayDateToInactive(ctx)
	if err != nil {
		println(err.Error())
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return res, er.Ebl(codes.Internal, "Failed Update to Active", err)
	}

	if !isSuccess {
		level.Error(gvars.Log).Log(lgr.LogErr, "Failed to Update Status")
		return res, er.Ebl(codes.Internal, "Failed Update to Active", err)
	}

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", res))

	return &pb.ProgramStatusUpdateRes{
		Status:  "Success",
		Message: "Status Program Has Been Updated",
	}, nil
}

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
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (s service) AddProgramBlacklistsBulk(ctx context.Context, req *pb.Blacklisting) (res *emptypb.Empty, err error) {
	const funcName = `AddProgramBlacklistsBulk`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	err = s.repo.ReadWriter.WriteProgramBlacklistBulk(ctx, req)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return nil, er.Ebl(codes.Internal, "failed to write bulk program blacklist", err)
	}

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", res))

	return nil, nil
}
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

func (s service) GetProgramDetail(ctx context.Context, req *pb.GetProgramDetailReq) (res *pb.Program, err error) {
	const funcName = `GetProgramDetail`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	programs, err := s.repo.ReadWriter.ReadProgramByID(ctx, req.Id)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return res, er.Ebl(codes.AlreadyExists, "failed to get data not found", err)
	}

	if programs.Deprecated {
		level.Warn(gvars.Log).Log(lgr.LogWarn, "program deprecated cannot be shown")
		return res, er.Ebl(codes.InvalidArgument, "program deprecated", err)
	}

	programChannels, err := s.repo.ReadWriter.ReadProgramChannelByProgramID(ctx, programs.Id)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return res, er.Ebl(codes.AlreadyExists, "failed to get data not found", err)
	}

	programs.ProgramChannels = append(programs.ProgramChannels, programChannels...)

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", programs))

	return programs, nil
}

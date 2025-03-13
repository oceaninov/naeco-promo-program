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

func (s service) GetProgramByTopicID(ctx context.Context, req *pb.GetProgramReq) (res *pb.Programs, err error) {
	const funcName = `GetProgramByTopicID`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	execTime := time.Now()

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("upper of %s function", funcName))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("request of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", req))

	res, err = s.repo.ReadWriter.ReadProgramByTopicID(ctx, req.TopicId)
	if err != nil {
		level.Error(gvars.Log).Log(lgr.LogErr, err.Error())
		return res, er.Ebl(codes.AlreadyExists,"failed to get data not found",err)
	}

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("downer of %s function execution start %d", funcName, time.Since(execTime)))

	level.Info(gvars.Log).Log(lgr.LogInfo, fmt.Sprintf("response of %s function", funcName), lgr.LogData, fmt.Sprintf("%+v", res))

	return res, nil
}

package _interface

import (
	"context"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
)

type ReadWrite interface {
	WriteProgram(ctx context.Context, req *pb.AddProgramReq) (string, bool, error)
	UpdateProgram(ctx context.Context, req *pb.EditProgramReq) (bool, error)
	DeleteProgram(ctx context.Context, req *pb.DeleteProgramReq) (bool, error)
	ReadProgramByTopicID(ctx context.Context, id string) (*pb.Programs, error)
	ReadProgramByID(ctx context.Context, id string) (*pb.Program, error)
	ReadProgramsByBetweenStartEndDate(ctx context.Context, topicID string, startAt, endAt int64) (int64, error)
	UpdateAllProgramByTodayDateToActive(ctx context.Context) (bool, error)
	UpdateAllProgramByTodayDateToInactive(ctx context.Context) (bool, error)
	ChangeProgramStatus(ctx context.Context, programID string, status int32) error
}

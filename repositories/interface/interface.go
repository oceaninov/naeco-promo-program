package _interface

import (
	"context"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
)

type Api interface {
	CheckBalance(account string) (res int64, err error)
}

type ReadWrite interface {
	WriteProgram(ctx context.Context, req *pb.AddProgramReq) (string, bool, error)
	UpdateProgram(ctx context.Context, req *pb.EditProgramReq) (bool, error)
	DeleteProgram(ctx context.Context, req *pb.DeleteProgramReq) (bool, error)
	ReadProgramByTopicID(ctx context.Context, id string) (*pb.Programs, error)
	ReadProgramByID(ctx context.Context, id string) (*pb.Program, error)
	ReadProgramsByBetweenStartEndDate(ctx context.Context, topicID string, status int64, startAt, endAt int64) (int64, error)
	ChangeProgramStatus(ctx context.Context, programID string, status int32) error
	ReadAllProgram(ctx context.Context) (*pb.Programs, error)
	ChangeDeprecatedState(ctx context.Context, programID string, state bool) error
	WriteProgramChannels(ctx context.Context, programID string, programChannels []*pb.ProgramChannel) error
	ReadProgramChannelByProgramID(ctx context.Context, id string) ([]*pb.ProgramChannel, error)
	WriteProgramBlacklistBulk(ctx context.Context, req *pb.Blacklisting) error
	RemoveProgramBlacklistBulk(ctx context.Context, req *pb.Blacklisting) error
	ReadProgramBlacklists(ctx context.Context, programID string) (*pb.Blacklists, error)
}

type Cache interface {
	RedisSetSourceOfFundBalance(ctx context.Context, accountNumber string, balance int64) error
}

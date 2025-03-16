package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeGetProgramBlacklistsEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetProgramBlacklists(ctx, request.(*pb.GetBlacklistReq))
		return res, err
	}
}

func (e ProgramEndpoint) GetProgramBlacklists(ctx context.Context, req *pb.GetBlacklistReq) (*pb.Blacklists, error) {
	res, err := e.GetProgramBlacklistsEndpoint(ctx, req)
	if err != nil {
		return &pb.Blacklists{}, err
	}
	return res.(*pb.Blacklists), nil
}


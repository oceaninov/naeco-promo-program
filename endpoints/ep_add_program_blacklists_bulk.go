package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/service/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

func makeAddProgramBlacklistsBulkEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.AddProgramBlacklistsBulk(ctx, request.(*pb.Blacklisting))
		return res, err
	}
}

func (e ProgramEndpoint) AddProgramBlacklistsBulk(ctx context.Context, req *pb.Blacklisting) (*emptypb.Empty, error) {
	res, err := e.AddProgramBlacklistsBulkEndpoint(ctx, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return res.(*emptypb.Empty), nil
}

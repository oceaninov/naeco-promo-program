package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

func makeDeleteProgramBlacklistsBulkEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.DeleteProgramBlacklistsBulk(ctx, request.(*pb.Blacklisting))
		return res, err
	}
}

func (e ProgramEndpoint) DeleteProgramBlacklistsBulk(ctx context.Context, req *pb.Blacklisting) (*emptypb.Empty, error) {
	res, err := e.DeleteProgramBlacklistsBulkEndpoint(ctx, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return res.(*emptypb.Empty), nil
}

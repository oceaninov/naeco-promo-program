package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

func makeGetProgramEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetProgram(ctx, request.(*emptypb.Empty))
		return res, err
	}
}

func (e ProgramEndpoint) GetProgram(ctx context.Context, req *emptypb.Empty) (*pb.Programs, error) {
	res, err := e.GetProgramEndpoint(ctx, req)
	if err != nil {
		return &pb.Programs{}, err
	}
	return res.(*pb.Programs), nil
}

package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeChangeStatusProgramEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.ChangeStatusProgram(ctx, request.(*pb.ProgramStatus))
		return res, err
	}
}

func (e ProgramEndpoint) ChangeStatusProgram(ctx context.Context, req *pb.ProgramStatus) (*pb.Response, error) {
	res, err := e.ChangeStatusProgramEndpoint(ctx, req)
	if err != nil {
		return &pb.Response{}, err
	}
	return res.(*pb.Response), nil
}

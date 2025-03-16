package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeAddProgramEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.AddProgram(ctx, request.(*pb.AddProgramReq))
		return res, err
	}
}

func (e ProgramEndpoint) AddProgram(ctx context.Context, req *pb.AddProgramReq) (*pb.ProgramRes, error) {
	res, err := e.AddProgramEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.ProgramRes), err
	}
	return res.(*pb.ProgramRes), nil
}

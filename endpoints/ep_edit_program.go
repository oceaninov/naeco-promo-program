package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeEditProgramEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.EditProgram(ctx, request.(*pb.EditProgramReq))
		return res, err
	}
}

func (e ProgramEndpoint) EditProgram(ctx context.Context, req *pb.EditProgramReq) (*pb.ProgramRes, error) {
	res, err := e.EditProgramEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.ProgramRes), err
	}
	return res.(*pb.ProgramRes), nil
}

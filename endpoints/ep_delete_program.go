package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/service/interface"
)

func makeDeleteProgramEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.DeleteProgram(ctx, request.(*pb.DeleteProgramReq))
		return res, err
	}
}

func (e ProgramEndpoint) DeleteProgram(ctx context.Context, req *pb.DeleteProgramReq) (*pb.ProgramRes, error) {
	res, err := e.DeleteProgramEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.ProgramRes), err
	}
	return res.(*pb.ProgramRes), nil
}

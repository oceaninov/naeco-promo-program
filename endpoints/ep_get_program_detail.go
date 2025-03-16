package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeGetProgramDetailEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetProgramDetail(ctx, request.(*pb.GetProgramDetailReq))
		return res, err
	}
}

func (e ProgramEndpoint) GetProgramDetail(ctx context.Context, req *pb.GetProgramDetailReq) (*pb.Program, error) {
	res, err := e.GetProgramDetailEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.Program), err
	}
	return res.(*pb.Program), nil
}

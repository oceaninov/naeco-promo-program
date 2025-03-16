package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeGetProgramByTopicIDEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetProgramByTopicID(ctx, request.(*pb.GetProgramReq))
		return res, err
	}
}

func (e ProgramEndpoint) GetProgramByTopicID(ctx context.Context, req *pb.GetProgramReq) (*pb.Programs, error) {
	res, err := e.GetProgramByTopicIDEndpoint(ctx, req)
	if err != nil {
		return &pb.Programs{}, err
	}
	return res.(*pb.Programs), nil
}

package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
)

func makeProgramStatusUpdateEndpoint(usecase _interface.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.ProgramStatusUpdate(ctx, request.(*pb.ProgramStatusUpdateRes))
		return res, err
	}
}

func (e ProgramEndpoint) ProgramStatusUpdate(ctx context.Context, req *pb.ProgramStatusUpdateRes) (*pb.ProgramStatusUpdateRes, error) {
	res, err := e.ProgramStatusUpdateEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.ProgramStatusUpdateRes), err
	}
	return res.(*pb.ProgramStatusUpdateRes), nil
}

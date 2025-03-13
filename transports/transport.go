package transports

import (
	"context"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ep "gitlab.com/nbdgocean6/nobita-promo-program/endpoints"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcProgramServer struct {
	addProgram          grpctransport.Handler
	editProgram         grpctransport.Handler
	deleteProgram       grpctransport.Handler
	getProgramByTopicID grpctransport.Handler
	getProgramDetail    grpctransport.Handler
	programStatusUpdate grpctransport.Handler
	programChangeStatus grpctransport.Handler
}

func (g *grpcProgramServer) AddProgram(ctx context.Context, req *pb.AddProgramReq) (*pb.ProgramRes, error) {
	_, res, err := g.addProgram.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ProgramRes), nil
}

func (g *grpcProgramServer) EditProgram(ctx context.Context, req *pb.EditProgramReq) (*pb.ProgramRes, error) {
	_, res, err := g.editProgram.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ProgramRes), nil
}

func (g *grpcProgramServer) DeleteProgram(ctx context.Context, req *pb.DeleteProgramReq) (*pb.ProgramRes, error) {
	_, res, err := g.deleteProgram.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ProgramRes), nil
}

func (g *grpcProgramServer) GetProgramByTopicID(ctx context.Context, req *pb.GetProgramReq) (*pb.Programs, error) {
	_, res, err := g.getProgramByTopicID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Programs), nil
}

func (g *grpcProgramServer) GetProgramDetail(ctx context.Context, req *pb.GetProgramDetailReq) (*pb.Program, error) {
	_, res, err := g.getProgramDetail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Program), nil
}

func (g *grpcProgramServer) ProgramStatusUpdate(ctx context.Context, req *pb.ProgramStatusUpdateRes) (*pb.ProgramStatusUpdateRes, error) {
	_, res, err := g.programStatusUpdate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ProgramStatusUpdateRes), nil
}

func (g *grpcProgramServer) ProgramChangeStatus(ctx context.Context, req *pb.ProgramStatus) (*emptypb.Empty, error) {
	_, res, err := g.programChangeStatus.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func NewProgramServer(endpoints ep.ProgramEndpoint) pb.ProgramServiceServer {
	options := []grpctransport.ServerOption{
		kitoc.GRPCServerTrace(),
	}
	return &grpcProgramServer{
		addProgram: grpctransport.NewServer(
			endpoints.AddProgramEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		editProgram: grpctransport.NewServer(
			endpoints.EditProgramEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		deleteProgram: grpctransport.NewServer(
			endpoints.DeleteProgramEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		getProgramByTopicID: grpctransport.NewServer(
			endpoints.GetProgramByTopicIDEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		getProgramDetail: grpctransport.NewServer(
			endpoints.GetProgramDetailEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		programStatusUpdate: grpctransport.NewServer(
			endpoints.ProgramStatusUpdateEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		programChangeStatus: grpctransport.NewServer(
			endpoints.ProgramChangeStatusEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
			options...,
		),
	}
}

func decodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeEmptyPbResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &emptypb.Empty{}, nil
}

//

func encodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func decodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

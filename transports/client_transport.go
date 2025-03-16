package transports

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ep "gitlab.com/nbdgocean6/nobita-promo-program/endpoints"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/service/interface"
	"google.golang.org/grpc"
)

func NewGRPProgramClient(conn *grpc.ClientConn) _interface.Service {

	var addProgramEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.ProgramService`
			rpcMethod = `AddProgram`
		)

		addProgramEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.ProgramRes{},
		).Endpoint()
	}

	var editProgramEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.ProgramService`
			rpcMethod = `EditProgram`
		)

		editProgramEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.ProgramRes{},
		).Endpoint()
	}

	var deleteProgramEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.ProgramService`
			rpcMethod = `DeleteProgram`
		)

		deleteProgramEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.ProgramRes{},
		).Endpoint()
	}

	var getProgramByTopicsIDEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.ProgramService`
			rpcMethod = `GetProgramByTopicID`
		)

		getProgramByTopicsIDEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.Programs{},
		).Endpoint()
	}

	var getProgramDetailEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.ProgramService`
			rpcMethod = `GetProgramDetail`
		)

		getProgramDetailEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.Program{},
		).Endpoint()
	}


	return &ep.ProgramEndpoint{
		AddProgramEndpoint:          addProgramEp,
		EditProgramEndpoint:         editProgramEp,
		DeleteProgramEndpoint:       deleteProgramEp,
		GetProgramByTopicIDEndpoint: getProgramByTopicsIDEp,
		GetProgramDetailEndpoint:    getProgramDetailEp,
	}
}

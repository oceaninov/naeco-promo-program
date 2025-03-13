package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	logger "github.com/go-kit/kit/log"
	authsvc "gitlab.com/nbdgocean6/nobita-promo-auth/service/interface"
	"gitlab.com/nbdgocean6/nobita-promo-program/constants"
	"gitlab.com/nbdgocean6/nobita-promo-program/middleware"
	"gitlab.com/nbdgocean6/nobita-util/vlt"

	kitoc "github.com/go-kit/kit/tracing/opencensus"
	authMiddleware "gitlab.com/nbdgocean6/nobita-promo-auth/middleware"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/service/interface"
)

type ProgramEndpoint struct {
	AddProgramEndpoint          endpoint.Endpoint
	EditProgramEndpoint         endpoint.Endpoint
	DeleteProgramEndpoint       endpoint.Endpoint
	GetProgramByTopicIDEndpoint endpoint.Endpoint
	GetProgramDetailEndpoint    endpoint.Endpoint
	ProgramStatusUpdateEndpoint endpoint.Endpoint
	ProgramChangeStatusEndpoint endpoint.Endpoint
}

func NewProgramEndpoint(programSvc _interface.Service, authSvc authsvc.Service, logger logger.Logger, vault vlt.VLT) (ProgramEndpoint, error) {

	var addProgramEp endpoint.Endpoint
	{
		const name = `AddProgram`
		addProgramEp = makeAddProgramEndpoint(programSvc)
		addProgramEp = middleware.LoggingMiddleware(logger)(addProgramEp)
		addProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(addProgramEp)
		addProgramEp = authMiddleware.JwtMiddleware(authSvc, vault)(addProgramEp)
		addProgramEp = kitoc.TraceEndpoint(name)(addProgramEp)
	}

	var editProgramEp endpoint.Endpoint
	{
		const name = `EditProgram`
		editProgramEp = makeEditProgramEndpoint(programSvc)
		editProgramEp = middleware.LoggingMiddleware(logger)(editProgramEp)
		editProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(editProgramEp)
		editProgramEp = authMiddleware.JwtMiddleware(authSvc, vault)(editProgramEp)
		editProgramEp = kitoc.TraceEndpoint(name)(editProgramEp)
	}

	var deleteProgramEp endpoint.Endpoint
	{
		const name = `DeleteProgram`
		deleteProgramEp = makeDeleteProgramEndpoint(programSvc)
		deleteProgramEp = middleware.LoggingMiddleware(logger)(deleteProgramEp)
		deleteProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(deleteProgramEp)
		deleteProgramEp = authMiddleware.JwtMiddleware(authSvc, vault)(deleteProgramEp)
		deleteProgramEp = kitoc.TraceEndpoint(name)(deleteProgramEp)
	}

	var getProgramByTopicsIDEp endpoint.Endpoint
	{
		const name = `GetProgramByTopicID`
		getProgramByTopicsIDEp = makeGetProgramByTopicIDEndpoint(programSvc)
		getProgramByTopicsIDEp = middleware.LoggingMiddleware(logger)(getProgramByTopicsIDEp)
		getProgramByTopicsIDEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(getProgramByTopicsIDEp)
		//getProgramByTopicsIDEp = authMiddleware.JwtMiddleware(authSvc, vault)(getProgramByTopicsIDEp)
		getProgramByTopicsIDEp = kitoc.TraceEndpoint(name)(getProgramByTopicsIDEp)
	}

	var getProgramDetailEp endpoint.Endpoint
	{
		const name = `GetProgramDetail`
		getProgramDetailEp = makeGetProgramDetailEndpoint(programSvc)
		getProgramDetailEp = middleware.LoggingMiddleware(logger)(getProgramDetailEp)
		getProgramDetailEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(getProgramDetailEp)
		getProgramDetailEp = authMiddleware.JwtMiddleware(authSvc, vault)(getProgramDetailEp)
		getProgramDetailEp = kitoc.TraceEndpoint(name)(getProgramDetailEp)
	}

	var programStatusUpdateEp endpoint.Endpoint
	{
		const name = `ProgramStatusUpdate`
		programStatusUpdateEp = makeProgramStatusUpdateEndpoint(programSvc)
		programStatusUpdateEp = middleware.LoggingMiddleware(logger)(programStatusUpdateEp)
		programStatusUpdateEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(programStatusUpdateEp)
		programStatusUpdateEp = authMiddleware.JwtMiddleware(authSvc, vault)(programStatusUpdateEp)
		programStatusUpdateEp = kitoc.TraceEndpoint(name)(programStatusUpdateEp)
	}

	var programChangeStatusEndpoint endpoint.Endpoint
	{
		const name = `ProgramChangeStatus`
		programChangeStatusEndpoint = makeProgramChangeStatusEndpoint(programSvc)
		programChangeStatusEndpoint = middleware.LoggingMiddleware(logger)(programChangeStatusEndpoint)
		programChangeStatusEndpoint = middleware.CircuitBreakerMiddleware(constants.ServiceName)(programChangeStatusEndpoint)
		programChangeStatusEndpoint = middleware.BasicAuthMiddleware()(programChangeStatusEndpoint)
		programChangeStatusEndpoint = kitoc.TraceEndpoint(name)(programChangeStatusEndpoint)
	}

	return ProgramEndpoint{
		AddProgramEndpoint:          addProgramEp,
		EditProgramEndpoint:         editProgramEp,
		DeleteProgramEndpoint:       deleteProgramEp,
		GetProgramByTopicIDEndpoint: getProgramByTopicsIDEp,
		GetProgramDetailEndpoint:    getProgramDetailEp,
		ProgramStatusUpdateEndpoint: programStatusUpdateEp,
		ProgramChangeStatusEndpoint: programChangeStatusEndpoint,
	}, nil
}

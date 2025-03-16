package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	logger "github.com/go-kit/kit/log"
	authsvc "github.com/oceaninov/naeco-promo-auth/service/interface"
	"github.com/oceaninov/naeco-promo-util/vlt"
	"gitlab.com/nbdgocean6/nobita-promo-program/constants"
	"gitlab.com/nbdgocean6/nobita-promo-program/middleware"

	kitoc "github.com/go-kit/kit/tracing/opencensus"
	authMiddleware "github.com/oceaninov/naeco-promo-auth/middleware"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/service/interface"
)

type ProgramEndpoint struct {
	AddProgramEndpoint                  endpoint.Endpoint
	EditProgramEndpoint                 endpoint.Endpoint
	DeleteProgramEndpoint               endpoint.Endpoint
	GetProgramByTopicIDEndpoint         endpoint.Endpoint
	GetProgramDetailEndpoint            endpoint.Endpoint
	ProgramChangeStatusEndpoint         endpoint.Endpoint
	GetProgramEndpoint                  endpoint.Endpoint
	ChangeStatusProgramEndpoint         endpoint.Endpoint
	AddProgramBlacklistsBulkEndpoint    endpoint.Endpoint
	DeleteProgramBlacklistsBulkEndpoint endpoint.Endpoint
	GetProgramBlacklistsEndpoint        endpoint.Endpoint
}

func NewProgramEndpoint(programSvc _interface.Service, authSvc authsvc.Service, logger logger.Logger, vault vlt.VLT) (ProgramEndpoint, error) {

	var addProgramEp endpoint.Endpoint
	{
		const name = `AddProgram`
		addProgramEp = makeAddProgramEndpoint(programSvc)
		addProgramEp = middleware.LoggingMiddleware(logger)(addProgramEp)
		addProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(addProgramEp)
		addProgramEp = authMiddleware.JwtMiddleware(authSvc)(addProgramEp)
		addProgramEp = kitoc.TraceEndpoint(name)(addProgramEp)
	}

	var editProgramEp endpoint.Endpoint
	{
		const name = `EditProgram`
		editProgramEp = makeEditProgramEndpoint(programSvc)
		editProgramEp = middleware.LoggingMiddleware(logger)(editProgramEp)
		editProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(editProgramEp)
		editProgramEp = authMiddleware.JwtMiddleware(authSvc)(editProgramEp)
		editProgramEp = kitoc.TraceEndpoint(name)(editProgramEp)
	}

	var deleteProgramEp endpoint.Endpoint
	{
		const name = `DeleteProgram`
		deleteProgramEp = makeDeleteProgramEndpoint(programSvc)
		deleteProgramEp = middleware.LoggingMiddleware(logger)(deleteProgramEp)
		deleteProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(deleteProgramEp)
		deleteProgramEp = authMiddleware.JwtMiddleware(authSvc)(deleteProgramEp)
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
		getProgramDetailEp = authMiddleware.JwtMiddleware(authSvc)(getProgramDetailEp)
		getProgramDetailEp = kitoc.TraceEndpoint(name)(getProgramDetailEp)
	}

	var getProgramEp endpoint.Endpoint
	{
		const name = `GetProgram`
		getProgramEp = makeGetProgramEndpoint(programSvc)
		getProgramEp = middleware.LoggingMiddleware(logger)(getProgramEp)
		getProgramEp = middleware.CircuitBreakerMiddleware(constants.ServiceName)(getProgramEp)
		getProgramEp = authMiddleware.JwtMiddleware(authSvc)(getProgramEp)
		getProgramEp = kitoc.TraceEndpoint(name)(getProgramEp)
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

	var changeStatusProgramEndpoint endpoint.Endpoint
	{
		const name = `ChangeStatusProgram`
		changeStatusProgramEndpoint = makeChangeStatusProgramEndpoint(programSvc)
		//changeStatusProgramEndpoint = middleware.LoggingMiddleware(logger)(changeStatusProgramEndpoint)
		//changeStatusProgramEndpoint = middleware.CircuitBreakerMiddleware(constants.ServiceName)(changeStatusProgramEndpoint)
		//changeStatusProgramEndpoint = authMiddleware.JwtMiddleware(authSvc, vault)(changeStatusProgramEndpoint)
		changeStatusProgramEndpoint = kitoc.TraceEndpoint(name)(changeStatusProgramEndpoint)
	}

	var addProgramBlacklistsBulkEndpoint endpoint.Endpoint
	{
		const name = `AddProgramBlacklistsBulk`
		addProgramBlacklistsBulkEndpoint = makeAddProgramBlacklistsBulkEndpoint(programSvc)
		addProgramBlacklistsBulkEndpoint = middleware.LoggingMiddleware(logger)(addProgramBlacklistsBulkEndpoint)
		addProgramBlacklistsBulkEndpoint = middleware.CircuitBreakerMiddleware(constants.ServiceName)(addProgramBlacklistsBulkEndpoint)
		addProgramBlacklistsBulkEndpoint = authMiddleware.JwtMiddleware(authSvc)(addProgramBlacklistsBulkEndpoint)
		addProgramBlacklistsBulkEndpoint = kitoc.TraceEndpoint(name)(addProgramBlacklistsBulkEndpoint)
	}

	var deleteProgramBlacklistsBulkEndpoint endpoint.Endpoint
	{
		const name = `DeleteProgramBlacklistsBulk`
		deleteProgramBlacklistsBulkEndpoint = makeDeleteProgramBlacklistsBulkEndpoint(programSvc)
		deleteProgramBlacklistsBulkEndpoint = middleware.LoggingMiddleware(logger)(deleteProgramBlacklistsBulkEndpoint)
		deleteProgramBlacklistsBulkEndpoint = middleware.CircuitBreakerMiddleware(constants.ServiceName)(deleteProgramBlacklistsBulkEndpoint)
		deleteProgramBlacklistsBulkEndpoint = authMiddleware.JwtMiddleware(authSvc)(deleteProgramBlacklistsBulkEndpoint)
		deleteProgramBlacklistsBulkEndpoint = kitoc.TraceEndpoint(name)(deleteProgramBlacklistsBulkEndpoint)
	}

	var getProgramBlacklistsEndpoint endpoint.Endpoint
	{
		const name = `GetProgramBlacklists`
		getProgramBlacklistsEndpoint = makeGetProgramBlacklistsEndpoint(programSvc)
		getProgramBlacklistsEndpoint = middleware.LoggingMiddleware(logger)(getProgramBlacklistsEndpoint)
		getProgramBlacklistsEndpoint = middleware.CircuitBreakerMiddleware(constants.ServiceName)(getProgramBlacklistsEndpoint)
		getProgramBlacklistsEndpoint = authMiddleware.JwtMiddleware(authSvc)(getProgramBlacklistsEndpoint)
		getProgramBlacklistsEndpoint = kitoc.TraceEndpoint(name)(getProgramBlacklistsEndpoint)
	}

	return ProgramEndpoint{
		AddProgramEndpoint:                  addProgramEp,
		EditProgramEndpoint:                 editProgramEp,
		DeleteProgramEndpoint:               deleteProgramEp,
		GetProgramByTopicIDEndpoint:         getProgramByTopicsIDEp,
		GetProgramDetailEndpoint:            getProgramDetailEp,
		ProgramChangeStatusEndpoint:         programChangeStatusEndpoint,
		GetProgramEndpoint:                  getProgramEp,
		ChangeStatusProgramEndpoint:         changeStatusProgramEndpoint,
		AddProgramBlacklistsBulkEndpoint:    addProgramBlacklistsBulkEndpoint,
		DeleteProgramBlacklistsBulkEndpoint: deleteProgramBlacklistsBulkEndpoint,
		GetProgramBlacklistsEndpoint:        getProgramBlacklistsEndpoint,
	}, nil
}

package service

import (
	_repointerface "gitlab.com/nbdgocean6/nobita-promo-program/repositories"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/service/interface"
	"gitlab.com/nbdgocean6/nobita-util/vlt"
	"go.opencensus.io/trace"
)

type service struct {
	tracer trace.Tracer
	vault  vlt.VLT
	repo   _repointerface.Repository
}

func NewUsecases(repo _repointerface.Repository, tracer trace.Tracer, vault vlt.VLT) _interface.Service {
	return &service{
		tracer: tracer,
		repo:   repo,
		vault:  vault,
	}
}

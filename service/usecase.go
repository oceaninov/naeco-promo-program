package service

import (
	_repointerface "github.com/oceaninov/naeco-promo-program/repositories"
	_interface "github.com/oceaninov/naeco-promo-program/service/interface"
	"github.com/oceaninov/naeco-promo-util/vlt"
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

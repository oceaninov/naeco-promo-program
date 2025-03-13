package repositories

import (
	"context"
	clockwerksvc "github.com/nightsilvertech/clockwerk/service/interface"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/repositories/interface"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/microservices"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/scheduler"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/sql"
	"gitlab.com/nbdgocean6/nobita-util/dbc"
	"go.opencensus.io/trace"
)

type Repository struct {
	ReadWriter   _interface.ReadWrite
	Microservice microservices.Microservices
	Scheduler    clockwerksvc.Clockwerk
}

type RepoConf struct {
	SQL              dbc.Config
	MicroserviceConf microservices.Config
	SchedulerConf    scheduler.Config
}

func NewRepositories(ctx context.Context, rc RepoConf, tracer trace.Tracer) (*Repository, error) {
	readWriter, err := sql.NewSQL(rc.SQL, tracer)
	if err != nil {
		return nil, err
	}
	microsvc, err := microservices.NewMicroservice(ctx, rc.MicroserviceConf)
	if err != nil {
		return nil, err
	}
	schedulers, err := scheduler.NewScheduler(rc.SchedulerConf)
	return &Repository{
		ReadWriter:   readWriter,
		Microservice: *microsvc,
		Scheduler:    schedulers,
	}, nil
}

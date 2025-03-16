package repositories

import (
	"context"
	clockwerk "github.com/nightsilvertech/clockwerk/client"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/api"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/repositories/interface"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/microservices"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/redis"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/scheduler"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/sql"
	"github.com/oceaninov/naeco-promo-util/dbc"
	"github.com/oceaninov/naeco-promo-util/vlt"
	"go.opencensus.io/trace"
)

type Repository struct {
	ReadWriter   _interface.ReadWrite
	API          _interface.Api
	Microservice microservices.Microservices
	Scheduler    clockwerk.ClockwerkClient
	Cache        _interface.Cache
}

type RepoConf struct {
	SQL              dbc.Config
	MicroserviceConf microservices.Config
	SchedulerConf    scheduler.Config
}

func NewRepositories(ctx context.Context, rc RepoConf, tracer trace.Tracer, vault vlt.VLT) (*Repository, error) {
	readWriter, err := sql.NewSQL(rc.SQL, tracer)
	if err != nil {
		return nil, err
	}
	microsvc, err := microservices.NewMicroservice(ctx, rc.MicroserviceConf)
	if err != nil {
		return nil, err
	}
	schedulers, err := scheduler.NewSchedulerV2(rc.SchedulerConf, vault)
	return &Repository{
		ReadWriter:   readWriter,
		Microservice: *microsvc,
		Scheduler:    schedulers,
		API:          api.NewAPI(),
		Cache:        redis.NewRedis(),
	}, nil
}

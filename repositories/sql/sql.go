package sql

import (
	"database/sql"
	"go.opencensus.io/trace"
	"sync"

	_interface "gitlab.com/nbdgocean6/nobita-promo-program/repositories/interface"
	"gitlab.com/nbdgocean6/nobita-util/dbc"
)

var mutex = &sync.RWMutex{}

type readWrite struct {
	tracer trace.Tracer
	db     *sql.DB
}

func NewSQL(config dbc.Config, tracer trace.Tracer) (_interface.ReadWrite, error) {
	sqlDB, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}

	return &readWrite{
		db:     sqlDB,
		tracer: tracer,
	}, nil
}

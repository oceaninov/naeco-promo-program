package sql

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"strings"
	"time"
)

func (r *readWrite) WriteProgramChannels(ctx context.Context, programID string, programChannels []*pb.ProgramChannel) error {
	const funcName = `WriteProgramChannels`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var valueStrings []string
	var valueArgs []interface{}
	for _, pc := range programChannels {
		id := uuid.NewV4().String()
		valueStrings = append(valueStrings, "(?,?,?,?,?)")
		valueArgs = append(valueArgs, id)
		valueArgs = append(valueArgs, programID)
		valueArgs = append(valueArgs, pc.Id)
		valueArgs = append(valueArgs, time.Now())
		valueArgs = append(valueArgs, time.Now())
	}
	if len(valueStrings) == 0 {
		return errors.New("no program channels data to insert the length is zero")
	}

	query := fmt.Sprintf("INSERT INTO program_channels(id, program_id, channel_id, created_at, updated_at) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := r.db.Exec(query, valueArgs...)
	if err != nil {
		return err
	}

	return nil
}

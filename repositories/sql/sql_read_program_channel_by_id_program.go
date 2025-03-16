package sql

import (
	"context"
	"database/sql"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"time"
)

func (r *readWrite) ReadProgramChannelByProgramID(ctx context.Context, id string) ([]*pb.ProgramChannel, error) {
	const funcName = `ReadProgramChannelByProgramID`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var program pb.ProgramChannel
	var programs []*pb.ProgramChannel
	var created, updated time.Time
	const query = `
		SELECT 
			channels.id,
			channels.key,
			channels.status,
			channels.title,
			channels.created_at,
			channels.created_by,
			channels.updated_at,
			channels.updated_by,
			program_channels.id AS program_channel_id
		FROM program_channels
		LEFT JOIN channels ON channels.id = program_channels.channel_id 
		WHERE program_channels.program_id = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return programs, err
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		return programs, err
	}
	mutex.Unlock()
	for row.Next() {
		err := row.Scan(
			&program.Id,               // id
			&program.Key,              // key
			&program.Status,           // status
			&program.Title,            // title
			&created,                  // created_at
			&program.CreatedBy,        // created_by
			&updated,                  // updated_at
			&program.UpdatedBy,        // updated_by
			&program.ProgramChannelId, // program_channel_id
		)
		if err != nil {
			return nil, err
		}
		program.CreatedAt = created.Unix()
		program.UpdatedAt = updated.Unix()
		programs = append(programs, &pb.ProgramChannel{
			Id:               program.Id,
			Key:              program.Key,
			Status:           program.Status,
			Title:            program.Title,
			CreatedAt:        program.CreatedAt,
			CreatedBy:        program.CreatedBy,
			UpdatedAt:        program.UpdatedAt,
			UpdatedBy:        program.UpdatedBy,
			ProgramChannelId: program.ProgramChannelId,
		})
	}

	return programs, nil
}

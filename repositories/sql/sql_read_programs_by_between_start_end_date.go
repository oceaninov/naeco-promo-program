package sql

import (
	"context"
	"database/sql"
)

func (r *readWrite) ReadProgramsByBetweenStartEndDate(ctx context.Context, topicID string, status int64, startAt, endAt int64) (int64, error) {
	const funcName = `ReadProgramsByBetweenStartEndDate`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var data int64
	const query = `SELECT count(*) FROM programs WHERE topic_id = ? AND deprecated = 0 AND status = ? AND ((? >= start_at AND ? <= end_at) OR (? <= end_at AND ? >= start_at))`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return data, err
	}
	mutex.Lock()
	row := stmt.QueryRowContext(ctx, topicID, status, startAt, startAt, endAt, endAt)
	mutex.Unlock()
	err = row.Scan(&data)
	if err != nil && err == sql.ErrNoRows {
		return data, err
	}
	return data, nil
}

package sql

import (
	"context"
	"database/sql"
)

func (r *readWrite) ReadTopicsTotalProgram(ctx context.Context, id string) (int64, error) {
	const funcName = `ReadTopicsTotalProgram`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var data int64
	const query = `SELECT count(id) FROM programs WHERE topic_id = ?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return data, err
	}
	mutex.Lock()
	row := stmt.QueryRowContext(ctx, id)
	mutex.Unlock()
	err = row.Scan(&data)
	if err != nil && err == sql.ErrNoRows {
		return data, err
	}
	return data, nil
}

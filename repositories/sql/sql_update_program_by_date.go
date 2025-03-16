package sql

import (
	"context"
	"github.com/oceaninov/naeco-promo-util/dbtrx"
)

func (r *readWrite) UpdateAllProgramByTodayDateToActive(ctx context.Context) (bool, error) {
	const funcName = `UpdateAllProgramByTodayDateToActive`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer dbtrx.Trx(tx, err)
	const query = `UPDATE programs SET
					status = 1
					WHERE (unix_timestamp(current_timestamp)*1000) >= start_at AND (unix_timestamp(current_timestamp)*1000) <= end_at`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *readWrite) UpdateAllProgramByTodayDateToInactive(ctx context.Context) (bool, error) {
	const funcName = `UpdateAllProgramByTodayDateToInactive`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer dbtrx.Trx(tx, err)
	const query = `UPDATE programs SET
					status = 0
					WHERE (unix_timestamp(current_timestamp)*1000) <= start_at OR (unix_timestamp(current_timestamp)*1000) >= end_at`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

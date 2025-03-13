package sql

import (
	"context"
	"errors"
	"gitlab.com/nbdgocean6/nobita-util/dbtrx"
	"time"
)

func (r *readWrite) ChangeProgramStatus(ctx context.Context, programID string, status int32) error {
	const funcName = `ChangeProgramStatus`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)
	const query = `UPDATE programs SET  
            status = ?,
            updated_at = ?
		WHERE id = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(
		ctx,
		status,     // status
		time.Now(), // updated_at
		programID,  // id
	)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil && rowAffected == 0 {
		return err
	}
	if rowAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

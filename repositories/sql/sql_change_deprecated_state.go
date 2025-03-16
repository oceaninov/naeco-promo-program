package sql

import (
	"context"
	"errors"
	"github.com/oceaninov/naeco-promo-util/dbtrx"
	"time"
)

func (r *readWrite) ChangeDeprecatedState(ctx context.Context, programID string, state bool) error {
	const funcName = `ChangeDeprecatedState`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)
	const query = `
		UPDATE programs SET  
            deprecated = ?,
            updated_at = ?
		WHERE id = ?
	`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(
		ctx,
		state,      // deprecated
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

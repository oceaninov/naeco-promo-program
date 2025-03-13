package sql

import (
	"context"

	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-util/dbtrx"
)

func (r *readWrite) DeleteProgram(ctx context.Context, req *pb.DeleteProgramReq) (bool, error) {
	const funcName = `DeleteProgram`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer dbtrx.Trx(tx, err)
	const query = `DELETE FROM programs WHERE id = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return false, err
	}
	result, err := stmt.ExecContext(ctx, req.Id)
	if err != nil {
		return false, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil && rowAffected == 0 {
		return false, err
	}
	if rowAffected == 0 {
		return false, nil
	}
	return true, nil
}

package sql

import (
	"context"
	"time"

	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-util/dbtrx"
)

func (r *readWrite) UpdateProgram(ctx context.Context, req *pb.EditProgramReq) (bool, error) {
	const funcName = `UpdateProgram`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer dbtrx.Trx(tx, err)
	const query = `
		UPDATE programs SET 
			channel_id = ?,
            topic_id = ?,
            description = ?, 
            memo_url = ?, 
            start_at = ?, 
            end_at = ?, 
            allocated_amount = ?, 
            available_allocated_amount = ?, 
            eligibility_check = ?, 
            status = ?,
            source_of_fund = ?, 
            discount_calculation = ?,
            allocated_quota = ?, 
            available_allocated_quota = ?, 
        	discount_amount = ?,
        	discount_percent = ?,
        	merchant_csv_url = ?,
        	customer_csv_url = ?,
            updated_at = ?, 
            updated_by = ?,
            refresh_program_quota_daily = ?,
            on_boarding_date_start = ?,
            on_boarding_date_to = ?,
            range_trx_amount_minimum = ?,
            range_trx_amount_maximum = ?
		WHERE id = ?
	`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return false, err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.ChannelId,                // channel_id
		req.TopicId,                  // topic_id
		req.Description,              // description
		req.MemoUrl,                  // memo_url
		req.StartAt,                  // start_at
		req.EndAt,                    // end_at
		req.AllocatedAmount,          // allocated_amount
		req.AvailableAllocatedAmount, // available_allocated_amount
		req.EligibilityCheck,         // eligibility_check
		req.Status,                   // status
		req.SourceOfFund,             // source_of_fund
		req.DiscountCalculation,      // discount_calculation
		req.AllocatedQuota,           // allocated_quota
		req.AvailableAllocatedQuota,  // available_allocated_quota
		req.DiscountAmount,           // discount_amount
		req.DiscountPercent,          // discount_percent
		req.MerchantCsvUrl,           // merchant_csv_url
		req.CustomerCsvUrl,           // customer_csv_url
		time.Now(),                   // updated_at
		req.UpdatedBy,                // updated_by
		req.RefreshProgramQuotaDaily, // refresh_program_quota_daily
		req.OnBoardingDateStart,      // on_boarding_date_start
		req.OnBoardingDateTo,         // on_boarding_date_to
		req.RangeTrxAmountMinimum,    // range_trx_amount_minimum
		req.RangeTrxAmountMaximum,    // range_trx_amount_maximum
		req.Id,                       // id
	)
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

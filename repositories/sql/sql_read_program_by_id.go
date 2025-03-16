package sql

import (
	"context"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"time"
)

func (r *readWrite) ReadProgramByID(ctx context.Context, id string) (*pb.Program, error) {
	const funcName = `ReadProgramByID`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var program pb.Program
	var created, updated time.Time
	const query = `
		SELECT 
			programs.id,
			programs.channel_id,
			programs.topic_id,
			programs.description,
			programs.memo_url,
			programs.start_at,
			programs.end_at,
			programs.allocated_amount,
			programs.available_allocated_amount,
			programs.eligibility_check,
			programs.status,
			programs.created_at,
			programs.created_by,
			programs.updated_at,
			programs.updated_by,
			programs.source_of_fund,
			programs.discount_calculation,
			programs.allocated_quota,
			programs.available_allocated_quota,
			programs.discount_percent,
			programs.discount_amount,
			programs.merchant_csv_url,
			programs.customer_csv_url,
			programs.refresh_program_quota_daily,
			programs.on_boarding_date_start,
			programs.on_boarding_date_to,
			programs.range_trx_amount_minimum,
			programs.range_trx_amount_maximum,
			programs.deprecated,
			programs.history_group_id,
			topics.title 
		FROM programs
		LEFT JOIN topics ON topics.id = programs.topic_id 
		WHERE programs.id = ?
	`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	mutex.Lock()
	row := stmt.QueryRowContext(ctx, id)
	mutex.Unlock()
	err = row.Scan(
		&program.Id,                       // id
		&program.ChannelId,                // channel_id
		&program.TopicId,                  // topic_id
		&program.Description,              // description
		&program.MemoUrl,                  // memo_url
		&program.StartAt,                  // start_at
		&program.EndAt,                    // end_at
		&program.AllocatedAmount,          // allocated_amount
		&program.AvailableAllocatedAmount, // available_allocated_amount
		&program.EligibilityCheck,         // eligibility_check
		&program.Status,                   // status
		&created,                          // created_at
		&program.CreatedBy,                // created_by
		&updated,                          // updated_at
		&program.UpdatedBy,                // updated_by
		&program.SourceOfFund,             // source_of_fund
		&program.DiscountCalculation,      // discount_calculation
		&program.AllocatedQuota,           // allocated_quota
		&program.AvailableAllocatedQuota,  // available_allocated_quota
		&program.DiscountPercent,          // discount_amount
		&program.DiscountAmount,           // discount_percent
		&program.MerchantCsvUrl,           // merchant_csv_url
		&program.CustomerCsvUrl,           // customer_csv_url
		&program.RefreshProgramQuotaDaily, // created_at
		&program.OnBoardingDateStart,      // on_boarding_date_start
		&program.OnBoardingDateTo,         // on_boarding_date_to
		&program.RangeTrxAmountMinimum,    // range_trx_amount_minimum
		&program.RangeTrxAmountMaximum,    // range_trx_amount_maximum
		&program.Deprecated,               // deprecated
		&program.HistoryGroupId,           // history_group_id
		&program.TopicTitle,               // topic_title
	)
	if err != nil {
		return nil, err
	}
	program.CreatedAt = created.Unix()
	program.UpdatedAt = updated.Unix()
	return &program, nil
}

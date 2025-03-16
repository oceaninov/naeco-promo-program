package sql

import (
	"context"
	"database/sql"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"time"
)

func (r *readWrite) ReadAllProgram(ctx context.Context) (*pb.Programs, error) {
	const funcName = `ReadAllProgram`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var programs pb.Programs
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
			topics.title
		FROM programs 
		INNER JOIN topics ON topics.id = programs.topic_id
		WHERE programs.deprecated = 0 
		ORDER BY created_at DESC
	`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return &programs, err
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx)
	if err != nil && err == sql.ErrNoRows {
		return &programs, err
	}
	mutex.Unlock()
	for row.Next() {
		err := row.Scan(
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
			&program.DiscountPercent,          // discount_percent
			&program.DiscountAmount,           // discount_amount
			&program.MerchantCsvUrl,           // merchant_csv_url
			&program.CustomerCsvUrl,           // customer_csv_url
			&program.RefreshProgramQuotaDaily, // refresh_program_quota_daily
			&program.OnBoardingDateStart,      // on_boarding_date_start
			&program.OnBoardingDateTo,         // on_boarding_date_to
			&program.RangeTrxAmountMinimum,    // range_trx_amount_minimum
			&program.RangeTrxAmountMaximum,    // range_trx_amount_maximum
			&program.TopicTitle,               // TopicTitle
		)
		if err != nil {
			return nil, err
		}
		program.CreatedAt = created.Unix()
		program.UpdatedAt = updated.Unix()
		programs.Programs = append(programs.Programs, &pb.Program{
			Id:                       program.Id,
			ChannelId:                program.ChannelId,
			TopicId:                  program.TopicId,
			Description:              program.Description,
			MemoUrl:                  program.MemoUrl,
			StartAt:                  program.StartAt,
			EndAt:                    program.EndAt,
			AllocatedAmount:          program.AllocatedAmount,
			AvailableAllocatedAmount: program.AvailableAllocatedAmount,
			EligibilityCheck:         program.EligibilityCheck,
			Status:                   program.Status,
			CreatedAt:                program.CreatedAt,
			CreatedBy:                program.CreatedBy,
			UpdatedAt:                program.UpdatedAt,
			UpdatedBy:                program.UpdatedBy,
			SourceOfFund:             program.SourceOfFund,
			DiscountCalculation:      program.DiscountCalculation,
			AllocatedQuota:           program.AllocatedQuota,
			AvailableAllocatedQuota:  program.AvailableAllocatedQuota,
			DiscountAmount:           program.DiscountAmount,
			DiscountPercent:          program.DiscountPercent,
			MerchantCsvUrl:           program.MerchantCsvUrl,
			CustomerCsvUrl:           program.CustomerCsvUrl,
			RefreshProgramQuotaDaily: program.RefreshProgramQuotaDaily,
			OnBoardingDateStart:      program.OnBoardingDateStart,
			OnBoardingDateTo:         program.OnBoardingDateTo,
			RangeTrxAmountMinimum:    program.RangeTrxAmountMinimum,
			RangeTrxAmountMaximum:    program.RangeTrxAmountMaximum,
			TopicTitle:               program.TopicTitle,
		})
	}
	return &programs, nil
}

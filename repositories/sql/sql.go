package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	_interface "gitlab.com/nbdgocean6/nobita-promo-program/repositories/interface"
	"github.com/oceaninov/naeco-promo-util/dbc"
	"go.opencensus.io/trace"
	"strings"
	"sync"
	"time"
)

var mutex = &sync.RWMutex{}

type readWrite struct {
	tracer trace.Tracer
	db     *sql.DB
}

func (r *readWrite) WriteProgramBlacklistBulk(ctx context.Context, req *pb.Blacklisting) error {
	const funcName = `WriteProgramBlacklistBulk`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var valueStrings []string
	var valueArgs []interface{}
	for _, program := range req.BlacklistsId {
		valueStrings = append(valueStrings, "(?,?)")
		valueArgs = append(valueArgs, req.ProgramId)
		valueArgs = append(valueArgs, program.Id)
	}
	if len(valueStrings) == 0 {
		return errors.New("no blacklist data to create")
	}

	query := fmt.Sprintf("INSERT INTO program_blacklists(program_id, blacklist_program_id) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := r.db.Exec(query, valueArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (r *readWrite) RemoveProgramBlacklistBulk(ctx context.Context, req *pb.Blacklisting) error {
	const funcName = `RemoveProgramBlacklistBulk`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var valueStrings []string
	var valueArgs []interface{}
	for i, program := range req.BlacklistsId {
		if i == 0 {
			valueStrings = append(valueStrings, "WHERE (program_id = ? AND blacklist_program_id = ?)")
		} else {
			valueStrings = append(valueStrings, "OR (program_id = ? AND blacklist_program_id = ?)")
		}
		valueArgs = append(valueArgs, req.ProgramId)
		valueArgs = append(valueArgs, program.Id)
	}
	if len(valueStrings) == 0 {
		return errors.New("no blacklist data to create")
	}

	query := fmt.Sprintf("DELETE FROM program_blacklists %s",
		strings.Join(valueStrings, " "))
	_, err := r.db.Exec(query, valueArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (r *readWrite) ReadProgramBlacklists(ctx context.Context, programID string) (res *pb.Blacklists, err error) {
	const funcName = `ReadProgramBlacklists`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var blacklists pb.Blacklists
	var blacklist pb.Blacklist
	var program pb.Program
	var created, updated time.Time
	const query = `
		SELECT
			program_blacklists.program_id,
			program_blacklists.blacklist_program_id,
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
			programs.range_trx_amount_maximum
		FROM program_blacklists
		LEFT JOIN programs ON programs.id = program_blacklists.blacklist_program_id
		WHERE program_blacklists.program_id = ?
	`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return res, err
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx, programID)
	if err != nil && err == sql.ErrNoRows {
		return res, err
	}
	mutex.Unlock()
	for row.Next() {
		err := row.Scan(
			&blacklist.ProgramId,              // program_id
			&blacklist.BlacklistProgramId,     // blacklist_program_id
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
		)
		if err != nil {
			return nil, err
		}
		program.CreatedAt = created.Unix()
		program.UpdatedAt = updated.Unix()
		blacklists.Blacklists = append(blacklists.Blacklists, &pb.Blacklist{
			ProgramId:          blacklist.ProgramId,
			BlacklistProgramId: blacklist.BlacklistProgramId,
			Program: &pb.Program{
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
			},
		})
	}
	return &blacklists, nil
}

func NewSQL(config dbc.Config, tracer trace.Tracer) (_interface.ReadWrite, error) {
	sqlDB, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}

	return &readWrite{
		db:     sqlDB,
		tracer: tracer,
	}, nil
}

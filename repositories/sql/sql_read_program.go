package sql

import (
	"context"
	"database/sql"
	"gitlab.com/nbdgocean6/nobita-promo-program/models"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
)

func (r *readWrite) ReadProgramByTopicID(ctx context.Context, id string) (*pb.Programs, error) {
	const funcName = `ReadProgramByTopicID`
	_, span := r.tracer.StartSpan(ctx, funcName)
	defer span.End()

	var programs pb.Programs
	var program models.Programs
	const query = `SELECT programs.*, topics.title FROM programs
					LEFT JOIN topics ON topics.id = programs.topic_id WHERE topic_id = ? ORDER BY created_at DESC`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return &programs, err
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx, id)
	if err != nil && err == sql.ErrNoRows {
		return &programs, err
	}
	mutex.Unlock()
	for row.Next() {
		err := row.Scan(program.FastScan()...)
		if err != nil {
			return nil, err
		}
		program.UseUnixTimestamp()
		programs.Programs = append(programs.Programs, &pb.Program{
			Id:                       program.ID,
			ChannelId:                program.ChannelId,
			TopicId:                  program.TopicsID,
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
			TopicTitle:               program.TopicTitle,
			RefreshProgramQuotaDaily: program.RefreshProgramQuotaDaily,
		})
	}
	return &programs, nil
}

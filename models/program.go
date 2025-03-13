package models

import (
	"time"
)

type Programs struct {
	ID                       string
	ChannelId                string
	TopicsID                 string
	Description              string
	MemoUrl                  string
	StartAt                  int64
	EndAt                    int64
	AllocatedAmount          int64
	AvailableAllocatedAmount int64
	EligibilityCheck         string
	Status                   int64
	CreatedAt, UpdatedAt     int64
	CreatedBy, UpdatedBy     string
	Created, Updated         time.Time
	SourceOfFund             string
	DiscountCalculation      string
	AllocatedQuota           int64
	AvailableAllocatedQuota  int64
	DiscountPercent          int64
	DiscountAmount           int64
	RefreshProgramQuotaDaily int64
	MerchantCsvUrl           string
	CustomerCsvUrl           string
	TopicTitle               string
}

func (us *Programs) UseUnixTimestamp() {
	us.CreatedAt = us.Created.Unix()
	us.UpdatedAt = us.Updated.Unix()
}

func (us *Programs) FastScan() []interface{} {
	return []interface{}{
		&us.ID,
		&us.ChannelId,
		&us.TopicsID,
		&us.Description,
		&us.MemoUrl,
		&us.StartAt,
		&us.EndAt,
		&us.AllocatedAmount,
		&us.AvailableAllocatedAmount,
		&us.EligibilityCheck,
		&us.Status,
		&us.Created,
		&us.CreatedBy,
		&us.Updated,
		&us.UpdatedBy,
		&us.SourceOfFund,
		&us.DiscountCalculation,
		&us.AllocatedQuota,
		&us.AvailableAllocatedQuota,
		&us.DiscountPercent,
		&us.DiscountAmount,
		&us.MerchantCsvUrl,
		&us.CustomerCsvUrl,
		&us.RefreshProgramQuotaDaily,
		&us.TopicTitle,
	}
}

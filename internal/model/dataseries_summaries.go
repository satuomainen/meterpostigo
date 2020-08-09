package model

import "time"

type DataSeriesSummaries struct {
	ID           uint64    `gorm:"column:id;primaryKey;not null;type:int(10);autoIncrement"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;autoCreateTime;default:current_timestamp()"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;autoCreateTime;default:current_timestamp()"`
	CurrentValue string    `gorm:"column:value;type:varchar(255);not null"`
	MinValue     string    `gorm:"column:value;type:varchar(255);not null"`
	MaxValue     string    `gorm:"column:value;type:varchar(255);not null"`
	DataSeriesID uint64    `gorm:"column:dataseries_id;type:int(10);not null;"`
}

func (t *DataSeriesSummaries) TableName() string {
	return "dataseries_summaries"
}

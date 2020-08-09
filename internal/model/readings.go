package model

import "time"

type Readings struct {
	ID           uint64     `gorm:"column:id;primaryKey;not null;type:int(10);autoIncrement"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null;autoCreateTime;default:current_timestamp()"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null;default:current_timestamp()"`
	Value        string     `gorm:"column:value;type:varchar(255);not null"`
	DataSeriesId uint64     `gorm:"column:dataseries_id;type:int(10);not null"`
}

func (Readings) TableName() string {
	return "readings"
}

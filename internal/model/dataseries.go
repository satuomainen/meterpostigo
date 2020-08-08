package model

import "time"

type DataSeries struct {
	ID          uint64    `gorm:"column:id;primaryKey;not null;type:int(10);autoIncrement"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime;default:current_timestamp()"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoCreateTime;default:current_timestamp()"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Description *string   `gorm:"column:description;type:text;default:NULL"`
	Label       *string   `gorm:"column:label;type:varchar(255);default:NULL"`
	ApiKey      string    `gorm:"column:api_key;type:varchar(255);not null"`
}

func (DataSeries) TableName() string {
	return "dataseries"
}

package service

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/model"
	"time"
)

func AddReading(dataSeriesId int64, value string) (createdReading *serverapi.Reading, err error) {
	reading := model.Readings{
		Value:        value,
		DataSeriesId: uint64(dataSeriesId),
	}

	createResult := db.DB.Create(&reading)

	updateSummaryWithLatestValue(dataSeriesId, &reading)

	dto := mapReadingToDTO(reading)
	return &dto, createResult.Error
}

func FindLatest(dataSeriesId int64, limit int) (*[]serverapi.Reading, error) {
	var readings []model.Readings

	queryStatus := db.DB.
		Model(&model.Readings{}).
		Where("dataseries_id = ?", dataSeriesId).
		Order("created_at DESC").
		Limit(limit).
		Find(&readings)

	if queryStatus.Error != nil {
		return nil, queryStatus.Error
	}

	var dtoList []serverapi.Reading
	for _, reading := range readings {
		dtoList = append(dtoList, mapReadingToDTO(reading))
	}

	return &dtoList, nil
}

type ReadingAverage struct {
	Date time.Time
	Value string
}

func FindAverages(dataSeriesId int64, days *int) (*[]serverapi.AverageReading, error) {
	limit := 30
	if days != nil {
		limit = *days
	}

	var averages []ReadingAverage

	queryStatus := db.DB.
		Model(&model.Readings{}).
		Select("DATE(created_at) AS Date, AVG(CAST(value AS DECIMAL(20,5))) AS Value").
		Where("dataseries_id = ?", dataSeriesId).
		Group("DATE(created_at)").
		Order("DATE(created_at) DESC").
		Limit(limit).
		Find(&averages)

	if queryStatus.Error != nil {
		return nil, queryStatus.Error
	}

	var dtoList []serverapi.AverageReading
	for _, average := range averages {
		dtoList = append(dtoList, mapAverageReadingToDTO(average))
	}

	return &dtoList, nil
}

package service

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/model"
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



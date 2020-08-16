package service

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/model"
	"github.com/labstack/gommon/log"
)

func FindSummaries() (*[]serverapi.DataSeriesSummary, error) {
	var summaries []model.DataSeriesSummaries
	queryStatus := db.DB.Preload("DataSeries").Find(&summaries)

	if queryStatus.Error != nil {
		return nil, queryStatus.Error
	}

	var dtoList []serverapi.DataSeriesSummary
	for _, summary := range summaries {
		dtoList = append(dtoList, mapDataSeriesSummaryToDTO(summary))
	}

	return &dtoList, nil
}

func updateSummaryWithLatestValue(dataSeriesId int64, reading *model.Readings) {
	var currentSummary model.DataSeriesSummaries
	queryResult := db.DB.
		Where("dataseries_id = ?", dataSeriesId).
		Find(&currentSummary)

	if queryResult.Error != nil {
		log.Errorf("Failed to update summary because summary was not found - %s", queryResult.Error)
		return
	}

	currentSummary.CurrentValue = reading.Value
	currentSummary.MinValue = min(reading.Value, currentSummary.MinValue)
	currentSummary.MaxValue = max(reading.Value, currentSummary.MaxValue)

	saveResult := db.DB.Save(&currentSummary)
	if saveResult.Error != nil {
		log.Errorf("Failed to update summary because saving updated values failed - %s", queryResult.Error)
		return
	}

	log.Info("Updated summary values")
}

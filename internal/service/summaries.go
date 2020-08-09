package service

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/model"
)

func FindSummaries() (*[]serverapi.DataSeriesSummary, error) {
	var summaries []model.DataSeriesSummaries
	queryStatus := db.DB.Find(&summaries)

	if queryStatus.Error != nil {
		return nil, queryStatus.Error
	}

	var dtoList []serverapi.DataSeriesSummary
	for _, summary := range summaries {
		dtoList = append(dtoList, mapDataSeriesSummaryToDTO(summary))
	}

	return &dtoList, nil
}



package service

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/model"
	"time"
)

func mapReadingToDTO(reading model.Readings) serverapi.Reading {
	return serverapi.Reading{
		CreatedAt: *toTimestampString(reading.CreatedAt),
		Id:        int64(reading.ID),
		UpdatedAt: toTimestampString(reading.UpdatedAt),
		Value:     reading.Value,
	}
}

func mapDataSeriesSummaryToDTO(model model.DataSeriesSummaries) serverapi.DataSeriesSummary {
	return serverapi.DataSeriesSummary{
		CreatedAt:    toTimestampString(model.CreatedAt),
		CurrentValue: &model.CurrentValue,
		DataSeriesId: int64(model.DataSeriesID),
		Id:           int64(model.ID),
		MaxValue:     &model.MaxValue,
		MinValue:     &model.MinValue,
		UpdatedAt:    toTimestampString(model.UpdatedAt),
	}
}

func toTimestampString(t time.Time) *string {
	timestamp := t.Format(time.RFC3339)
	return &timestamp
}

package service

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/model"
	"fmt"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"math"
	"strconv"
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


func mapAverageReadingToDTO(average ReadingAverage) serverapi.AverageReading {

	return serverapi.AverageReading{
		Date:  openapi_types.Date{
			Time: average.Date,
		},
		Value: average.Value,
	}
}

func mapDataSeriesSummaryToDTO(model model.DataSeriesSummaries) serverapi.DataSeriesSummary {
	return serverapi.DataSeriesSummary{
		CreatedAt:             toTimestampString(model.CreatedAt),
		CurrentValue:          toStringPointer(model.CurrentValue),
		DataSeriesId:          int64(model.DataSeriesID),
		DataSeriesDescription: model.DataSeries.Description,
		DataSeriesName:        toStringPointer(model.DataSeries.Name),
		DataSeriesLabel:       model.DataSeries.Label,
		Id:                    int64(model.ID),
		MaxValue:              toStringPointer(model.MaxValue),
		MinValue:              toStringPointer(model.MinValue),
		UpdatedAt:             toTimestampString(model.UpdatedAt),
	}
}

func toTimestampString(t time.Time) *string {
	timestamp := t.Format(time.RFC3339)
	return &timestamp
}

func min(a string, b string) string {
	aNumeric, aErr := strconv.ParseFloat(a, 64)
	bNumeric, bErr := strconv.ParseFloat(b, 64)

	if aErr != nil || bErr != nil {
		if a < b {
			return a
		}
		return b
	}

	return fmt.Sprintf("%.2f", math.Min(aNumeric, bNumeric))
}

func max(a string, b string) string {
	aNumeric, aErr := strconv.ParseFloat(a, 64)
	bNumeric, bErr := strconv.ParseFloat(b, 64)

	if aErr != nil || bErr != nil {
		if a > b {
			return a
		}
		return b
	}

	return fmt.Sprintf("%.2f", math.Max(aNumeric, bNumeric))
}

func toStringPointer(s string) *string {
	return &s
}

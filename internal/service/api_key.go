package service

import (
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckApiKey(dataSeriesId int64, apiKey string) error {
	var dataSeries []model.DataSeries
	queryResult := db.DB.
		Model(&model.DataSeries{}).
		Where("id = ?", dataSeriesId).
		Where("api_key = ?", apiKey).
		Find(&dataSeries)

	if queryResult.Error != nil || queryResult.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return nil
}

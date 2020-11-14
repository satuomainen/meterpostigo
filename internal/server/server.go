package server

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

const defaultReadingsLimit = 100

type MetricsServer struct {
}

// GetDataseriesSummaries serves (GET /dataseries/summaries)
func (MetricsServer) GetDataseriesSummaries(ctx echo.Context) error {

	summaries, err := service.FindSummaries()

	if err != nil {
		log.Errorf("Failed to fetch dataseries summaries - %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, summaries)
}

// GetDataseriesSummariesDataSeriesId serves (GET /dataseries/summaries/{dataSeriesId})
func (MetricsServer) GetDataseriesSummariesDataSeriesId(ctx echo.Context, dataSeriesId int64) error {
	summary, err := service.FindSummary(dataSeriesId)

	if err != nil {
		log.Errorf("Failed to fetch dataseries summary for series %d - %s", dataSeriesId, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, summary)
}

// GetDataseriesDataSeriesIdReadings serves (GET /dataseries/{dataSeriesId}/readings)
func (MetricsServer) GetDataseriesDataSeriesIdReadings(
	ctx echo.Context,
	dataSeriesId int64,
	params serverapi.GetDataseriesDataSeriesIdReadingsParams,
) error {
	limit := defaultReadingsLimit
	if params.Limit != nil {
		limit = *params.Limit
	}

	readings, err := service.FindLatest(dataSeriesId, limit)

	if err != nil {
		log.Errorf("Failed to fetch readings for data series '%d' - %s", dataSeriesId, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, readings)
}

func (MetricsServer) GetDataseriesAveragesByDataSeriesId(
	ctx echo.Context,
	dataSeriesId int64,
	params serverapi.GetDataseriesAveragesByDataSeriesIdParams,
) error {
	if averages, err := service.FindAverages(dataSeriesId, params.Days); err != nil {
		log.Errorf("Failed to fetch averages for data series '%d' for %d days - %s", dataSeriesId, params.Days, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	} else {
		return ctx.JSON(http.StatusOK, averages)
	}
}

// PostDataseriesDataSeriesIdReadings serves (POST /dataseries/{dataSeriesId}/readings)
func (MetricsServer) PostDataseriesDataSeriesIdReadings(ctx echo.Context, dataSeriesId int64) error {
	req := serverapi.NewReading{}
	if err := ctx.Bind(&req); err != nil {
		ctx.Logger().Warnf("Failed to bind add reading request - %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to extract reading from request")
	}

	return saveNewValue(ctx, dataSeriesId, req.ApiKey, req.Value)
}

// PostSeriesDataSeriesIdAdd serves (POST /series/{dataSeriesId}/add)
func (MetricsServer) PostSeriesDataSeriesIdAdd(ctx echo.Context, dataSeriesId int64) error {
	value := ctx.FormValue("value")
	apiKey := ctx.FormValue("api_key")

	return saveNewValue(ctx, dataSeriesId, apiKey, value)
}

func saveNewValue(ctx echo.Context, dataSeriesId int64, apiKey string, value string) error {
	if err := service.CheckApiKey(dataSeriesId, apiKey); err != nil {
		log.Warnf("Unauthorized add request to series '%d' with key '%s'", dataSeriesId, apiKey)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	addedReading, err := service.AddReading(dataSeriesId, value)
	if err != nil {
		log.Errorf("Failed to add reading - %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Debugf("New value added to series %d", dataSeriesId)

	return ctx.JSON(http.StatusCreated, addedReading)
}

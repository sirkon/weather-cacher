package geo

import (
	"context"
)

// Geo работа с геграфическими координатами
type Geo interface {
	// GetNearby получение идентификаторов уже полученного от данного провайдера прогноза вблизи данного положения
	// возвращает идентификатор прогноза
	GetNearby(ctx context.Context, provID string, lat, lon float64) (map[string]float64, error)

	// Set привязка прогноза с данным идентификатором к данной точке
	Set(ctx context.Context, provID string, lat, lon float64, forecastID string) error
}

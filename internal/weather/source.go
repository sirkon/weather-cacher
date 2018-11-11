package weather

import (
	"context"
)

// Source общее описание источника данных
type Source interface {
	WeatherFor(ctx context.Context, lat, lon float64) (*Forecast, error)
}

package weather

import (
	"context"
	"github.com/sirkon/weather-cacher/internal/schema"
)

// Source общее описание источника данных
type Source interface {
	WeatherFor(ctx context.Context, lat, lon float64) (*schema.Forecast, error)
}

package cache

import (
	"context"
	"github.com/sirkon/weather-cacher/internal/schema"
)

// Cache абстракция для хранения и получения кешированных данных
type Cache interface {
	// Get получение прогноза с данным идентификатором. Возвращает (nil, nil), если данный прогноз уже протух
	Get(ctx context.Context, forecastID string) (*schema.Forecast, error)

	// Set сохранение прогноза с данным идентификатором
	Set(ctx context.Context, forecastID string, forecast *schema.Forecast) error
}

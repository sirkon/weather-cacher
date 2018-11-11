package weather

import (
	"context"
)

// Cache абстракция для хранения и получения кешированных данных
type Cache interface {
	// Get получение прогноза с данным идентификатором. Возвращает (nil, nil), если данный прогноз уже протух
	Get(ctx context.Context, forecastID string) (*Forecast, error)

	// Set сохранение прогноза с данным идентификатором
	Set(ctx context.Context, forecastID string, forecast *Forecast) error
}

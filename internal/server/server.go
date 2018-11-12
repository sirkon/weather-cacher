package server

import (
	"context"
	"fmt"
	"github.com/sirkon/weather-cacher/internal/cache"
	"github.com/sirkon/weather-cacher/internal/geo"
	"github.com/sirkon/weather-cacher/internal/idgen"
	"github.com/sirkon/weather-cacher/internal/schema"
	"github.com/sirkon/weather-cacher/internal/weather"
	"google.golang.org/grpc/metadata"
	"log"
	"math"
	"time"
)

const userAuthToken = "x-user-token"

// Weather конструктор реализации schema.WeatherServer
func Weather(providers map[string]weather.Source, cache cache.Cache, geo geo.Geo, idGen idgen.IDGen) schema.WeatherServer {
	return &server{
		providers: providers,
		cache:     cache,
		geo:       geo,
		idGen:     idGen,
	}
}

type server struct {
	providers map[string]weather.Source
	cache     cache.Cache
	geo       geo.Geo
	idGen     idgen.IDGen
}

func errorResult(msg string) *schema.WeatherResponse {
	return &schema.WeatherResponse{
		Result: &schema.WeatherResponse_Error_{
			Error: &schema.WeatherResponse_Error{
				Msg: "wrong request",
			},
		},
	}
}

// Get ...
func (s *server) Get(ctx context.Context, req *schema.WeatherRequest) (*schema.WeatherResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || !s.validateRequest(md, req) {
		return errorResult("bad request"), nil
	}

	prov, ok := s.providers[req.ProviderId]
	if !ok {
		return errorResult(fmt.Sprintf("unknown provider %s", req.ProviderId)), nil
	}

	// ищем предсказания для ближних точек
	forecasts, err := s.geo.GetNearby(ctx, req.ProviderId, req.Latitude, req.Longitude)
	if err != nil {
		return errorResult(err.Error()), nil
	}

	// если нашли, то пытаемся получить данные из кеша
	if len(forecasts) != 0 {
		for key := range forecasts {
			forecast, err := s.cache.Get(ctx, key)
			if err != nil {
				log.Printf("failed to get forecast with ID %s", key)
				continue
			}
			if forecast != nil {
				return &schema.WeatherResponse{
					Result: &schema.WeatherResponse_Forecast{
						Forecast: forecast,
					},
				}, nil
			}
		}
	}

	// данные в кеше не найдены, делаем запрос к поставщику
	forecast, err := prov.WeatherFor(ctx, req.Latitude, req.Longitude)
	if err != nil {
		return errorResult(err.Error()), nil
	}

	// сохраняем полученный прогноз, по-возможности, в кеше, делаем это в фоне
	go s.storeForecast(req, forecast)

	return &schema.WeatherResponse{
		Result: &schema.WeatherResponse_Forecast{
			Forecast: forecast,
		},
	}, nil
}

func (s *server) storeForecast(req *schema.WeatherRequest, forecast *schema.Forecast) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	forecastID, err := s.idGen.ID(req.ProviderId, req.Latitude, req.Longitude, forecast)
	if err != nil {
		log.Println("failed to compute forecast id", err)
		return
	}
	if err := s.cache.Set(ctx, forecastID, forecast); err != nil {
		log.Println("failed to write forecast into the cache", err)
		return
	}
	if err := s.geo.Set(ctx, req.ProviderId, req.Latitude, req.Longitude, forecastID); err != nil {
		log.Println("failed to register forecast into the DB", err)
	}
}

func (s *server) validateRequest(md metadata.MD, req *schema.WeatherRequest) bool {
	if len(req.UserId) == 0 || len(md.Get(userAuthToken)) == 0 {
		// дальше здесь по-хорошему проверяем токен для данного пользователя, но пока для этого ничего ещё не сделано
		return false
	}
	if math.Abs(req.Latitude) > 90 || math.Abs(req.Longitude) > 180 {
		return false
	}
	return true
}

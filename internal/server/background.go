package server

import (
	"context"
	"github.com/sirkon/weather-cacher/internal/schema"
	"log"
	"time"
)

// Background реализация
type Background interface {
	Job()
	Stop() <-chan struct{}
}

// Job запуск наполнителя кешей
func (s *server) Job() {
	s.wg.Add(1)
	go func() {
	again:
		select {
		case task := <-s.tasks:
			s.storeForecast(task.provID, task.lat, task.lon, task.forecat)
		case <-s.stop:
			s.wg.Done()
			return
		}
		goto again // не нравится отступ цикла
	}()
}

// Stop жобе остановиться
func (s *server) Stop() <-chan struct{} {
	close(s.stop)
	res := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(res)
	}()
	return res
}

func (s *server) storeForecast(provID string, lat, lon float64, forecast *schema.Forecast) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	forecastID, err := s.idGen.ID(provID, lat, lon, forecast)
	if err != nil {
		log.Println("failed to compute forecast id", err)
		return
	}
	if err := s.cache.Set(ctx, forecastID, forecast); err != nil {
		log.Println("failed to write forecast into the cache", err)
		return
	}
	if err := s.geo.Set(ctx, provID, lat, lon, forecastID); err != nil {
		log.Println("failed to register forecast into the DB", err)
	}
}

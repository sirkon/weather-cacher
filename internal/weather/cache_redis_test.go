// +build sirkon

package weather

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCacheRedis(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong := c.Ping()
	if err := pong.Err(); err != nil {
		t.Fatal(err)
	}

	cr := CacheRedis(c)

	f := &Forecast{
		Current: Hourly{
			Time:    0,
			Summary: "forecast",
			Temperature: struct {
				Measured float64 `json:"measured,omitempty"`
				Apparent float64 `json:"apparent,omitempty"`
				DewPoint float64 `json:"dew_point,omitempty"`
			}{},
			Precipitation: struct {
				Type        string  `json:"type,omitempty"`
				Intensity   float64 `json:"intensity,omitempty"`
				Probability float64 `json:"probability,omitempty"`
			}{},
			Humidity: 0,
			Pressure: 0,
			Wind: struct {
				Speed   float64 `json:"speed,omitempty"`
				Bearing int64   `json:"bearing,omitempty"`
				Gust    float64 `json:"gust,omitempty"`
			}{},
			CloudCover: 0,
			UVIndex:    0,
			Visibility: 0,
		},
		Hourly: nil,
		Daily:  nil,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	const forecastID = "forecast-id"
	if err := cr.Set(ctx, forecastID, f); err != nil {
		t.Fatal(err)
	}

	res, err := cr.Get(ctx, forecastID)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, f, res)
}

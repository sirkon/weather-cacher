// +build sirkon

/*
Ставим redis как здесь:	https://hub.docker.com/_/redis/

И запускаем его:

	sudo docker run --name some-redis -d redis
*/

package cache

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirkon/weather-cacher/internal/schema"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong := c.Ping()
	if err := pong.Err(); err != nil {
		t.Fatal(err)
	}

	cr := Redis(c)

	f := &schema.Forecast{
		Current: &schema.Hourly{
			Summary: "forecast",
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

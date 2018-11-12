package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirkon/weather-cacher/internal/cache"
	"github.com/sirkon/weather-cacher/internal/geo"
	"github.com/sirkon/weather-cacher/internal/idgen"
	"github.com/sirkon/weather-cacher/internal/schema"
	"github.com/sirkon/weather-cacher/internal/server"
	"github.com/sirkon/weather-cacher/internal/weather"
	"github.com/sirkon/weather-cacher/internal/weather/darksky"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/urfave/cli.v2"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func startServer(c *cli.Context) error {
	darkSkyToken := c.String("token-dark-sky")
	pgAddr := c.String("postgres")
	redisAddr := c.String("redis")
	cachingJobsNo := c.Uint("caching-jobs")
	port := c.Uint("port")

	pgDB, err := sql.Open("postgres", "postgres://"+pgAddr)
	if err != nil {
		return fmt.Errorf("failed to initiate postgres connection")
	}
	g := geo.Postgis(pgDB)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	cach := cache.Redis(redisClient)

	dsrc := darksky.RawClient(darkSkyToken, &http.Client{})
	dsc := darksky.Client(dsrc)

	srv, bgjob := server.Weather(
		map[string]weather.Source{
			"dark-sky": dsc,
		},
		cach,
		g,
		idgen.New(md5.New()),
	)

	for i := uint(0); i < cachingJobsNo; i++ {
		bgjob.Job()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt)
	go func() {
		for range osSig {
			log.Println("stopping server")
			grpcServer.GracefulStop()
			<-bgjob.Stop()
			log.Println("stopped")
		}
	}()

	schema.RegisterWeatherServer(grpcServer, srv)
	reflection.Register(grpcServer)
	log.Println("start serving")
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	log.Println("server stopped")

	return nil
}

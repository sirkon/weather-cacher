package main

import (
	"gopkg.in/urfave/cli.v2"
	"log"
	"os"
)

func main() {
	app := cli.App{
		Name: "weather-cacher",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:    "port",
				Usage:   "port to serve",
				Value:   50005,
				EnvVars: []string{"SERVER_PORT"},
			},
			&cli.StringFlag{
				Name:    "postgres",
				Usage:   "postgres connection parameters in `user:pass@host:port/database` form",
				Value:   "docker:docker@localhost:25432/data",
				EnvVars: []string{"POSTGRES_CONN"},
			},
			&cli.StringFlag{
				Name:    "redis",
				Usage:   "redis connection parameters in `host:port` form",
				Value:   "localhost:6379",
				EnvVars: []string{"REDIS_CONN"},
			},
			&cli.StringFlag{
				Name:    "token-dark-sky",
				Usage:   "token for dark sky service",
				EnvVars: []string{"TOKEN_DARK_SKY"},
			},
			&cli.UintFlag{
				Name:    "caching-jobs",
				Usage:   "amount of jobs for forecast caching. 0 means no job",
				EnvVars: []string{"CACHING_JOBS_NO"},
				Value:   1,
			},
		},
		Action: startServer,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

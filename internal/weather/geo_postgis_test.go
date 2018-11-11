// +build sirkon

/*
Используем этот docker-образ: https://hub.docker.com/r/kartoza/postgis/
После установки запускаем образ и создаём БД data:

    sudo docker run --name "postgis" -p 25432:5432 -d -t kartoza/postgis

    psql -h localhost -U docker -p 25432 -d gis
    CREATE DATABASE data;

    psql -h localhost -U docker -p 25432 -d gis
    CREATE TABLE public.forecast
    (
      id character varying NOT NULL,
      provider_id character varying NOT NULL,
      created timestamp without time zone NOT NULL,
      location geography,
      CONSTRAINT forecast_pkey PRIMARY KEY (id, provider_id)
    );
*/

package weather

import (
	"context"
	"database/sql"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGeoPostgis(t *testing.T) {
	connStr := "postgres://docker:docker@localhost:25432/data"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	type (
		node struct {
			id  string
			lat float64
			lon float64
		}
	)

	data := []node{
		{
			id:  "moscow-1",
			lat: 55.751944,
			lon: 37.615555,
		},
		{
			id:  "moscow-2",
			lat: 55.752,
			lon: 37.61555,
		},
		{
			id:  "leningrad-1",
			lat: 59.9386111,
			lon: 30.313888,
		},
		{
			id:  "leningrad-2",
			lat: 59.94,
			lon: 30.31389,
		},
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	_, err = db.ExecContext(ctx,
		`DELETE FROM forecast WHERE id IN ($1, $2, $3, $4)`,
		data[0].id, data[1].id, data[2].id, data[3].id,
	)
	if err != nil {
		t.Fatal(err)
	}

	g := GeoPostgis(db)

	const providerID = "prov-id"

	if err := g.Set(ctx, providerID, data[0].lat, data[0].lon, data[0].id); err != nil {
		t.Fatal(err)
	}
	if err := g.Set(ctx, providerID, data[2].lat, data[2].lon, data[2].id); err != nil {
		t.Fatal(err)
	}

	resp, err := g.GetNearby(ctx, providerID, data[1].lat, data[1].lon)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := resp[data[0].id]; !ok {
		t.Fatalf("there should be %s key in the response", data[0].id)
	}
	require.Len(t, resp, 1)
	t.Log(litter.Sdump(resp))

	resp, err = g.GetNearby(ctx, providerID, data[3].lat, data[3].lon)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := resp[data[2].id]; !ok {
		t.Fatalf("there should be %s key in the response", data[2].id)
	}
	require.Len(t, resp, 1)
	t.Log(litter.Sdump(resp))
}

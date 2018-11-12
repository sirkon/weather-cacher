// +build sirkon

/*
Используем этот docker-образ: https://hub.docker.com/r/kartoza/postgis/
После установки запускаем образ и создаём БД data:

    .

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

package geo

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

	moscow1 := node{
		id:  "moscow-1",
		lat: 55.751944,
		lon: 37.615555,
	}
	moscow2 := node{
		id:  "moscow-2",
		lat: 55.752,
		lon: 37.61555,
	}
	leningrad1 := node{
		id:  "leningrad-1",
		lat: 59.9386111,
		lon: 30.313888,
	}
	leningrad2 := node{
		id:  "leningrad-2",
		lat: 59.94,
		lon: 30.31389,
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	_, err = db.ExecContext(ctx,
		`DELETE FROM forecast WHERE id IN ($1, $2)`,
		moscow1.id, leningrad1.id,
	)
	if err != nil {
		t.Fatal(err)
	}

	g := Postgis(db)

	const providerID = "prov-id"

	if err := g.Set(ctx, providerID, moscow1.lat, moscow1.lon, moscow1.id); err != nil {
		t.Fatal(err)
	}
	if err := g.Set(ctx, providerID, leningrad1.lat, leningrad1.lon, leningrad1.id); err != nil {
		t.Fatal(err)
	}

	resp, err := g.GetNearby(ctx, providerID, moscow2.lat, moscow2.lon)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := resp[moscow1.id]; !ok {
		t.Fatalf("there should be %s key in the response", moscow1.id)
	}
	require.Len(t, resp, 1)
	t.Log(litter.Sdump(resp))

	resp, err = g.GetNearby(ctx, providerID, leningrad2.lat, leningrad2.lon)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := resp[leningrad1.id]; !ok {
		t.Fatalf("there should be %s key in the response", leningrad1.id)
	}
	require.Len(t, resp, 1)
	t.Log(litter.Sdump(resp))
}

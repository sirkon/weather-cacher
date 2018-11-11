package weather

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// GeoPostgis конструктор Geo работающего поверх postgis
func GeoPostgis(conn *sql.DB) Geo {
	return &geoPostgis{
		conn: conn,
	}
}

type geoPostgis struct {
	conn *sql.DB
}

func (pg *geoPostgis) GetNearby(ctx context.Context, provID string, lat, lon float64) (map[string]float64, error) {
	fifteenMinutesBefore := time.Now().Add(-time.Minute * 15)
	rows, err := pg.conn.QueryContext(ctx, `
SELECT * FROM (
	SELECT
		id,
		ST_Distance(ST_SetSRID(ST_Point($3, $4), 4326)::geography, location) AS dist
	FROM data.forecast
  	WHERE provider = $2 AND created > $1
) AS subquery WHERE dist < 5000
`, fifteenMinutesBefore, provID, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve nearby forecasts: %s", err)
	}
	var id string
	var dist float64
	res := map[string]float64{}
	for rows.Next() {
		if err := rows.Scan(&id, &dist); err != nil {
			return nil, fmt.Errorf("failed to scan db response: %s", err)
		}
		res[id] = dist
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan db response: %s", err)
	}
	rows.NextResultSet()
	return res, nil
}

func (pg *geoPostgis) Set(ctx context.Context, provID string, lat, lon float64, forecastID string) error {
	_, err := pg.conn.ExecContext(ctx, `
INSERT INTO data.forecast (id, created, provider, location) VALUES ($1, $2, $3, ST_SetSRID(ST_Point($4, $5), 4326)::geography)
`, forecastID, time.Now(), provID, lat, lon)
	if err != nil {
		return fmt.Errorf("failed to insert geo data: %s", err)
	}
	return nil
}

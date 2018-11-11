package darksky

import (
	"fmt"
	"github.com/sirkon/weather-cacher/internal/rawclient"
	"net/http"
)

// RawClient сырой клиент для Dark Sky
func RawClient(token string, client *http.Client) rawclient.RawClient {
	return rawclient.GetURLTokenRawClient(func(lat, lon float64) string {
		return fmt.Sprintf("https://api.darksky.net/forecast/%s/%f,%f?units=si", token, lat, lon)
	}, client)
}

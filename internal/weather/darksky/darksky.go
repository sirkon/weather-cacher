package darksky

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirkon/weather-cacher/internal/rawclient"
	"github.com/sirkon/weather-cacher/internal/weather"
)

// Client клиент DarkSky
func Client(rawClient rawclient.RawClient) weather.Source {
	return client{
		rawClient: rawClient,
	}
}

type client struct {
	rawClient rawclient.RawClient
}

// WeatherFor ...
func (c client) WeatherFor(ctx context.Context, lat, lon float64) (*weather.Forecast, error) {
	r, err := c.rawClient.Get(ctx, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from Dark Sky: %s", err)
	}

	defer r.Close()

	var resp Response
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed to decode data from Dark Sky: %s", err)
	}

	return transformResponse(resp), nil
}

func transformResponse(r Response) *weather.Forecast {
	cur := r.Currently
	res := &weather.Forecast{
		Current: weather.Hourly{
			Time:    cur.Time,
			Summary: cur.Summary,
			Temperature: struct {
				Measured float64 `json:"measured,omitempty"`
				Apparent float64 `json:"apparent,omitempty"`
				DewPoint float64 `json:"dew_point,omitempty"`
			}{
				Measured: cur.Temperature,
				Apparent: cur.ApparentTemperature,
				DewPoint: cur.DewPoint,
			},
			Precipitation: struct {
				Type        string  `json:"type,omitempty"`
				Intensity   float64 `json:"intensity,omitempty"`
				Probability float64 `json:"probability,omitempty"`
			}{
				Type:        cur.PrecipType,
				Intensity:   cur.PrecipIntensity,
				Probability: cur.PrecipProbability,
			},
			Humidity: cur.Humidity,
			Pressure: cur.Pressure,
			Wind: struct {
				Speed   float64 `json:"speed,omitempty"`
				Bearing int64   `json:"bearing,omitempty"`
				Gust    float64 `json:"gust,omitempty"`
			}{
				Speed:   cur.WindSpeed,
				Bearing: cur.WindBearing,
				Gust:    cur.WindGust,
			},
			CloudCover: cur.CloudCover,
			UVIndex:    uint(cur.UvIndex),
			Visibility: cur.Visibility,
		},
	}

	res.Daily = make([]weather.Daily, len(r.Daily.Data))
	for i, day := range r.Daily.Data {
		f := weather.Daily{
			Time:    day.Time,
			Summary: day.Summary,
			Sun: struct {
				Rise int64 `json:"rise,omitempty"`
				Set  int64 `json:"set,omitempty"`
			}{
				Rise: day.SunriseTime,
				Set:  day.SunsetTime,
			},
			MoonPhase: day.MoonPhase,
			Temperature: struct {
				Measured weather.TemperatureMeasurement `json:"measured,omitempty"`
				Apparent weather.TemperatureMeasurement `json:"apparent,omitempty"`
				DewPoint float64                        `json:"dew_point,omitempty"`
			}{
				Measured: weather.TemperatureMeasurement{
					Min: weather.Measurement{
						Value: day.TemperatureMin,
						Time:  day.TemperatureMinTime,
					},
					Low: weather.Measurement{
						Value: day.TemperatureLow,
						Time:  day.TemperatureLowTime,
					},
					High: weather.Measurement{
						Value: day.TemperatureHigh,
						Time:  day.TemperatureHighTime,
					},
					Max: weather.Measurement{
						Value: day.TemperatureMax,
						Time:  day.TemperatureMaxTime,
					},
				},
				Apparent: weather.TemperatureMeasurement{
					Min: weather.Measurement{
						Value: day.ApparentTemperatureMin,
						Time:  day.ApparentTemperatureMinTime,
					},
					Low: weather.Measurement{
						Value: day.ApparentTemperatureLow,
						Time:  day.ApparentTemperatureLowTime,
					},
					High: weather.Measurement{
						Value: day.ApparentTemperatureHigh,
						Time:  day.ApparentTemperatureHighTime,
					},
					Max: weather.Measurement{
						Value: day.ApparentTemperatureMax,
						Time:  day.ApparentTemperatureMaxTime,
					},
				},
				DewPoint: day.DewPoint,
			},
			Precipitation: struct {
				Type             string  `json:"type,omitempty"`
				Intensity        float64 `json:"intensity,omitempty"`
				MaxIntensity     float64 `json:"max_intensity,omitempty"`
				MaxIntensityTime int64   `json:"max_intensity_time,omitempty"`
				Probability      float64 `json:"probability,omitempty"`
			}{
				Type:             day.PrecipType,
				Intensity:        day.PrecipIntensity,
				MaxIntensity:     day.PrecipIntensityMax,
				MaxIntensityTime: day.PrecipIntensityMaxTime,
				Probability:      day.PrecipProbability,
			},
			Humidity: day.Humidity,
			Pressure: day.Pressure,
			Wind: struct {
				Speed    float64 `json:"speed,omitempty"`
				Bearing  int64   `json:"bearing,omitempty"`
				Gust     float64 `json:"gust,omitempty"`
				GustTime int64   `json:"gust_time,omitempty"`
			}{
				Speed:    day.WindSpeed,
				Bearing:  day.WindBearing,
				Gust:     day.WindGust,
				GustTime: day.WindGustTime,
			},
			CloudCover: day.CloudCover,
			UV: struct {
				Index uint  `json:"index,omitempty"`
				Time  int64 `json:"time,omitempty"`
			}{
				Index: uint(day.UvIndex),
				Time:  day.UvIndexTime,
			},
			Visibility: day.Visibility,
		}
		res.Daily[i] = f
	}

	res.Hourly = make([]weather.Hourly, len(r.Hourly.Data))
	for i, hour := range r.Hourly.Data {
		h := weather.Hourly{
			Time:    hour.Time,
			Summary: hour.Summary,
			Temperature: struct {
				Measured float64 `json:"measured,omitempty"`
				Apparent float64 `json:"apparent,omitempty"`
				DewPoint float64 `json:"dew_point,omitempty"`
			}{
				Measured: hour.Temperature,
				Apparent: hour.ApparentTemperature,
				DewPoint: hour.DewPoint,
			},
			Precipitation: struct {
				Type        string  `json:"type,omitempty"`
				Intensity   float64 `json:"intensity,omitempty"`
				Probability float64 `json:"probability,omitempty"`
			}{
				Type:        hour.PrecipType,
				Intensity:   hour.PrecipIntensity,
				Probability: hour.PrecipProbability,
			},
			Humidity: hour.Humidity,
			Pressure: hour.Pressure,
			Wind: struct {
				Speed   float64 `json:"speed,omitempty"`
				Bearing int64   `json:"bearing,omitempty"`
				Gust    float64 `json:"gust,omitempty"`
			}{
				Speed:   hour.WindSpeed,
				Bearing: hour.WindBearing,
				Gust:    hour.WindGust,
			},
			CloudCover: hour.CloudCover,
			UVIndex:    uint(hour.UvIndex),
			Visibility: hour.Visibility,
		}
		res.Hourly[i] = h
	}

	return res
}

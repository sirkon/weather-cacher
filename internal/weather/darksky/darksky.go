package darksky

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirkon/weather-cacher/internal/rawclient"
	"github.com/sirkon/weather-cacher/internal/schema"
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
func (c client) WeatherFor(ctx context.Context, lat, lon float64) (*schema.Forecast, error) {
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

func transformResponse(r Response) *schema.Forecast {
	cur := r.Currently
	res := &schema.Forecast{
		Current: &schema.Hourly{
			Time:    cur.Time,
			Summary: cur.Summary,
			Temperature: &schema.Hourly_Temperature{
				Measured: cur.Temperature,
				Apparent: cur.ApparentTemperature,
				DewPoint: cur.DewPoint,
			},
			Precipation: &schema.Hourly_Precipitation{
				Type:        cur.PrecipType,
				Intensity:   cur.PrecipIntensity,
				Probability: cur.PrecipProbability,
			},
			Humidity: cur.Humidity,
			Pressure: cur.Pressure,
			Wind: &schema.Hourly_Wind{
				Speed:   cur.WindSpeed,
				Bearing: cur.WindBearing,
				Gust:    cur.WindGust,
			},
			CloudCover: cur.CloudCover,
			UvIndex:    int32(cur.UvIndex),
			Visibility: cur.Visibility,
		},
	}

	res.Daily = make([]*schema.Daily, len(r.Daily.Data))
	for i, day := range r.Daily.Data {
		f := &schema.Daily{
			Time:    day.Time,
			Summary: day.Summary,
			Sun: &schema.Daily_Sun{
				Rise: day.SunriseTime,
				Set:  day.SunsetTime,
			},
			MoonPhase: day.MoonPhase,
			Temperature: &schema.Daily_Temperature{
				Measured: &schema.Daily_Temperature_Measurement{
					Min: &schema.Daily_Temperature_Measurement_Value{
						Value: day.TemperatureMin,
						Time:  day.TemperatureMinTime,
					},
					Low: &schema.Daily_Temperature_Measurement_Value{
						Value: day.TemperatureLow,
						Time:  day.TemperatureLowTime,
					},
					High: &schema.Daily_Temperature_Measurement_Value{
						Value: day.TemperatureHigh,
						Time:  day.TemperatureHighTime,
					},
					Max: &schema.Daily_Temperature_Measurement_Value{
						Value: day.TemperatureMax,
						Time:  day.TemperatureMaxTime,
					},
				},
				Apparent: &schema.Daily_Temperature_Measurement{
					Min: &schema.Daily_Temperature_Measurement_Value{
						Value: day.ApparentTemperatureMin,
						Time:  day.ApparentTemperatureMinTime,
					},
					Low: &schema.Daily_Temperature_Measurement_Value{
						Value: day.ApparentTemperatureLow,
						Time:  day.ApparentTemperatureLowTime,
					},
					High: &schema.Daily_Temperature_Measurement_Value{
						Value: day.ApparentTemperatureHigh,
						Time:  day.ApparentTemperatureHighTime,
					},
					Max: &schema.Daily_Temperature_Measurement_Value{
						Value: day.ApparentTemperatureMax,
						Time:  day.ApparentTemperatureMaxTime,
					},
				},
				DewPoint: day.DewPoint,
			},
			Precipation: &schema.Daily_Precipitation{
				Type:             day.PrecipType,
				Intensity:        day.PrecipIntensity,
				MaxIntensity:     day.PrecipIntensityMax,
				MaxIntensityTime: day.PrecipIntensityMaxTime,
				Probability:      day.PrecipProbability,
			},
			Humidity: day.Humidity,
			Pressure: day.Pressure,
			Wind: &schema.Daily_Wind{
				Speed:    day.WindSpeed,
				Bearing:  day.WindBearing,
				Gust:     day.WindGust,
				GustTime: day.WindGustTime,
			},
			CloudCover: day.CloudCover,
			Uv: &schema.Daily_UVValue{
				Index: int32(day.UvIndex),
				Time:  day.UvIndexTime,
			},
			Visibility: day.Visibility,
		}
		res.Daily[i] = f
	}

	res.Hourly = make([]*schema.Hourly, len(r.Hourly.Data))
	for i, hour := range r.Hourly.Data {
		h := &schema.Hourly{
			Time:    hour.Time,
			Summary: hour.Summary,
			Temperature: &schema.Hourly_Temperature{
				Measured: hour.Temperature,
				Apparent: hour.ApparentTemperature,
				DewPoint: hour.DewPoint,
			},
			Precipation: &schema.Hourly_Precipitation{
				Type:        hour.PrecipType,
				Intensity:   hour.PrecipIntensity,
				Probability: hour.PrecipProbability,
			},
			Humidity: hour.Humidity,
			Pressure: hour.Pressure,
			Wind: &schema.Hourly_Wind{
				Speed:   hour.WindSpeed,
				Bearing: hour.WindBearing,
				Gust:    hour.WindGust,
			},
			CloudCover: hour.CloudCover,
			UvIndex:    int32(hour.UvIndex),
			Visibility: hour.Visibility,
		}
		res.Hourly[i] = h
	}

	return res
}

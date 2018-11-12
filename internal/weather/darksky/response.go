package darksky

// Response формат ответа от DarkSky
type Response struct {
	Currently struct {
		ApparentTemperature  float64 `json:"apparentTemperature"`
		CloudCover           float64 `json:"cloudCover"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Icon                 string  `json:"icon"`
		NearestStormBearing  int64   `json:"nearestStormBearing"`
		NearestStormDistance int64   `json:"nearestStormDistance"`
		Ozone                float64 `json:"ozone"`
		PrecipIntensity      float64 `json:"precipIntensity"`
		PrecipProbability    float64 `json:"precipProbability"`
		PrecipType           string  `json:"precipType"`
		Pressure             float64 `json:"pressure"`
		Summary              string  `json:"summary"`
		Temperature          float64 `json:"temperature"`
		Time                 int64   `json:"time"`
		UvIndex              int64   `json:"uvIndex"`
		Visibility           float64 `json:"visibility"`
		WindBearing          int64   `json:"windBearing"`
		WindGust             float64 `json:"windGust"`
		WindSpeed            float64 `json:"windSpeed"`
	} `json:"currently"`
	Daily struct {
		Data []struct {
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime"`
			CloudCover                  float64 `json:"cloudCover"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Icon                        string  `json:"icon"`
			MoonPhase                   float64 `json:"moonPhase"`
			Ozone                       float64 `json:"ozone"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipType                  string  `json:"precipType"`
			Pressure                    float64 `json:"pressure"`
			Summary                     string  `json:"summary"`
			SunriseTime                 int64   `json:"sunriseTime"`
			SunsetTime                  int64   `json:"sunsetTime"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int64   `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int64   `json:"temperatureLowTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int64   `json:"temperatureMaxTime"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int64   `json:"temperatureMinTime"`
			Time                        int64   `json:"time"`
			UvIndex                     int64   `json:"uvIndex"`
			UvIndexTime                 int64   `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			WindBearing                 int64   `json:"windBearing"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int64   `json:"windGustTime"`
			WindSpeed                   float64 `json:"windSpeed"`
		} `json:"data"`
		Icon    string `json:"icon"`
		Summary string `json:"summary"`
	} `json:"daily"`
	Flags struct {
		Nearest_station float64  `json:"nearest-station"`
		Sources         []string `json:"sources"`
		Units           string   `json:"units"`
	} `json:"flags"`
	Hourly struct {
		Data []struct {
			ApparentTemperature float64 `json:"apparentTemperature"`
			CloudCover          float64 `json:"cloudCover"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Icon                string  `json:"icon"`
			Ozone               float64 `json:"ozone"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			PrecipType          string  `json:"precipType"`
			Pressure            float64 `json:"pressure"`
			Summary             string  `json:"summary"`
			Temperature         float64 `json:"temperature"`
			Time                int64   `json:"time"`
			UvIndex             int64   `json:"uvIndex"`
			Visibility          float64 `json:"visibility"`
			WindBearing         int64   `json:"windBearing"`
			WindGust            float64 `json:"windGust"`
			WindSpeed           float64 `json:"windSpeed"`
		} `json:"data"`
		Icon    string `json:"icon"`
		Summary string `json:"summary"`
	} `json:"hourly"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Minutely  struct {
		Data []struct {
			PrecipIntensity   int64 `json:"precipIntensity"`
			PrecipProbability int64 `json:"precipProbability"`
			Time              int64 `json:"time"`
		} `json:"data"`
		Icon    string `json:"icon"`
		Summary string `json:"summary"`
	} `json:"minutely"`
	Offset   int64  `json:"offset"`
	Timezone string `json:"timezone"`
}

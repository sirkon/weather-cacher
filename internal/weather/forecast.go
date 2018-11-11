package weather

// Hourly данные по погоде за час
type Hourly struct {
	Time int64 `json:"time,omitempty"` // время в Unix Time для которого сделан прогноз

	Summary     string `json:"summary,omitempty"` // краткое описание
	Temperature struct {
		Measured float64 `json:"measured,omitempty"`  // Температура в °C
		Apparent float64 `json:"apparent,omitempty"`  // Восприятие температуры в °C
		DewPoint float64 `json:"dew_point,omitempty"` // точка росы
	} `json:"temperature,omitempty"`
	Precipitation struct {
		Type        string  `json:"type,omitempty"`        // тип осадков
		Intensity   float64 `json:"intensity,omitempty"`   // количество осадков в мм
		Probability float64 `json:"probability,omitempty"` // вероятность осадков [0,1]
	} `json:"precipitation,omitempty"`
	Humidity float64 `json:"humidity,omitempty"` // влажность (отношение в [0,1])
	Pressure float64 `json:"pressure,omitempty"` // давление в мм ртутного столба
	Wind     struct {
		Speed   float64 `json:"speed,omitempty"`   // скорость ветра в м/с
		Bearing int64   `json:"bearing,omitempty"` // направление ветра
		Gust    float64 `json:"gust,omitempty"`    // максимальные порывы ветра
	}
	CloudCover float64 `json:"cloud_cover,omitempty"` // количество облачности в [0,1]
	UVIndex    uint    `json:"uv_index,omitempty"`    // индекс ультрафиолетового облучения
	Visibility float64 `json:"visibility,omitempty"`  // видимость
}

// Measurement измерение данных
type Measurement struct {
	Value float64 `json:"value,omitempty"` // значение измерения
	Time  int64   `json:"time,omitempty"`  // время измерения
}

// TemperatureMeasurement дневное измерение температуры
type TemperatureMeasurement struct {
	Min  Measurement `json:"min,omitempty"`  // минимальная температура
	Low  Measurement `json:"low,omitempty"`  // "low" температура
	High Measurement `json:"high,omitempty"` // "high" температура
	Max  Measurement `json:"max,omitempty"`  // максимальная температура
}

// Daily данные по погоде за день
type Daily struct {
	Time int64 `json:"time,omitempty"` // время в Unix Time для которого сделан прогноз

	Summary string `json:"summary,omitempty"` // краткое описание
	Sun     struct {
		Rise int64 `json:"rise,omitempty"` // время восхода солнца
		Set  int64 `json:"set,omitempty"`  // время заката солнца
	} `json:"sun,omitempty"`
	MoonPhase   float64 `json:"moon_phase,omitempty"` // фаза луны
	Temperature struct {
		Measured TemperatureMeasurement `json:"measured,omitempty"`  // инструментальное дневная температура
		Apparent TemperatureMeasurement `json:"apparent,omitempty"`  // воспринимаемая дневная температура
		DewPoint float64                `json:"dew_point,omitempty"` // точка росы
	}
	Precipitation struct {
		Type             string  `json:"type,omitempty"`               // тип осадков
		Intensity        float64 `json:"intensity,omitempty"`          // количество осадков в мм
		MaxIntensity     float64 `json:"max_intensity,omitempty"`      // макисмальное количество осадков
		MaxIntensityTime int64   `json:"max_intensity_time,omitempty"` // время максимального кол-ва осадков
		Probability      float64 `json:"probability,omitempty"`        // вероятность осадков [0,1]
	} `json:"precipitation,omitempty"`
	Humidity float64 `json:"humidity,omitempty"` // влажность (отношение в [0,1])
	Pressure float64 `json:"pressure,omitempty"` // давление в мм ртутного столба
	Wind     struct {
		Speed    float64 `json:"speed,omitempty"`     // скорость ветра в м/с
		Bearing  int64   `json:"bearing,omitempty"`   // направление ветра
		Gust     float64 `json:"gust,omitempty"`      // максимальные порывы ветра
		GustTime int64   `json:"gust_time,omitempty"` // время достижения максимальных порывов
	} `json:"wind,omitempty"`
	CloudCover float64 `json:"cloud_cover,omitempty"` // количество облачности в [0,1]
	UV         struct {
		Index uint  `json:"index,omitempty"` // индекс ультрафиолетового облучения
		Time  int64 `json:"time,omitempty"`  // время достижения максимального индекса облучения
	}
	Visibility float64 `json:"visibility,omitempty"` // видимость
}

// Forecast предсказание погоды на определённое число дней + почасовой прогноз погоды на сколько-то часов вперёд
type Forecast struct {
	Current Hourly   `json:"current,omitempty"` // текущая погода
	Hourly  []Hourly `json:"hourly,omitempty"`  // почасовой прогноз погоды
	Daily   []Daily  `json:"daily,omitempty"`   // дневной прогноз погоды
}

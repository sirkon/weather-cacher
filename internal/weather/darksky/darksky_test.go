package darksky

import (
	"bytes"
	"context"
	"io"
	"testing"
)

type closingReader struct {
	*bytes.Buffer
}

func (cr closingReader) Close() error {
	return nil
}

type darkskyRawClient struct{}

func (darkskyRawClient) Get(ctx context.Context, lat, log float64) (io.ReadCloser, error) {
	reader := bytes.NewBufferString(`{"latitude":55.75,"longitude":36.6,"timezone":"Europe/Moscow","currently":{"time":1541838285,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0004,"precipProbability":0.05,"precipType":"snow","temperature":34.7,"apparentTemperature":27.87,"dewPoint":32.15,"humidity":0.9,"pressure":1028.31,"windSpeed":8.47,"windGust":12.48,"windBearing":101,"cloudCover":1,"uvIndex":1,"visibility":8.83,"ozone":286.78},"hourly":{"summary":"Overcast throughout the day.","icon":"cloudy","data":[{"time":1541836800,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0004,"precipProbability":0.04,"precipAccumulation":0,"precipType":"snow","temperature":34.47,"apparentTemperature":27.47,"dewPoint":32.17,"humidity":0.91,"pressure":1028.15,"windSpeed":8.68,"windGust":12.6,"windBearing":101,"cloudCover":1,"uvIndex":0,"visibility":7.58,"ozone":286.78},{"time":1541840400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0005,"precipProbability":0.06,"precipType":"rain","temperature":35.03,"apparentTemperature":28.44,"dewPoint":32.11,"humidity":0.89,"pressure":1028.53,"windSpeed":8.17,"windGust":12.31,"windBearing":102,"cloudCover":1,"uvIndex":1,"visibility":10,"ozone":286.78},{"time":1541844000,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0004,"precipProbability":0.05,"precipType":"rain","temperature":35.45,"apparentTemperature":29,"dewPoint":31.88,"humidity":0.87,"pressure":1028.74,"windSpeed":8.09,"windGust":12.65,"windBearing":102,"cloudCover":1,"uvIndex":1,"visibility":10,"ozone":286.85},{"time":1541847600,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0004,"precipProbability":0.05,"precipType":"rain","temperature":35.67,"apparentTemperature":28.98,"dewPoint":31.67,"humidity":0.85,"pressure":1028.88,"windSpeed":8.6,"windGust":13.29,"windBearing":102,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":286.89},{"time":1541851200,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0005,"precipProbability":0.07,"precipType":"rain","temperature":35.74,"apparentTemperature":28.74,"dewPoint":31.58,"humidity":0.85,"pressure":1029.11,"windSpeed":9.23,"windGust":14.14,"windBearing":101,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":287.24},{"time":1541854800,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0004,"precipProbability":0.07,"precipType":"rain","temperature":35.62,"apparentTemperature":28.57,"dewPoint":31.64,"humidity":0.85,"pressure":1029.54,"windSpeed":9.27,"windGust":15.19,"windBearing":100,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":287.81},{"time":1541858400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0005,"precipProbability":0.08,"precipType":"rain","temperature":35.45,"apparentTemperature":28.47,"dewPoint":31.74,"humidity":0.86,"pressure":1030.06,"windSpeed":9.06,"windGust":16.41,"windBearing":99,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":288.5},{"time":1541862000,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0007,"precipProbability":0.08,"precipType":"rain","temperature":35.36,"apparentTemperature":28.36,"dewPoint":31.82,"humidity":0.87,"pressure":1030.57,"windSpeed":9.04,"windGust":17.57,"windBearing":99,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":289.07},{"time":1541865600,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0008,"precipProbability":0.07,"precipType":"rain","temperature":35.65,"apparentTemperature":28.57,"dewPoint":31.82,"humidity":0.86,"pressure":1031.07,"windSpeed":9.36,"windGust":18.5,"windBearing":100,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":289.07},{"time":1541869200,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0008,"precipProbability":0.07,"precipType":"rain","temperature":35.74,"apparentTemperature":28.45,"dewPoint":31.83,"humidity":0.86,"pressure":1031.57,"windSpeed":9.8,"windGust":19.34,"windBearing":101,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":288.84},{"time":1541872800,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0012,"precipProbability":0.07,"precipType":"rain","temperature":35.68,"apparentTemperature":28.16,"dewPoint":31.7,"humidity":0.85,"pressure":1032.11,"windSpeed":10.27,"windGust":20.21,"windBearing":103,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":288.55},{"time":1541876400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0018,"precipProbability":0.09,"precipAccumulation":0.013,"precipType":"snow","temperature":35.46,"apparentTemperature":27.71,"dewPoint":31.52,"humidity":0.85,"pressure":1032.68,"windSpeed":10.67,"windGust":21.24,"windBearing":104,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":288.39},{"time":1541880000,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0024,"precipProbability":0.11,"precipAccumulation":0.019,"precipType":"snow","temperature":35.08,"apparentTemperature":27.07,"dewPoint":31.3,"humidity":0.86,"pressure":1033.26,"windSpeed":11.03,"windGust":22.34,"windBearing":106,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":288.2},{"time":1541883600,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0026,"precipProbability":0.12,"precipAccumulation":0.02,"precipType":"snow","temperature":34.66,"apparentTemperature":26.34,"dewPoint":30.9,"humidity":0.86,"pressure":1033.91,"windSpeed":11.51,"windGust":23.29,"windBearing":108,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":287.79},{"time":1541887200,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0021,"precipProbability":0.11,"precipAccumulation":0.016,"precipType":"snow","temperature":34.28,"apparentTemperature":25.43,"dewPoint":30.23,"humidity":0.85,"pressure":1034.66,"windSpeed":12.54,"windGust":24.15,"windBearing":110,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":287.22},{"time":1541890800,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0016,"precipProbability":0.1,"precipAccumulation":0.012,"precipType":"snow","temperature":34,"apparentTemperature":24.53,"dewPoint":29.44,"humidity":0.83,"pressure":1035.45,"windSpeed":13.94,"windGust":24.9,"windBearing":113,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":286.59},{"time":1541894400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0013,"precipProbability":0.06,"precipAccumulation":0.01,"precipType":"snow","temperature":33.54,"apparentTemperature":23.53,"dewPoint":28.4,"humidity":0.81,"pressure":1036.28,"windSpeed":15.08,"windGust":25.35,"windBearing":116,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":286.02},{"time":1541898000,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0008,"precipProbability":0.05,"precipAccumulation":0.006,"precipType":"snow","temperature":32.7,"apparentTemperature":22.28,"dewPoint":26.86,"humidity":0.79,"pressure":1037.03,"windSpeed":15.59,"windGust":25.3,"windBearing":117,"cloudCover":0.99,"uvIndex":0,"visibility":10,"ozone":285.65},{"time":1541901600,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0003,"precipProbability":0.04,"precipAccumulation":0,"precipType":"snow","temperature":31.48,"apparentTemperature":20.66,"dewPoint":24.91,"humidity":0.76,"pressure":1037.8,"windSpeed":15.72,"windGust":24.95,"windBearing":119,"cloudCover":0.97,"uvIndex":0,"visibility":10,"ozone":285.42},{"time":1541905200,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":30.18,"apparentTemperature":18.99,"dewPoint":23.11,"humidity":0.75,"pressure":1038.61,"windSpeed":15.74,"windGust":24.63,"windBearing":120,"cloudCover":0.96,"uvIndex":0,"visibility":10,"ozone":285.42},{"time":1541908800,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":28.67,"apparentTemperature":17.04,"dewPoint":21.43,"humidity":0.74,"pressure":1039.5,"windSpeed":15.75,"windGust":24.6,"windBearing":120,"cloudCover":0.95,"uvIndex":0,"visibility":10,"ozone":285.88},{"time":1541912400,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":27.1,"apparentTemperature":15.07,"dewPoint":19.67,"humidity":0.73,"pressure":1040.4,"windSpeed":15.64,"windGust":24.61,"windBearing":120,"cloudCover":0.95,"uvIndex":0,"visibility":10,"ozone":286.59},{"time":1541916000,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":26.13,"apparentTemperature":13.89,"dewPoint":18.35,"humidity":0.72,"pressure":1041.19,"windSpeed":15.45,"windGust":24.55,"windBearing":120,"cloudCover":0.93,"uvIndex":0,"visibility":10,"ozone":287.19},{"time":1541919600,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":26.77,"apparentTemperature":14.82,"dewPoint":17.98,"humidity":0.69,"pressure":1041.81,"windSpeed":15.18,"windGust":24.28,"windBearing":123,"cloudCover":0.88,"uvIndex":0,"visibility":10,"ozone":287.54},{"time":1541923200,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":28.11,"apparentTemperature":16.65,"dewPoint":18.08,"humidity":0.66,"pressure":1042.34,"windSpeed":14.89,"windGust":23.91,"windBearing":126,"cloudCover":0.81,"uvIndex":0,"visibility":10,"ozone":287.65},{"time":1541926800,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":29.84,"apparentTemperature":18.98,"dewPoint":18.12,"humidity":0.61,"pressure":1042.77,"windSpeed":14.6,"windGust":23.53,"windBearing":129,"cloudCover":0.72,"uvIndex":1,"visibility":10,"ozone":287.65},{"time":1541930400,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":30.98,"apparentTemperature":20.51,"dewPoint":17.9,"humidity":0.58,"pressure":1043.05,"windSpeed":14.39,"windGust":23.05,"windBearing":131,"cloudCover":0.61,"uvIndex":1,"visibility":10,"ozone":287.63},{"time":1541934000,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":31.59,"apparentTemperature":21.38,"dewPoint":17.57,"humidity":0.56,"pressure":1043.26,"windSpeed":14.14,"windGust":22.58,"windBearing":132,"cloudCover":0.49,"uvIndex":0,"visibility":10,"ozone":287.45},{"time":1541937600,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":31.07,"apparentTemperature":20.84,"dewPoint":17.2,"humidity":0.56,"pressure":1043.55,"windSpeed":13.84,"windGust":22.35,"windBearing":134,"cloudCover":0.36,"uvIndex":0,"visibility":10,"ozone":287.47},{"time":1541941200,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":29.85,"apparentTemperature":19.47,"dewPoint":16.91,"humidity":0.58,"pressure":1043.91,"windSpeed":13.4,"windGust":22.72,"windBearing":134,"cloudCover":0.23,"uvIndex":0,"visibility":10,"ozone":287.61},{"time":1541944800,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":28.01,"apparentTemperature":17.35,"dewPoint":16.65,"humidity":0.62,"pressure":1044.34,"windSpeed":12.89,"windGust":23.36,"windBearing":134,"cloudCover":0.09,"uvIndex":0,"visibility":10,"ozone":287.74},{"time":1541948400,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":26.49,"apparentTemperature":15.62,"dewPoint":16.07,"humidity":0.64,"pressure":1044.75,"windSpeed":12.49,"windGust":23.82,"windBearing":134,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":288.2},{"time":1541952000,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":25.21,"apparentTemperature":14.06,"dewPoint":15.16,"humidity":0.65,"pressure":1045.07,"windSpeed":12.37,"windGust":23.91,"windBearing":136,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":289.12},{"time":1541955600,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":24.3,"apparentTemperature":12.91,"dewPoint":13.93,"humidity":0.64,"pressure":1045.33,"windSpeed":12.38,"windGust":23.79,"windBearing":138,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":290.3},{"time":1541959200,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":23.56,"apparentTemperature":12.02,"dewPoint":12.79,"humidity":0.63,"pressure":1045.58,"windSpeed":12.31,"windGust":23.62,"windBearing":141,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":291.53},{"time":1541962800,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":22.96,"apparentTemperature":11.39,"dewPoint":11.79,"humidity":0.62,"pressure":1045.85,"windSpeed":12.04,"windGust":23.36,"windBearing":144,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":292.77},{"time":1541966400,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":22.49,"apparentTemperature":10.99,"dewPoint":10.76,"humidity":0.6,"pressure":1046.15,"windSpeed":11.68,"windGust":23.01,"windBearing":146,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":294.05},{"time":1541970000,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":22,"apparentTemperature":10.58,"dewPoint":10,"humidity":0.59,"pressure":1046.37,"windSpeed":11.3,"windGust":22.73,"windBearing":149,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":295.86},{"time":1541973600,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":21.54,"apparentTemperature":10.23,"dewPoint":9.52,"humidity":0.59,"pressure":1046.52,"windSpeed":10.89,"windGust":22.7,"windBearing":152,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":298.58},{"time":1541977200,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":21.06,"apparentTemperature":9.88,"dewPoint":9.29,"humidity":0.6,"pressure":1046.67,"windSpeed":10.46,"windGust":22.75,"windBearing":153,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":301.9},{"time":1541980800,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":20.56,"apparentTemperature":9.47,"dewPoint":9.1,"humidity":0.61,"pressure":1046.75,"windSpeed":10.13,"windGust":22.54,"windBearing":154,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":305.28},{"time":1541984400,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":20,"apparentTemperature":8.95,"dewPoint":8.89,"humidity":0.61,"pressure":1046.7,"windSpeed":9.84,"windGust":21.82,"windBearing":155,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":309.42},{"time":1541988000,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":19.39,"apparentTemperature":8.32,"dewPoint":8.75,"humidity":0.63,"pressure":1046.63,"windSpeed":9.64,"windGust":20.85,"windBearing":155,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":313.79},{"time":1541991600,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":18.88,"apparentTemperature":7.73,"dewPoint":8.66,"humidity":0.64,"pressure":1046.6,"windSpeed":9.59,"windGust":20.21,"windBearing":156,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":316.16},{"time":1541995200,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":18.42,"apparentTemperature":7.03,"dewPoint":8.63,"humidity":0.65,"pressure":1046.82,"windSpeed":9.8,"windGust":20.28,"windBearing":158,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":315.41},{"time":1541998800,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":18.15,"apparentTemperature":6.44,"dewPoint":8.76,"humidity":0.66,"pressure":1047.05,"windSpeed":10.19,"windGust":20.69,"windBearing":159,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":312.63},{"time":1542002400,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":18.65,"apparentTemperature":6.84,"dewPoint":9.1,"humidity":0.66,"pressure":1047.17,"windSpeed":10.54,"windGust":20.89,"windBearing":161,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":309.51},{"time":1542006000,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":20.88,"apparentTemperature":9.45,"dewPoint":9.85,"humidity":0.62,"pressure":1047.12,"windSpeed":10.82,"windGust":20.63,"windBearing":165,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":306.21},{"time":1542009600,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":23.99,"apparentTemperature":13.2,"dewPoint":10.79,"humidity":0.56,"pressure":1046.97,"windSpeed":11.07,"windGust":20.19,"windBearing":168,"cloudCover":0,"uvIndex":0,"visibility":10,"ozone":302.36}]},"daily":{"summary":"Light snow (< 1 in.) next Saturday, with high temperatures falling to 30°F on Tuesday.","icon":"snow","data":[{"time":1541797200,"summary":"Overcast throughout the day.","icon":"cloudy","sunriseTime":1541825993,"sunsetTime":1541857069,"moonPhase":0.09,"precipIntensity":0.0008,"precipIntensityMax":0.0024,"precipIntensityMaxTime":1541880000,"precipProbability":0.22,"precipType":"rain","temperatureHigh":35.74,"temperatureHighTime":1541851200,"temperatureLow":27.1,"temperatureLowTime":1541912400,"apparentTemperatureHigh":29,"apparentTemperatureHighTime":1541844000,"apparentTemperatureLow":15.07,"apparentTemperatureLowTime":1541912400,"dewPoint":31.26,"humidity":0.87,"pressure":1028.78,"windSpeed":7.96,"windGust":22.34,"windGustTime":1541880000,"windBearing":104,"cloudCover":0.97,"uvIndex":1,"uvIndexTime":1541840400,"visibility":10,"ozone":287.96,"temperatureMin":33.76,"temperatureMinTime":1541804400,"temperatureMax":35.74,"temperatureMaxTime":1541851200,"apparentTemperatureMin":26.44,"apparentTemperatureMinTime":1541818800,"apparentTemperatureMax":35.57,"apparentTemperatureMaxTime":1541800800},{"time":1541883600,"summary":"Mostly cloudy until afternoon.","icon":"partly-cloudy-day","sunriseTime":1541912518,"sunsetTime":1541943358,"moonPhase":0.12,"precipIntensity":0.0004,"precipIntensityMax":0.0026,"precipIntensityMaxTime":1541883600,"precipProbability":0.2,"precipAccumulation":0.07,"precipType":"snow","temperatureHigh":31.59,"temperatureHighTime":1541934000,"temperatureLow":18.15,"temperatureLowTime":1541998800,"apparentTemperatureHigh":21.38,"apparentTemperatureHighTime":1541934000,"apparentTemperatureLow":6.44,"apparentTemperatureLowTime":1541998800,"dewPoint":19.76,"humidity":0.69,"pressure":1041.36,"windSpeed":13.68,"windGust":25.35,"windGustTime":1541894400,"windBearing":126,"cloudCover":0.58,"uvIndex":1,"uvIndexTime":1541926800,"visibility":10,"ozone":287.94,"temperatureMin":22.49,"temperatureMinTime":1541966400,"temperatureMax":34.66,"temperatureMaxTime":1541883600,"apparentTemperatureMin":10.99,"apparentTemperatureMinTime":1541966400,"apparentTemperatureMax":26.34,"apparentTemperatureMaxTime":1541883600},{"time":1541970000,"summary":"Clear throughout the day.","icon":"clear-day","sunriseTime":1541999043,"sunsetTime":1542029649,"moonPhase":0.15,"precipIntensity":0,"precipIntensityMax":0,"precipProbability":0,"temperatureHigh":30.58,"temperatureHighTime":1542020400,"temperatureLow":17.39,"temperatureLowTime":1542085200,"apparentTemperatureHigh":21.39,"apparentTemperatureHighTime":1542020400,"apparentTemperatureLow":4.79,"apparentTemperatureLowTime":1542085200,"dewPoint":9.61,"humidity":0.57,"pressure":1046.26,"windSpeed":10.32,"windGust":26.12,"windGustTime":1542052800,"windBearing":166,"cloudCover":0,"uvIndex":1,"uvIndexTime":1542013200,"visibility":10,"ozone":299.57,"temperatureMin":18.15,"temperatureMinTime":1541998800,"temperatureMax":30.58,"temperatureMaxTime":1542020400,"apparentTemperatureMin":6.44,"apparentTemperatureMinTime":1541998800,"apparentTemperatureMax":21.39,"apparentTemperatureMaxTime":1542020400},{"time":1542056400,"summary":"Mostly cloudy throughout the day.","icon":"partly-cloudy-day","sunriseTime":1542085568,"sunsetTime":1542115942,"moonPhase":0.18,"precipIntensity":0,"precipIntensityMax":0.0002,"precipIntensityMaxTime":1542056400,"precipProbability":0,"temperatureHigh":30.04,"temperatureHighTime":1542110400,"temperatureLow":24.15,"temperatureLowTime":1542132000,"apparentTemperatureHigh":19.61,"apparentTemperatureHighTime":1542110400,"apparentTemperatureLow":12.08,"apparentTemperatureLowTime":1542132000,"dewPoint":3.9,"humidity":0.45,"pressure":1041.1,"windSpeed":12.33,"windGust":32.89,"windGustTime":1542139200,"windBearing":198,"cloudCover":0.53,"uvIndex":1,"uvIndexTime":1542099600,"visibility":10,"ozone":275.75,"temperatureMin":17.39,"temperatureMinTime":1542085200,"temperatureMax":30.04,"temperatureMaxTime":1542110400,"apparentTemperatureMin":4.79,"apparentTemperatureMinTime":1542085200,"apparentTemperatureMax":19.61,"apparentTemperatureMaxTime":1542110400},{"time":1542142800,"summary":"Foggy in the evening.","icon":"fog","sunriseTime":1542172092,"sunsetTime":1542202237,"moonPhase":0.21,"precipIntensity":0.0015,"precipIntensityMax":0.0065,"precipIntensityMaxTime":1542211200,"precipProbability":0.26,"precipAccumulation":0.332,"precipType":"snow","temperatureHigh":33.72,"temperatureHighTime":1542196800,"temperatureLow":29.09,"temperatureLowTime":1542214800,"apparentTemperatureHigh":23.81,"apparentTemperatureHighTime":1542196800,"apparentTemperatureLow":18.49,"apparentTemperatureLowTime":1542214800,"dewPoint":10.47,"humidity":0.5,"pressure":1030.5,"windSpeed":14.45,"windGust":32.99,"windGustTime":1542146400,"windBearing":206,"cloudCover":1,"uvIndex":1,"uvIndexTime":1542186000,"visibility":10,"ozone":277.55,"temperatureMin":24.69,"temperatureMinTime":1542142800,"temperatureMax":33.72,"temperatureMaxTime":1542196800,"apparentTemperatureMin":12.34,"apparentTemperatureMinTime":1542175200,"apparentTemperatureMax":23.81,"apparentTemperatureMaxTime":1542196800},{"time":1542229200,"summary":"Overcast throughout the day.","icon":"cloudy","sunriseTime":1542258616,"sunsetTime":1542288535,"moonPhase":0.24,"precipIntensity":0.0016,"precipIntensityMax":0.004,"precipIntensityMaxTime":1542236400,"precipProbability":0.3,"precipAccumulation":0.311,"precipType":"snow","temperatureHigh":35.35,"temperatureHighTime":1542279600,"temperatureLow":30.54,"temperatureLowTime":1542319200,"apparentTemperatureHigh":26.98,"apparentTemperatureHighTime":1542279600,"apparentTemperatureLow":21.36,"apparentTemperatureLowTime":1542319200,"dewPoint":26.65,"humidity":0.8,"pressure":1029.51,"windSpeed":11.48,"windGust":19.54,"windGustTime":1542312000,"windBearing":198,"cloudCover":0.98,"uvIndex":0,"uvIndexTime":1542229200,"visibility":10,"ozone":297.97,"temperatureMin":30.37,"temperatureMinTime":1542229200,"temperatureMax":35.35,"temperatureMaxTime":1542279600,"apparentTemperatureMin":20.37,"apparentTemperatureMinTime":1542229200,"apparentTemperatureMax":26.98,"apparentTemperatureMaxTime":1542279600},{"time":1542315600,"summary":"Overcast throughout the day.","icon":"cloudy","sunriseTime":1542345138,"sunsetTime":1542374835,"moonPhase":0.27,"precipIntensity":0.0027,"precipIntensityMax":0.0077,"precipIntensityMaxTime":1542344400,"precipProbability":0.5,"precipAccumulation":0.294,"precipType":"snow","temperatureHigh":35.14,"temperatureHighTime":1542366000,"temperatureLow":20.8,"temperatureLowTime":1542405600,"apparentTemperatureHigh":27.97,"apparentTemperatureHighTime":1542366000,"apparentTemperatureLow":13.96,"apparentTemperatureLowTime":1542412800,"dewPoint":25.2,"humidity":0.8,"pressure":1027.94,"windSpeed":9.42,"windGust":29.39,"windGustTime":1542344400,"windBearing":194,"cloudCover":1,"uvIndex":0,"uvIndexTime":1542315600,"visibility":10,"ozone":301.04,"temperatureMin":23.01,"temperatureMinTime":1542398400,"temperatureMax":35.14,"temperatureMaxTime":1542366000,"apparentTemperatureMin":16.4,"apparentTemperatureMinTime":1542398400,"apparentTemperatureMax":27.97,"apparentTemperatureMaxTime":1542366000},{"time":1542402000,"summary":"Foggy in the evening.","icon":"fog","sunriseTime":1542431660,"sunsetTime":1542461138,"moonPhase":0.3,"precipIntensity":0.0028,"precipIntensityMax":0.0133,"precipIntensityMaxTime":1542466800,"precipProbability":0.39,"precipAccumulation":0.469,"precipType":"snow","temperatureHigh":35.6,"temperatureHighTime":1542452400,"temperatureLow":23.19,"temperatureLowTime":1542517200,"apparentTemperatureHigh":28.89,"apparentTemperatureHighTime":1542452400,"apparentTemperatureLow":13.86,"apparentTemperatureLowTime":1542517200,"dewPoint":22.54,"humidity":0.82,"pressure":1022.55,"windSpeed":6.26,"windGust":26.62,"windGustTime":1542434400,"windBearing":156,"cloudCover":0.99,"uvIndex":0,"uvIndexTime":1542402000,"visibility":10,"ozone":310.21,"temperatureMin":20.8,"temperatureMinTime":1542405600,"temperatureMax":35.6,"temperatureMaxTime":1542452400,"apparentTemperatureMin":13.96,"apparentTemperatureMinTime":1542412800,"apparentTemperatureMax":28.89,"apparentTemperatureMaxTime":1542452400}]},"flags":{"sources":["cmc","gfs","icon","isd","madis"],"nearest-station":26.986,"units":"us"},"offset":3}`)
	return closingReader{
		Buffer: reader,
	}, nil
}

func TestClient(t *testing.T) {
	c := Client(darkskyRawClient{})
	_, err := c.WeatherFor(context.Background(), 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	// здесь пока нет настоящих тестов
}

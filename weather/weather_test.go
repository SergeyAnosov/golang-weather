package weather_test

import (
	"golang/weather/geo"
	"golang/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	//arrange
	format := 3
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}

	//act
	result, err := weather.GetWeather(geoData, format)

	//assert
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	bool := strings.Contains(result, expected)
	if !bool {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "big_format", format: 147},
	{name: "zero_format", format: 0},
	{name: "minus_format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrorWrongFormat {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrorWrongFormat, err)
			}
		})
	}

}

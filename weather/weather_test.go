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

func TestGetWeatherWrongFormat(t *testing.T) {
	format := 125
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}

	_, err := weather.GetWeather(geoData, format)
	if err != weather.ErrorWrongFormat {
		t.Errorf("Ожидалось %v, получено %v", weather.ErrorWrongFormat, err)
	}
}
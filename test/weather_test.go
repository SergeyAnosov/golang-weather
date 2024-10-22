package test

import (
	"golang/weather/geo"
	"golang/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	//arrange
	var format = 1
	geo := geo.GeoData{
		City: "London",
	}

	//act
	result := weather.GetWeather(geo, format)

	//assert
	bool := strings.Contains(result, "London")
	if !bool {
		t.Errorf("Ожидалось %v, получено %v", geo.City, result)
	}
}

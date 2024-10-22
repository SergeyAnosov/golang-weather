package geo_test

import (
	"golang/weather/geo"
	"testing"
)

func TestGetMyLocation_(t *testing.T) {
	// Arange - подготовка, expected result, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - вызываем функцию
	got, err := geo.GetMyLocation(city)
	
	// Assert - проверка результата с ожиданием
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}	
}
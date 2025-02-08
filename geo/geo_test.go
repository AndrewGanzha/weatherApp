package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arange - подготовка, expected результат, данные для функции
	city := "London"

	expected := geo.GeoData{City: "London"}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != nil {
		t.Error("Ошибка получения города")
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Beleberda"

	_, err := geo.GetMyLocation(city)

	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}

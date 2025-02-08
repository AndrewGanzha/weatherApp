package weather_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

func TestGetWeather(t *testing.T) {
	city := "London"
	format := 3
	geodata := geo.GeoData{
		City: city,
	}

	result, err := weather.GetWeather(geodata, format)

	if err != nil {
		t.Errorf("Пришла ошибка" + err.Error())
	}

	if !strings.Contains(result, city) {
		t.Errorf("city does not contains %s", city)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -28},
}

func TestGetWeatherWithInvalidFormat(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			city := "London"
			geodata := geo.GeoData{
				City: city,
			}

			_, err := weather.GetWeather(geodata, testCase.format)

			fmt.Println(err)

			if !errors.Is(err, weather.ErrWrongFormat) {
				t.Errorf("Ожидалось получение другого формата")
			}
		})
	}
}

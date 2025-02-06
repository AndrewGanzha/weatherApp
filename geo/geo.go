package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := CheckCity(city)

		if !isCity {
			return nil, errors.New(fmt.Sprintf("city %s is not exist", city))
		}

		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://ipapi.co/json")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var geo GeoData

	defer resp.Body.Close()

	json.Unmarshal(body, &geo)

	return &geo, nil
}

func CheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{"city": city})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var population CityPopulationResponse

	json.Unmarshal(body, &population)

	return !population.error
}

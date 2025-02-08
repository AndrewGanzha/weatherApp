package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

var ErrWrongFormat = errors.New("ERROR_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())

	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return "", errors.New("ERROR_BODY")
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(body), nil
}

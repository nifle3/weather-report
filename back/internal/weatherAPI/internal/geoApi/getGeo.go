package geoApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseBody struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func (g GeoApi) GetLatLon(country, state, city string) (float64, float64, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, g.apiKey)
	respBody, err := http.Get(url)

	if err != nil {
		return 0, 0, err
	}

	var respResult []responseBody
	if err = json.NewDecoder(respBody.Body).Decode(&respResult); err != nil {
		return 0, 0, err
	}

	var lat, lon float64
	for _, value := range respResult {
		if value.Country == country && value.State == state {
			lat, lon = value.Lat, value.Lon
		}
	}

	return lat, lon, nil
}

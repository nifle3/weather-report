package weatherAPI

import (
	"back/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type currentJson struct {
	Current domain.Weather `json:"main"`
}

func (w *WeatherAPI) GetWeatherNow(city, state, country string) (domain.Weather, error) {
	lat, lon, err := w.geoApi.GetLatLon(city, state, country)
	if err != nil {
		return domain.Weather{}, err
	}

	apiURI := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=&appid=%s",
		lat, lon, w.apiAccessCode)

	body, err := http.Get(apiURI)
	if err != nil {
		return domain.Weather{}, err
	}

	var current currentJson

	err = json.NewDecoder(body.Body).Decode(&current)
	if err != nil {
		return domain.Weather{}, err
	}

	return current.Current, nil
}

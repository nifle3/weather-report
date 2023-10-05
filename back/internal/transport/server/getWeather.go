package server

import (
	"encoding/json"
	"net/http"
)

func (c *customMux) getWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type cityAndCountry struct {
		City    string `json:"City"`
		State   string `json:"State"`
		Country string `json:"Country"`
	}

	var result cityAndCountry
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	weather, err := c.api.GetWeatherNow(result.City, result.State, result.Country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonWeather, err := json.Marshal(weather)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonWeather)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

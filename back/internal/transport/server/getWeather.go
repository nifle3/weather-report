package server

import (
	"back/internal/domain"
	"encoding/json"
	"net/http"
)

func (c *customMux) getWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var result domain.CityAndCountry
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	return
}

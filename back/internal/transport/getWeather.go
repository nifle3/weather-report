package transport

import (
	"encoding/json"
	"net/http"
)

func (h *HttpServer) getWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type cityAndCountry struct {
		City    string `json:"City"`
		Country string `json:"Country"`
		Time    string `json:"Time"`
	}

	var result cityAndCountry
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	weather := h.api.GetWeather(result.City, result.Country)

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

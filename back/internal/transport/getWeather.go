package transport

import (
	"encoding/json"
	"net/http"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
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

}

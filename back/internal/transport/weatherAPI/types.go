package weatherAPI

import (
	"back/internal/transport/weatherAPI/internal/geoAPI"
)

type geoApi interface {
	GetLatLon(country, state, city string) (float64, float64, error)
}

type WeatherAPI struct {
	apiAccessCode string
	geoApi        geoApi
}

func New(code string) *WeatherAPI {
	geoApi := geoAPI.New(code)

	return &WeatherAPI{
		apiAccessCode: code,
		geoApi:        geoApi,
	}
}

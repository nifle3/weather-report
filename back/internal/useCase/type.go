package useCase

import (
	"back/internal/domain"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=WeatherApi
type weatherApi interface {
	GetWeatherNow(city, state, country string) (domain.Weather, error)
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=Redis
type storage interface {
	GetWeather(key string) (domain.Weather, error)
	AddWeather(key string, weather domain.Weather) error
}

type Weather struct {
	storage *storage
	api     *weatherApi
}

func New(str storage, api weatherApi) Weather {
	return Weather{
		storage: &str,
		api:     &api,
	}
}

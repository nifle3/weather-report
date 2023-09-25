package weatherAPI

import "back/internal/domain"

type currentJson struct {
	Current currentInnerJson `json:"main"`
}

type currentInnerJson struct {
	Temp     float64 `json:"temp"`
	TempLike float64 `json:"feels_like"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

func (w *WeatherAPI) GetWeatherNow(city, state, country string) (domain.Weather, error) {
	lat, lon, err := w.geoApi.GetLatLon(city, state, country)
	if err != nil {
		return domain.Weather{}, err
	}

	return domain.Weather{}, nil
}

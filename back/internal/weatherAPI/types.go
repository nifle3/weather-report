package weatherAPI

type WeatherAPI struct {
}

func New() *WeatherAPI {
	return &WeatherAPI{}
}

func (w *WeatherAPI) GetWeather(city, country string) {

}

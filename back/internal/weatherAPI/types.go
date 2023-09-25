package weatherAPI

type geoApi interface {
	GetLatLon(country, state, city string) (float64, float64, error)
}

type WeatherAPI struct {
	apiAccessCode string
	geoApi        geoApi
}

func New(code string, geoApi geoApi) *WeatherAPI {
	return &WeatherAPI{
		apiAccessCode: code,
		geoApi:        geoApi,
	}
}

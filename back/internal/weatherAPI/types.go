package weatherAPI

type WeatherAPI struct {
	apiAccessCode string
}

func New(code string) *WeatherAPI {
	return &WeatherAPI{
		apiAccessCode: code,
	}
}

package geoAPI

type GeoApi struct {
	apiKey string
}

func New(apiKey string) GeoApi {
	return GeoApi{
		apiKey: apiKey,
	}
}

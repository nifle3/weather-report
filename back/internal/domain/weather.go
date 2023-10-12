package domain

type Weather struct {
	Temp     float64 `json:"temp" `
	TempLike float64 `json:"feels_like"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

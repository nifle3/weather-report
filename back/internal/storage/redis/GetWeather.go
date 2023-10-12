package redis

import (
	"back/internal/domain"
	"context"
	"encoding/json"
)

func (s *Storage) GetWeather(key string) (domain.Weather, error) {
	data, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		return domain.Weather{}, err
	}

	var result domain.Weather
	err = json.Unmarshal([]byte(data), &result)

	return result, err
}

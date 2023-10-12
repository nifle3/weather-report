package redis

import (
	"back/internal/domain"
	"context"
	"encoding/json"
	"time"
)

func (s *Storage) AddWeather(key string, weather domain.Weather) error {
	ctx := context.Background()

	jsonData, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	_, err = s.client.Set(ctx, key, jsonData, time.Second*30).Result()

	return err
}

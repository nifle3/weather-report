package app

import (
	"back/internal/app/cfg"
	"back/internal/transport"
	"back/internal/weatherAPI"
)

func Start() {
	// log := logger.New()
	config := cfg.New()
	API := weatherAPI.New(config.Key)
	server := transport.New(API)
	server.StartListen()
}

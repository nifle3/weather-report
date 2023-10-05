package app

import (
	"back/internal/app/cfg"
	"back/internal/transport/server"
	"back/internal/transport/weatherAPI"
)

func Start() {
	config := cfg.New()
	API := weatherAPI.New(config.Key)
	listener := server.New(API)
	listener.StartListen()
}

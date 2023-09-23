package app

import (
	"back/internal/transport"
	"back/internal/weatherAPI"
)

func Start() {
	// log := logger.New()
	API := weatherAPI.New()
	server := transport.New(API)
	server.StartListen()
}

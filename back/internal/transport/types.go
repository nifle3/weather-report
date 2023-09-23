package transport

import (
	"log"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=weatherApi
type weatherApi interface {
	GetWeather(city, country string)
}

type HttpServer struct {
	server *http.Server
	api    weatherApi
}

func New(api weatherApi) *HttpServer {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    "8080",
		Handler: mux,
	}

	return &HttpServer{
		server: server,
		api:    api,
	}
}

func (h *HttpServer) StartListen() {
	log.Fatal(h.server.ListenAndServe())
}

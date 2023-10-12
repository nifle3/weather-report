package server

import (
	"log"
	"net/http"
)

type useCase interface {
}

type customMux struct {
	api useCase
}

type HttpServer struct {
	server *http.Server
}

func New(api useCase) *HttpServer {
	mux := http.NewServeMux()

	controllers := customMux{api: api}

	mux.HandleFunc("/get/weather", controllers.getWeather)

	server := &http.Server{
		Addr:    "8080",
		Handler: mux,
	}

	return &HttpServer{
		server: server,
	}
}

func (h *HttpServer) StartListen() {
	log.Fatal(h.server.ListenAndServe())
}

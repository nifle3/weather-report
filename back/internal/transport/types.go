package transport

import (
	"back/internal/domain"
	"log"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=weatherApi
type weatherApi interface {
	GetWeather(city, state, country string) (domain.Weather, error)
}

type customMux struct {
	api weatherApi
}

type HttpServer struct {
	server *http.Server
}

func New(api weatherApi) *HttpServer {
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

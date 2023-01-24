package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIConfig struct {
	Host string
	Port int
}

type APIServer struct {
	config APIConfig
	router *mux.Router
}

func New(c *APIConfig) *APIServer {
	s := &APIServer{
		config: APIConfig{c.Host, c.Port},
		router: mux.NewRouter(),
	}

	s.configurateRouter()

	return s
}

func (a *APIServer) configurateRouter() {
	a.router.HandleFunc("/health", a.handleCheckHealth())
}

func (a *APIServer) handleCheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"status\": \"OK\"}"))
	}
}

func (a *APIServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", a.config.Host, a.config.Port), a.router)
}

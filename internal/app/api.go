package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIConfig struct {
	Host     string
	Port     int
	LogLevel string
}

type APIServer struct {
	config APIConfig
	router *mux.Router
	logger *logrus.Logger
}

func New(c APIConfig) *APIServer {
	return &APIServer{
		config: APIConfig{c.Host, c.Port, c.LogLevel},
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (a *APIServer) configurateLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}

func (a *APIServer) configurateRouter() {
	a.router.Use(a.logRequest)
	a.router.HandleFunc("/health", a.handleCheckHealth())
}

func (a *APIServer) handleCheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\": \"OK\"}"))
	}
}

func (a *APIServer) logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := a.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)
		code := http.StatusOK
		start := time.Now()
		w.WriteHeader(code)

		h.ServeHTTP(w, r)
		logger.Logf(
			logrus.InfoLevel,
			"completed with %d %s in %v",
			http.StatusOK,
			http.StatusText(code),
			time.Since(start),
		)
	})
}

func (a *APIServer) Start() error {
	if err := a.configurateLogger(); err != nil {
		return err
	}

	a.logger.Info("Starting logger ...")

	a.configurateRouter()

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", a.config.Host, a.config.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           a.router,
	}

	return server.ListenAndServe()
}

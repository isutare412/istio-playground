package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/isutare412/istio-playground/api-server/pkg/config"
	"github.com/isutare412/istio-playground/api-server/pkg/core/health"
	"github.com/isutare412/istio-playground/api-server/pkg/core/user"
	log "github.com/sirupsen/logrus"
)

type server struct {
	server *http.Server
	done   chan struct{}
}

func (s *server) Start(ctx context.Context) <-chan error {
	errNotify := make(chan error)
	go func() {
		defer close(errNotify)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errNotify <- err
		}
	}()
	log.Infof("HTTP server started on %s", s.server.Addr)

	go func() {
		<-ctx.Done()
		s.shutdown()
	}()

	return errNotify
}

func (s *server) Done() <-chan struct{} {
	return s.done
}

func (s *server) shutdown() {
	log.Info("HTTP server shutdown start")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	defer close(s.done)

	if err := s.server.Shutdown(ctx); err != nil {
		log.Errorf("failed to shutdown HTTP server: %v", err)
	}

	log.Info("HTTP server shutdown finished successfully")
}

func NewServer(cfg *config.HttpConfig, hSvc health.Service, uSvc user.Service) *server {
	accessLog := structAccessLog
	if config.IsDevelopmentMode() {
		accessLog = plainAccessLog
	}

	r := mux.NewRouter()
	r.Use(accessLog, tracing)

	r.HandleFunc("/liveness", liveness(hSvc)).Methods("GET")
	r.HandleFunc("/readiness", readiness(hSvc)).Methods("GET")

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/hello/{name}", sayHello(uSvc))

	return &server{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler: r,
		},
		done: make(chan struct{}),
	}
}

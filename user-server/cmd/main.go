package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/isutare412/istio-playground/user-server/pkg/adapter/http"
	"github.com/isutare412/istio-playground/user-server/pkg/adapter/tracer"
	"github.com/isutare412/istio-playground/user-server/pkg/config"
	"github.com/isutare412/istio-playground/user-server/pkg/core/health"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const cfgEnvStr = "USER_SERVER_CONFIG"

func main() {
	cfgPath := os.Getenv(cfgEnvStr)
	if cfgPath == "" {
		log.Fatalf("need environment variable: %s", cfgEnvStr)
	}
	cfg, err := readConfig(cfgPath)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	config.SetMode(config.Mode(cfg.Mode))
	if config.IsDevelopmentMode() {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}

	rootCtx, cancel := context.WithCancel(context.Background())

	tracer, tracerCloser, err := tracer.NewTracer(&cfg.Tracer)
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}
	defer tracerCloser.Close()
	opentracing.SetGlobalTracer(tracer)
	log.Info("registered global tracer")

	hSvc := health.NewService()
	log.Info("created health service")

	server := http.NewServer(&cfg.Http, hSvc)
	log.Info("created http server")

	srvErrors := server.Start(rootCtx)

	sig := make(chan os.Signal, 3)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case e := <-srvErrors:
		log.Errorf("got error from http server: %v", e)
	case s := <-sig:
		log.Infof("caught signal[%s]", s.String())
	}

	cancel()
	<-server.Done()
}

func readConfig(path string) (*config.Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

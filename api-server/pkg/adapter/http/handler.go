package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isutare412/istio-playground/api-server/pkg/core/health"
	"github.com/isutare412/istio-playground/api-server/pkg/core/user"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
)

func liveness(hSvc health.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hSvc.Liveness(); err != nil {
			log.Errorf("liveness failed: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func readiness(hSvc health.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hSvc.Readiness(); err != nil {
			log.Errorf("readiness failed: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func sayHello(uSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, ctx := opentracing.StartSpanFromContext(r.Context(), "http.handler.sayHello")
		defer span.Finish()

		name, ok := mux.Vars(r)["name"]
		if !ok {
			log.Warnf("name field not given")
			responseError(w, http.StatusBadRequest, "'name' is mandatory field")
			return
		}

		usr, err := uSvc.GetUser(ctx, name)
		if err != nil {
			log.Errorf("failed to get user: %v", err)
			responseError(w, http.StatusInternalServerError, "failed to get user")
			return
		}

		responseJson(w, &sayHelloResp{
			Name:     usr.Name,
			Age:      usr.Age,
			Sentence: usr.Sentence,
		})
	}
}

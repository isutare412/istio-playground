package http

import (
	"math/rand"
	"net/http"

	lorem "github.com/drhodes/golorem"
	"github.com/gorilla/mux"
	"github.com/isutare412/istio-playground/user-server/pkg/core/health"
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

func getUser(w http.ResponseWriter, r *http.Request) {
	span, _ := opentracing.StartSpanFromContext(r.Context(), "http.handler.getUser")
	defer span.Finish()

	name, ok := mux.Vars(r)["name"]
	if !ok {
		log.Warnf("name field not given")
		responseError(w, http.StatusBadRequest, "'name' is mandatory field")
		return
	}

	sentence := lorem.Sentence(10, 20)
	age := int32(rand.Intn(20) + 20)

	responseJson(w, &getUserResp{
		Name:     name,
		Age:      age,
		Sentence: sentence,
	})
}

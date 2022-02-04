package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isutare412/istio-playground/api-server/pkg/core/user"
	log "github.com/sirupsen/logrus"
)

func sayHello(uSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

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

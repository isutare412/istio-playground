package http

import (
	"math/rand"
	"net/http"

	lorem "github.com/drhodes/golorem"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getUser(w http.ResponseWriter, r *http.Request) {
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

package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const OCM_OPERATION_ID_HEADER = "X-Operation-Id"

func errorMessageResponse(err error, w http.ResponseWriter) {
	log.Error(err)
	http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
}

func invalidRequestVerbResponse(method string, w http.ResponseWriter) {
	log.Errorf("Invalid request verb: %s", method)
	http.Error(w, "Bad request body", http.StatusBadRequest)
}

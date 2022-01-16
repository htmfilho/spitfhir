package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "FHIR!\n")
	}
}

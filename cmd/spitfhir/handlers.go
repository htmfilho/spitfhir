package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "SpitFHIR!\n")
}

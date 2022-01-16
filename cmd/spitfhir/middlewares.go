package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) spit(h httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Fprint(res, "Spit ")
		h(res, req, params)
	}
}

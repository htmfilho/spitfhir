package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type server struct {
	name   string
	router *httprouter.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	srv := &server{name: "SpitFHIR"}
	srv.router = httprouter.New()
	srv.routes()
	return srv
}

func run() error {
	server := newServer()
	fmt.Printf("%s Serving at http://localhost:8000\n", server.name)
	log.Fatal(http.ListenAndServe(":8000", server.router))
	return nil
}

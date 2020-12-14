package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server is struct for HTTP server
type Server struct {
	router *mux.Router
	db     int
}

func main() {
	run()
}

func run() {
	srvr := &Server{
		router: mux.NewRouter().StrictSlash(true),
	}
	srvr.routes()
	log.Fatal(http.ListenAndServe(":8080", srvr.router))
}

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
	httpSrvr := &Server{
		router: mux.NewRouter().StrictSlash(true),
	}
	httpSrvr.routes()
	log.Printf("HTTP server started in port 8080")
	log.Fatal(http.ListenAndServe(":8080", httpSrvr.router))
}

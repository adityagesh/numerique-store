package main

import (
	"net/http"

	"github.com/adityagesh/numerique-store/bff/customer"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// ServerContext is for initialising global server resources
type serverContext struct {
	router *mux.Router
	db     int
}

func main() {
	run()
}

func run() {
	customerContext := newCustomerContext()
	httpSrvr := &serverContext{
		router: mux.NewRouter().StrictSlash(true),
	}
	httpSrvr.initRoutes(customerContext)
	log.Printf("HTTP server started in port 8080")
	log.Fatal(http.ListenAndServe(":8080", httpSrvr.router))
}

func newCustomerContext() *customer.Context {
	customerContext := &customer.Context{
		GRPCConnection: customer.NewGRPCClientConnection(),
	}
	return customerContext
}

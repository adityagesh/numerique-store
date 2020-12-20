package main

import (
	"github.com/adityagesh/numerique-store/bff/customer"
)

func (s *serverContext) initRoutes(ctx *customer.Context) {
	s.router.HandleFunc("/customer", customer.Handler)
	s.router.HandleFunc("/customer/register", ctx.Register).Methods("POST")
}

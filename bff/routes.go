package main

import (
	"github.com/adityagesh/numerique-store/bff/customer"
)

func (s *Server) routes() {
	s.router.HandleFunc("/customer", customer.Handler)
	s.router.HandleFunc("/customer/register", customer.Register)
}

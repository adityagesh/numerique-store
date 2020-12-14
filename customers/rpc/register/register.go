package register

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Server is struct for customer registration rpc
type Server struct {
	UnimplementedRegisterServiceServer
}

// RegisterCustomer is used to Register a customer
func (s *Server) RegisterCustomer(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	log.Printf("Recieved customer registration %v", request.UserName)
	return &RegisterResponse{Message: "User created successfully", Uid: "aZWE1", Error: nil}, nil
}

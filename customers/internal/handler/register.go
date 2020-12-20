package handler

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adityagesh/numerique-store/customers/internal/rpc/register"
	log "github.com/sirupsen/logrus"
)

// RegisterServer is struct for customer registration rpc
type RegisterServer struct {
	register.UnimplementedRegisterServiceServer
	Log *log.Logger
	DB  *sql.DB
}

// RegisterCustomer is handler that implementes gRPC method to Register Customer
func (s *RegisterServer) RegisterCustomer(ctx context.Context, request *register.RegisterRequest) (*register.RegisterResponse, error) {
	sql := "INSERT INTO customer(`username`, `first_name`, `last_name`, `email`, `phone`) VALUES (?, ?, ?, ?, ?)"
	res, err := s.DB.Exec(sql, request.UserName, request.FirstName, request.LastName, request.Email, request.Phone)
	if err != nil {
		log.Printf("Error registering customer %v", err)
		return &register.RegisterResponse{Message: "Error creating user",
			Error: &register.Error{Message: fmt.Sprint(err), Code: "500"}}, nil
	}
	id, _ := res.LastInsertId()
	log.Printf("Customer created successfully: %v", id)
	return &register.RegisterResponse{Message: "Customer created successfully", Id: id}, nil
}

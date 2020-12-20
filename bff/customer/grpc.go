package customer

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// var grpcConnection *grpc.ClientConn = newGRPCClientConnection()

// NewGRPCClientConnection is used to create GRPC Client connection for customer backend
func NewGRPCClientConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Errorf("Could not connect to GRPC server at port 9000")
	}

	return conn
}

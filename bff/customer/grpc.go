package customer

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func grpcClientConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Errorf("Could not connect to GRPC server at port 9000")
	}

	return conn, err
}

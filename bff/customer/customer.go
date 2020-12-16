package customer

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adityagesh/numerique-store/bff/customer/register"
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

// Handler is HTTP handler for customer
func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, req.URL.Path)

}

// Register is HTTP handler for customer register
func Register(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error reading body for request %v", r)
	}
	d := string(bytes.TrimSpace(b))
	fmt.Fprintf(w, d)
	conn, err := grpcClientConnection()
	defer conn.Close()
	registerServiceClient := register.NewRegisterServiceClient(conn)
	customerDetails := register.RegisterRequest{
		UserName:  "aditya",
		FirstName: "aditya",
		LastName:  "nagesh",
	}
	resp, err := registerServiceClient.RegisterCustomer(context.Background(), &customerDetails)
	if err != nil {
		log.Printf("Error creating user %v", err)
	}
	log.Printf("Response from Server %v", resp)

}

package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adityagesh/numerique-store/bff/customer/register"
	log "github.com/sirupsen/logrus"
)

// Handler is HTTP handler for customer
func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, req.URL.Path)

}

// Register is HTTP handler to register customer
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var customerDetails register.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&customerDetails)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		log.Printf("Registering customer: %v ", customerDetails.UserName)

		conn, err := grpcClientConnection()
		defer conn.Close()
		registerServiceClient := register.NewRegisterServiceClient(conn)

		resp, err := registerServiceClient.RegisterCustomer(context.Background(), &customerDetails)
		if err != nil {
			log.Printf("Error creating user %v", err)
		}
		log.Printf("Response from Server %v", resp)
		fmt.Fprintf(w, resp.Message)
	}

}

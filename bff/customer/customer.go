package customer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

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

}

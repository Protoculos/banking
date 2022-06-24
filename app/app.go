package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	// define router
	r.HandleFunc("/greet", greet).Methods(http.MethodGet)
	r.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	r.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

package app

import (
	"log"
	"net/http"

	"github.com/Protoculos/banking/domain"
	"github.com/Protoculos/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define router
	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

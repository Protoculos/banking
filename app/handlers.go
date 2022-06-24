package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/Protoculos/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name,omitempty" xml:"name,omitempty"`
	City    string `json:"city,omitempty" xml:"city,omitempty"`
	Zipcode string `json:"zipcode,omitempty" xml:"zipcode,omitempty"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		if err := xml.NewEncoder(w).Encode(customers); err != nil {
			fmt.Fprint(w, err)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(customers); err != nil {
			fmt.Fprint(w, err)
		}
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name,omitempty" xml:"name,omitempty"`
	City    string `json:"city,omitempty" xml:"city,omitempty"`
	Zipcode string `json:"zipcode,omitempty" xml:"zipcode,omitempty"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashish", City: "New Delhi", Zipcode: "110075"},
		{Name: "Rob", City: "New Delhi", Zipcode: "110075"},
	}
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

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

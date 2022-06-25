package domain

import "github.com/Protoculos/banking/errs"

type Customer struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	City        string `json:"city,omitempty"`
	Zipcode     string `json:"zipcode,omitempty"`
	DateofBirth string `json:"dateof_birth,omitempty"`
	Status      string `json:"status,omitempty"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}

package domain

type CustomerRepositoryStab struct {
	customers []Customer
}

func (s CustomerRepositoryStab) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStab() CustomerRepositoryStab {
	customers := []Customer{
		{ID: "1001", Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{ID: "1002", Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStab{customers: customers}
}

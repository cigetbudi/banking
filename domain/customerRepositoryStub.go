package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Sigit", City: "Jakarta", ZipCode: "56161", DateOfBirth: "1995-01-05", Status: "1"},
		{Id: "1002", Name: "Syifa", City: "Kuningan", ZipCode: "45551", DateOfBirth: "1997-08-03", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}

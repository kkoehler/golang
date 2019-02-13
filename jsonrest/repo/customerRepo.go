package repo

import (
	"github.com/kkoehler/golang/jsonrest/model"
)

type CustomerRepository interface {
	Get(id int) *model.Customer
	Add(customer *model.Customer)
}

type ArrayCustomerRepository struct {
	customers []model.Customer
}

func (repo *ArrayCustomerRepository) Init() {
	repo.customers = []model.Customer{model.Customer{Id: 1, Firstname: "Dummy"}}
}

func (repo *ArrayCustomerRepository) Get(id int) *model.Customer {

	for _, customer := range repo.customers {
		if customer.Id == id {
			return &customer
		}
	}

	return nil

}

func (repo *ArrayCustomerRepository) Add(customer *model.Customer) {

	repo.customers = append(repo.customers, *customer)

}

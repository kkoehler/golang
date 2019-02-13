package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/kkoehler/golang/jsonrest/model"

	"github.com/kkoehler/golang/jsonrest/repo"
)

type CustomerController struct {
	repo repo.CustomerRepository
}

func NewController(repo repo.CustomerRepository) CustomerController {
	return CustomerController{repo: repo}
}

func (controller *CustomerController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	path := request.URL.Path
	idx := strings.LastIndex(path, "/") + 1
	customerId, err := strconv.Atoi(path[idx:])

	if request.Method == "GET" {

		if err != nil {
			writer.WriteHeader(400)
			writer.Write([]byte("Not found"))
			return
		}

		customer := controller.repo.Get(customerId)

		if customer == nil {
			writer.WriteHeader(404)
			writer.Write([]byte("Not found"))
			return
		}

		json, _ := json.Marshal(customer)
		header := writer.Header()
		header.Add("content-type", "application/json")
		writer.Write(json)

	}

	if request.Method == "PUT" {

		customer := &model.Customer{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(customer)

		controller.repo.Add(customer)

		writer.WriteHeader(204)
		writer.Write([]byte("Not found"))

	}

}

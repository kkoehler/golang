package main

import (
	"log"
	"net/http"

	"github.com/kkoehler/golang/jsonrest/middleware"

	"github.com/kkoehler/golang/jsonrest/repo"

	"github.com/kkoehler/golang/jsonrest/controller"
)

func main() {

	port := ":8080"

	log.Printf("Listen on port %s", port)

	repository := repo.ArrayCustomerRepository{}
	repository.Init()

	customerController := controller.NewController(&repository)

	http.Handle("/customer/", middleware.SecurityMiddleware(&customerController))

	log.Fatal(http.ListenAndServe(port, nil))

}

package main

import (
	"context"
	"fmt"

	hello "github.com/kkoehler/golang/go-micro/client/hello"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("hello.client"))
	service.Init()

	// Create new greeter client
	greeter := hello.NewHelloService("hello", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &hello.HelloRequest{Name: "Kristian"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

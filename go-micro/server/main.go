package main

import (
	"context"
	"fmt"

	hello "github.com/kkoehler/golang/go-micro/server/hello"
	micro "github.com/micro/go-micro"
)

type HelloService struct{}

func (g *HelloService) Hello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("hello"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	hello.RegisterHelloServiceHandler(service.Server(), new(HelloService))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

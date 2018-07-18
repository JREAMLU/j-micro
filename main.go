package main

import (
	"github.com/JREAMLU/j-micro/controller"
	proto "github.com/JREAMLU/j-micro/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	RunMicroService()
}

// RunMicroService run micro service
func RunMicroService() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), controller.NewGreeterHandler())

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

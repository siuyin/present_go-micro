// 10 OMIT
package main

import (
	"context"
	"fmt"
	"log"

	micro "github.com/micro/go-micro"

	proto "gomicrotry/proto" // HL
)

// 20 OMIT
// 30 OMIT
type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context,
	req *proto.HelloRequest,
	rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	log.Println("served request")
	return nil
}

// 40 OMIT
// 50 OMIT
func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// 60 OMIT

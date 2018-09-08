// 10 OMIT
package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"

	proto "gomicrotry/proto" // HL
)

// 20 OMIT
// 30 OMIT
func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client()) // (1) // HL

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "SiuYin"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

// 40 OMIT

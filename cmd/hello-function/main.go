//10 OMIT
package main

import (
	"context"

	"github.com/micro/go-micro"

	proto "gomicrotry/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest,
	rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello Function: " + req.Name // (1) // HL
	return nil
}

//20 OMIT
//30 OMIT
func main() {
	fnc := micro.NewFunction(
		micro.Name("greeter-func"), // (2) // HL
	)

	fnc.Init()

	fnc.Handle(new(Greeter))

	fnc.Run()
}

//40 OMIT

package main

import (
	"fmt"
	"proto/greeter"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)


type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	fmt.Printf("Responding with %s\n", rsp.Greeting)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type":"helloworld"
		})

		service.Init()
		
		// Register handler
		proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	
		// Run the server
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	)
}
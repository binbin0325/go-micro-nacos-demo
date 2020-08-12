package main

import (
	"context"
	helloworld "go-micro-demo/proto"
	"nacos"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Hello(ctx context.Context, req *helloworld.HelloRequest, rsp *helloworld.HelloResponse) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}
func main() {
	addrs := make([]string, 1)
	addrs[0] = "192.168.23.178:8848"
	registry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = addrs
	})
	service := micro.NewService(
		// Set service name
		micro.Name("my.service"),
		// Set service registry
		micro.Registry(registry),
	)
	helloworld.RegisterGreeterHandler(service.Server(), new(Helloworld))
	service.Run()
}

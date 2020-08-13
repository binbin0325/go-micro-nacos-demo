package main

import (
	"context"
	"fmt"
	helloworld "go-micro-demo/proto"
	"nacos"

	_ "nacos"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
)

func main() {
	addrs := make([]string, 1)
	addrs[0] = "192.168.31.57:8848"
	registry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = addrs
	})
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("my.service.client"),
		micro.Registry(registry))

	// 创建新的客户端
	greeter := helloworld.NewGreeterService("my.service", service.Client())
	// 调用greeter
	rsp, err := greeter.Hello(context.TODO(), &helloworld.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}
	//获取所有服务
	fmt.Println(registry.ListServices())
	//获取某一个服务
	services, err := registry.GetService("my.service")
	if err != nil {
		fmt.Println(err)
	}
	//监听服务
	registry.Watch()
	fmt.Println(services)
	// 打印响应请求
	fmt.Println(rsp.Greeting)
	service.Run()
}

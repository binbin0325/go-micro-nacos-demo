package main

import (
	"context"
	"fmt"
	helloworld "go-micro-nacos-demo/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	nacos "github.com/micro/go-plugins/registry/nacos/v2"
)

const serverName = "my.micro.service"

func main() {
	addrs := make([]string, 1)
	addrs[0] = "console.nacos.io:80"
	r := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = addrs
	})
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("my.micro.service.client"),
		micro.Registry(r))

	// 创建新的客户端
	greeter := helloworld.NewGreeterService(serverName, service.Client())
	// 调用greeter
	rsp, err := greeter.Hello(context.TODO(), &helloworld.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}
	//获取所有服务
	fmt.Println(registry.ListServices())
	//获取某一个服务
	services, err := registry.GetService(serverName)
	if err != nil {
		fmt.Println(err)
	}

	//监听服务
	watch, err := registry.Watch()

	fmt.Println(services)
	// 打印响应请求
	fmt.Println(rsp.Greeting)
	go service.Run()
	for {
		result, err := watch.Next()
		if len(result.Action) > 0 {
			fmt.Println(result, err)
		}
	}

}

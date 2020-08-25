# go-micro-nacos-demo

此项目可以帮助你在go-micro中使用nacos.

Nacos:https://nacos.io/en-us/

Go-Micro:https://github.com/micro/go-micro

## 项目说明

此项目利用go-micro创建服务端(server.go)以及客户端(client.go).

1. 在服务端中使用protobuf文件定义了一个服务叫做Greeter的处理器,它有一个接收HelloRequest并返回HelloResponse的Hello方法。并将服务端注册到nacos

   ```protobuf
   syntax = "proto3";
   
   package helloworld;
   
   service Greeter {
       rpc Hello(HelloRequest) returns (HelloResponse) {}
   }
   
   message HelloRequest {
       string name = 1;
   }
   
   message HelloResponse {
       string greeting = 2;
   }
   ```

2. 在客户端完成服务注册,客户端调用Hello,获取服务列表,获取单个服务,以及监听服务的功能。

   ps:(本项目中client.go 使用的nacos客户端是go-mirco提供的是默认配置。go-micro 的registry 接口自由度较高,我们可以利用context完成nacos客户端参数的配置, 如何使用context设置nacos客户端参数 可参考:https://github.com/micro/go-plugins/blob/master/registry/nacos/nacos_test.go)

## 功能说明

1. server.go

   1. 服务端:使用go-micro创建服务端Demo,并注册到nacos

      ```go
         registry := nacos.NewRegistry(func(options *registry.Options) {
         		options.Addrs = addrs
         })
         service := micro.NewService(
         		// Set service name
         		micro.Name("my.micro.service"),
         		// Set service registry
         		micro.Registry(registry),
         )
         service.Run()
      
      
      ```

      

2. client.go

   1. 客户端:使用go-micro创建客户端Demo,注册到nacos.

      ```go
      	r := nacos.NewRegistry(func(options *registry.Options) {
      		options.Addrs = addrs
      	})
      	service := micro.NewService(
      		micro.Name("my.micro.service.client"),
      		micro.Registry(r))
      ```

   2. 客户端rpc调用

      ```go
      	// 创建新的客户端
      	greeter := helloworld.NewGreeterService(serverName, service.Client())
      	// 调用greeter
      	rsp, err := greeter.Hello(context.TODO(), &helloworld.HelloRequest{Name: "John"})
      ```

   3. 查询服务列表

      ```go
      	services,err:=registry.ListServices()
      ```

   4. 获取某一个服务

      ```go
      	service, err := registry.GetService(serverName)
      ```

   5. 监听服务

      ```go
      	//监听服务
      	watch, err := registry.Watch()
      	for {
      		result, err := watch.Next()
      		if len(result.Action) > 0 {
      			fmt.Println(result, err)
      		}
      	}
      ```

   ## Run

1. ```
   #clone项目
   $git clone git@github.com:sanxun0325/go-micro-nacos-demo.git
   ```

2. ```
   #启动服务端
   $go run server.go
   
   在控制台可看到打印如下日志：
   2020-08-25 20:11:52  file=v2@v2.9.1/service.go:200 level=info Starting [service] my.micro.service
   2020-08-25 20:11:52  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:64604
   2020-08-25 20:11:52  file=grpc/grpc.go:697 level=info Registry [nacos] Registering node: my.micro.service-c23d867d-4752-4dc2-8944-3098a1d2c404
   
   ```

3. ```
   #启动客户端
   $go run client.go
   
   在控制台可看到打印如下日志：
   2020-08-25 20:18:35  file=v2@v2.9.1/service.go:200 level=info Starting [service] my.micro.service.client
   2020-08-25 20:18:35  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:64957
   2020-08-25 20:18:35  file=grpc/grpc.go:697 level=info Registry [nacos] Registering node: my.micro.service.client-ae7d2a0b-3758-49ee-8bd3-d30377a416e4
   
   此时服务端控制台追加打印(服务调用成功日志)：
   2020-08-25 20:18:35  file=go-micro-demo/server.go:17 level=info Received Helloworld.Call request
   ```

4. 在nacos中可以看到 客户端和服务端都已经注册在服务列表中。
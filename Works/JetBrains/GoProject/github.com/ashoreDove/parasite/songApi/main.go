package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/ashoreDove/common"
	go_micro_service_song "github.com/ashoreDove/parasite-song/proto/song"
	"github.com/ashoreDove/parasite-songApi/handler"
	songApi "github.com/ashoreDove/parasite-songApi/proto/songApi"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
)

func main() {
	//注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	},
	)
	//链路追踪
	t, io, err := common.NewTracer("go.micro.api.songApi", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动端口监听
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9097"),
			hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.songApi"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("192.168.0.106:8087"),
		//添加consul 作为注册中心
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()
	// Register Handler
	us := go_micro_service_song.NewSongService("go.micro.service.song", service.Client())
	// Register Handler
	err = songApi.RegisterSongApiHandler(service.Server(), &handler.SongApi{SongModuleService: us})
	if err != nil {
		return
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		//错误处理
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{
			c,
		}
	}
}

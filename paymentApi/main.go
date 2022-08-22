package main

import (
	"context"
	"git.imooc.com/keegan/common"
	go_micro_service_payment "git.imooc.com/keegan/payment/proto/payment"
	"git.imooc.com/keegan/paymentApi/handler"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"

	paymentApi "git.imooc.com/keegan/paymentApi/proto/paymentApi"
)
//sb-xpx7b8299280@personal.example.com
//5qe:hf=I
func main() {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// jaeger 链路追踪
	t,io,err := common.NewTracer("go.micro.api.payment","localhost:6831")
	if err !=nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	//启动监听
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0","9192"),hystrixStreamHandler)
		common.Error(err)
		log.Error(err)
	}()

	//监控
	common.PrometheusBoot(9292)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.paymentApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:9092"),
		//注册中心
		micro.Registry(consul),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//作为服务端访问时候生效
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	paymentService := go_micro_service_payment.NewPaymentService("go.micro.service.payment",service.Client())

	// Register Handler
	if err:=paymentApi.RegisterPaymentApiHandler(service.Server(), &handler.PaymentApi{PaymentService:paymentService});err!=nil {
		common.Error(err)
	}

	// Run service

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context,req client.Request,rsp interface{},opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//正常执行
		common.Info(req.Service()+"."+req.Endpoint())
		return c.Client.Call(ctx,req,rsp,opts...)
	}, func(e error) error {
		common.Error(e)
		return e
	})
}

func NewClientHystrixWrapper() client.Wrapper  {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}

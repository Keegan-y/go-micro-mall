package main

import (
	"context"
	"fmt"
	"git.imooc.com/keegan/product/common"
	go_micro_service_product "git.imooc.com/keegan/product/proto/product"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main()  {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t,io,err:=common.NewTracer("go.micro.service.product.client","localhost:6831")
	if err !=nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService:=go_micro_service_product.NewProductService("go.micro.service.product",service.Client())

	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:        "imooc",
		ProductSku:         "cap",
		ProductPrice:       1.1,
		ProductDescription: "imooc-cap",
		ProductCategoryId:  1,
		ProductImage:       []*go_micro_service_product.ProductImage{
			{
				ImageName: "cap-image",
				ImageCode: "capimage01",
				ImageUrl:  "capimage01",
			},
			{
				ImageName: "cap-image02",
				ImageCode: "capimage02",
				ImageUrl:  "capimage02",
			},
		},
		ProductSize:        []*go_micro_service_product.ProductSize{
			{
				SizeName: "cap-size",
				SizeCode: "cap-size-code",
			},
		},
		ProductSeo:         &go_micro_service_product.ProductSeo{
			SeoTitle:       "cap-seo",
			SeoKeywords:    "cap-seo",
			SeoDescription: "cap-seo",
			SeoCode:        "cap-seo",
		},
	}
	response,err:=productService.AddProduct(context.TODO(),productAdd)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
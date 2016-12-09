package main

import (
	"github.com/facebookgo/inject"
	"github.com/micro/go-micro"
	"incremental/config/di"
	"incremental/handler"
	proto "incremental/proto/incremental"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("version_incremental"),
		micro.Version("latest"),
	)

	service.Init()

	var graph inject.Graph
	handler := new(handler.IncrementalVersionHandler)
	di.InitServices(graph, handler)
	proto.RegisterVersionIncrementalHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/facebookgo/inject"
	"github.com/micro/go-micro"
	"log"
	"project/config/di"
	"project/handler"
	proto "project/proto/project"
)

func main() {
	service := micro.NewService(
		micro.Name("project"),
		micro.Version("latest"),
	)

	service.Init()

	var graph inject.Graph
	handler := new(handler.ProjectHandler)
	di.InitServices(graph, handler)
	proto.RegisterProjectHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

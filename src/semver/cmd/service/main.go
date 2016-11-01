package main

import (
	"github.com/facebookgo/inject"
	"github.com/micro/go-micro"
	"log"
	"semver/config/di"
	"semver/handler"
	proto "semver/proto/semver"
)

func main() {
	service := micro.NewService(
		micro.Name("semver"),
		micro.Version("latest"),
	)

	service.Init()

	var graph inject.Graph
	handler := new(handler.SemverHandler)
	di.InitServices(graph, handler)
	proto.RegisterVersionSemverHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

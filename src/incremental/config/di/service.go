package di

import (
	"fmt"
	"os"

	"github.com/facebookgo/inject"
	"incremental/config"
	"incremental/database"
	"incremental/generator"
	"incremental/handler"
	"incremental/model"
)

//InitServices initialize DI configuration
//DIC
func InitServices(graph inject.Graph, handler *handler.IncrementalVersionHandler) {

	config := config.New()
	dbMap := database.InitORM(config)

	err := graph.Provide(
		&inject.Object{Value: dbMap, Name: "dbMap"},
		&inject.Object{Value: config, Name: "config"},
		&inject.Object{Value: &model.VersionRepository{}},
		&inject.Object{Value: &generator.IncrementalVersionGenerator{}},
		&inject.Object{Value: handler},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := graph.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

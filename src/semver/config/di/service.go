package di

import (
	"fmt"
	"os"

	"github.com/facebookgo/inject"
	"semver/config"
	"semver/database"
	"semver/generator"
	"semver/handler"
	"semver/model"
)

//InitServices initialize DI configuration
//DIC
func InitServices(graph inject.Graph, handler *handler.SemverHandler) {

	config := config.New()
	dbMap := database.InitORM(config)

	err := graph.Provide(
		&inject.Object{Value: dbMap, Name: "dbMap"},
		&inject.Object{Value: config, Name: "config"},
		&inject.Object{Value: &model.VersionRepository{}},
		&inject.Object{Value: &generator.SemverGenerator{}},
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

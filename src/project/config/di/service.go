package di

import (
	"fmt"
	"os"

	"github.com/facebookgo/inject"
	"project/config"
	"project/database"
	"project/handler"
	"project/model"
)

//InitServices initialize DI configuration
//DIC
func InitServices(graph inject.Graph, handler *handler.ProjectHandler) {

	config := config.New()
	dbMap := database.NewConnection(config)

	err := graph.Provide(
		&inject.Object{Value: dbMap, Name: "dbMap"},
		&inject.Object{Value: config, Name: "config"},
		&inject.Object{Value: &model.ProjectRepository{}},
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

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"Api/micro_gateway"
	"Api/restapi/operations"
	"Api/restapi/operations/project"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.ReleaseManagerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ReleaseManagerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	projectGateway := micro_gateway.NewProjectGateway()
	api.ProjectCreateProjectHandler = project.CreateProjectHandlerFunc(func(params project.CreateProjectParams) middleware.Responder {
		return projectGateway.CreateProjectAction(params)
	})
	api.ProjectReadProjectHandler = project.ReadProjectHandlerFunc(func(params project.ReadProjectParams) middleware.Responder {
		return projectGateway.ReadProjectAction(params)
	})
	api.ProjectUpdateProjectHandler = project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams) middleware.Responder {
		return projectGateway.UpdateProjectAction(params)
	})
	api.ProjectListProjectsHandler = project.ListProjectsHandlerFunc(func(params project.ListProjectsParams) middleware.Responder {
		return projectGateway.ListProjectsAction(params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

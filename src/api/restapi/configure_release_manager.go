package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"api/microGateway"
	"api/restapi/operations"
	"api/restapi/operations/project"
	"api/restapi/operations/version_incremental"
	"api/restapi/operations/version_semantic"
	"github.com/tylerb/graceful"
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

	//PROJECT MICRO SERVICE INTEGRATION
	projectGateway := microGateway.NewProjectGateway()
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

	//SEMANTIC VERSIONING MICRO SERVICE INTEGRATION
	semverGateway := microGateway.NewSemverGateway()
	api.VersionSemanticSemverGenerateHandler = version_semantic.SemverGenerateHandlerFunc(func(params version_semantic.SemverGenerateParams) middleware.Responder {
		return semverGateway.GenerateVersionAction(params)
	})

	//INCREMENTAL VERSIONING MICRO SERVICE INTEGRATION
	incrementalGateway := microGateway.NewIncrementalVersioningGateway()
	api.VersionIncrementalIncrementalGenerateHandler = version_incremental.IncrementalGenerateHandlerFunc(func(params version_incremental.IncrementalGenerateParams) middleware.Responder {
		return incrementalGateway.GenerateVersionAction(params)
	})

	api.VersionIncrementalIncrementalDeleteHandler = version_incremental.IncrementalDeleteHandlerFunc(func(params version_incremental.IncrementalDeleteParams) middleware.Responder {
		return incrementalGateway.DeleteVersionAction(params)
	})

	api.VersionIncrementalIncrementalUpdateHandler = version_incremental.IncrementalUpdateHandlerFunc(func(params version_incremental.IncrementalUpdateParams) middleware.Responder {
		return incrementalGateway.UpdateVersionAction(params)
	})

	//END OF MICRO SERVICES INTEGRATION

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}


// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
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

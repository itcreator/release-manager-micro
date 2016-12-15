package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"api/restapi/operations/project"
	"api/restapi/operations/version_incremental"
	"api/restapi/operations/version_semantic"
)

// NewReleaseManagerAPI creates a new ReleaseManager instance
func NewReleaseManagerAPI(spec *loads.Document) *ReleaseManagerAPI {
	o := &ReleaseManagerAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*ReleaseManagerAPI This application generate version for your software. */
type ReleaseManagerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/release-manager.v1+json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/release-manager.v1+json" mime type
	JSONProducer runtime.Producer

	// ProjectCreateProjectHandler sets the operation handler for the create project operation
	ProjectCreateProjectHandler project.CreateProjectHandler
	// VersionIncrementalIncrementalDeleteHandler sets the operation handler for the incremental delete operation
	VersionIncrementalIncrementalDeleteHandler version_incremental.IncrementalDeleteHandler
	// VersionIncrementalIncrementalGenerateHandler sets the operation handler for the incremental generate operation
	VersionIncrementalIncrementalGenerateHandler version_incremental.IncrementalGenerateHandler
	// VersionIncrementalIncrementalUpdateHandler sets the operation handler for the incremental update operation
	VersionIncrementalIncrementalUpdateHandler version_incremental.IncrementalUpdateHandler
	// ProjectListProjectsHandler sets the operation handler for the list projects operation
	ProjectListProjectsHandler project.ListProjectsHandler
	// ProjectReadProjectHandler sets the operation handler for the read project operation
	ProjectReadProjectHandler project.ReadProjectHandler
	// VersionSemanticSemverGenerateHandler sets the operation handler for the semver generate operation
	VersionSemanticSemverGenerateHandler version_semantic.SemverGenerateHandler
	// ProjectUpdateProjectHandler sets the operation handler for the update project operation
	ProjectUpdateProjectHandler project.UpdateProjectHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ReleaseManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ReleaseManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *ReleaseManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ReleaseManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ReleaseManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ReleaseManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ReleaseManagerAPI
func (o *ReleaseManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ProjectCreateProjectHandler == nil {
		unregistered = append(unregistered, "project.CreateProjectHandler")
	}

	if o.VersionIncrementalIncrementalDeleteHandler == nil {
		unregistered = append(unregistered, "version_incremental.IncrementalDeleteHandler")
	}

	if o.VersionIncrementalIncrementalGenerateHandler == nil {
		unregistered = append(unregistered, "version_incremental.IncrementalGenerateHandler")
	}

	if o.VersionIncrementalIncrementalUpdateHandler == nil {
		unregistered = append(unregistered, "version_incremental.IncrementalUpdateHandler")
	}

	if o.ProjectListProjectsHandler == nil {
		unregistered = append(unregistered, "project.ListProjectsHandler")
	}

	if o.ProjectReadProjectHandler == nil {
		unregistered = append(unregistered, "project.ReadProjectHandler")
	}

	if o.VersionSemanticSemverGenerateHandler == nil {
		unregistered = append(unregistered, "version_semantic.SemverGenerateHandler")
	}

	if o.ProjectUpdateProjectHandler == nil {
		unregistered = append(unregistered, "project.UpdateProjectHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ReleaseManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ReleaseManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ReleaseManagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/release-manager.v1+json":
			result["application/release-manager.v1+json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ReleaseManagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/release-manager.v1+json":
			result["application/release-manager.v1+json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ReleaseManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *ReleaseManagerAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects"] = project.NewCreateProject(o.context, o.ProjectCreateProjectHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/increamental_version/{projectName}"] = version_incremental.NewIncrementalDelete(o.context, o.VersionIncrementalIncrementalDeleteHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/increamental_version/{projectName}"] = version_incremental.NewIncrementalGenerate(o.context, o.VersionIncrementalIncrementalGenerateHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/increamental_version/{projectName}"] = version_incremental.NewIncrementalUpdate(o.context, o.VersionIncrementalIncrementalUpdateHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects"] = project.NewListProjects(o.context, o.ProjectListProjectsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{id}"] = project.NewReadProject(o.context, o.ProjectReadProjectHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/version/semantic"] = version_semantic.NewSemverGenerate(o.context, o.VersionSemanticSemverGenerateHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{id}"] = project.NewUpdateProject(o.context, o.ProjectUpdateProjectHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ReleaseManagerAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}

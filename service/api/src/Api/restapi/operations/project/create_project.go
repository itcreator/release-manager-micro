package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateProjectHandlerFunc turns a function with the right signature into a create project handler
type CreateProjectHandlerFunc func(CreateProjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateProjectHandlerFunc) Handle(params CreateProjectParams) middleware.Responder {
	return fn(params)
}

// CreateProjectHandler interface for that can handle valid create project params
type CreateProjectHandler interface {
	Handle(CreateProjectParams) middleware.Responder
}

// NewCreateProject creates a new http.Handler for the create project operation
func NewCreateProject(ctx *middleware.Context, handler CreateProjectHandler) *CreateProject {
	return &CreateProject{Context: ctx, Handler: handler}
}

/*CreateProject swagger:route POST /projects project createProject

Create new projects

*/
type CreateProject struct {
	Context *middleware.Context
	Handler CreateProjectHandler
}

func (o *CreateProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateProjectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

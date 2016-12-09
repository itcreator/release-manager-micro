package version_incremental

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// IncrementalDeleteHandlerFunc turns a function with the right signature into a incremental delete handler
type IncrementalDeleteHandlerFunc func(IncrementalDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IncrementalDeleteHandlerFunc) Handle(params IncrementalDeleteParams) middleware.Responder {
	return fn(params)
}

// IncrementalDeleteHandler interface for that can handle valid incremental delete params
type IncrementalDeleteHandler interface {
	Handle(IncrementalDeleteParams) middleware.Responder
}

// NewIncrementalDelete creates a new http.Handler for the incremental delete operation
func NewIncrementalDelete(ctx *middleware.Context, handler IncrementalDeleteHandler) *IncrementalDelete {
	return &IncrementalDelete{Context: ctx, Handler: handler}
}

/*IncrementalDelete swagger:route DELETE /increamental_version/{projectName} versionIncremental incrementalDelete

Delete incremental version number (RESET)

Delete generated version


*/
type IncrementalDelete struct {
	Context *middleware.Context
	Handler IncrementalDeleteHandler
}

func (o *IncrementalDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewIncrementalDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// Code generated by go-swagger; DO NOT EDIT.

package version_incremental

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIncrementalDeleteParams creates a new IncrementalDeleteParams object
// with the default values initialized.
func NewIncrementalDeleteParams() *IncrementalDeleteParams {
	var ()
	return &IncrementalDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIncrementalDeleteParamsWithTimeout creates a new IncrementalDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIncrementalDeleteParamsWithTimeout(timeout time.Duration) *IncrementalDeleteParams {
	var ()
	return &IncrementalDeleteParams{

		timeout: timeout,
	}
}

// NewIncrementalDeleteParamsWithContext creates a new IncrementalDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewIncrementalDeleteParamsWithContext(ctx context.Context) *IncrementalDeleteParams {
	var ()
	return &IncrementalDeleteParams{

		Context: ctx,
	}
}

// NewIncrementalDeleteParamsWithHTTPClient creates a new IncrementalDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIncrementalDeleteParamsWithHTTPClient(client *http.Client) *IncrementalDeleteParams {
	var ()
	return &IncrementalDeleteParams{
		HTTPClient: client,
	}
}

/*IncrementalDeleteParams contains all the parameters to send to the API endpoint
for the incremental delete operation typically these are written to a http.Request
*/
type IncrementalDeleteParams struct {

	/*ProjectName*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the incremental delete params
func (o *IncrementalDeleteParams) WithTimeout(timeout time.Duration) *IncrementalDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the incremental delete params
func (o *IncrementalDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the incremental delete params
func (o *IncrementalDeleteParams) WithContext(ctx context.Context) *IncrementalDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the incremental delete params
func (o *IncrementalDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the incremental delete params
func (o *IncrementalDeleteParams) WithHTTPClient(client *http.Client) *IncrementalDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the incremental delete params
func (o *IncrementalDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectName adds the projectName to the incremental delete params
func (o *IncrementalDeleteParams) WithProjectName(projectName string) *IncrementalDeleteParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the incremental delete params
func (o *IncrementalDeleteParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *IncrementalDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
package version_incremental

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"api/models"
)

// NewIncrementalUpdateParams creates a new IncrementalUpdateParams object
// with the default values initialized.
func NewIncrementalUpdateParams() IncrementalUpdateParams {
	var ()
	return IncrementalUpdateParams{}
}

// IncrementalUpdateParams contains all the bound params for the incremental update operation
// typically these are obtained from a http.Request
//
// swagger:parameters incrementalUpdate
type IncrementalUpdateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  In: body
	*/
	Body *models.IncrementalVersionNumber
	/*
	  Required: true
	  Max Length: 100
	  In: path
	*/
	ProjectName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *IncrementalUpdateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.IncrementalVersionNumber
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	rProjectName, rhkProjectName, _ := route.Params.GetOK("projectName")
	if err := o.bindProjectName(rProjectName, rhkProjectName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IncrementalUpdateParams) bindProjectName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ProjectName = raw

	if err := o.validateProjectName(formats); err != nil {
		return err
	}

	return nil
}

func (o *IncrementalUpdateParams) validateProjectName(formats strfmt.Registry) error {

	if err := validate.MaxLength("projectName", "path", string(o.ProjectName), 100); err != nil {
		return err
	}

	return nil
}

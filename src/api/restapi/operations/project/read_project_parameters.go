// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewReadProjectParams creates a new ReadProjectParams object
// no default values defined in spec.
func NewReadProjectParams() ReadProjectParams {

	return ReadProjectParams{}
}

// ReadProjectParams contains all the bound params for the read project operation
// typically these are obtained from a http.Request
//
// swagger:parameters readProject
type ReadProjectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Project ID in UUID format
	  Required: true
	  In: path
	*/
	UUID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReadProjectParams() beforehand.
func (o *ReadProjectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUUID, rhkUUID, _ := route.Params.GetOK("uuid")
	if err := o.bindUUID(rUUID, rhkUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReadProjectParams) bindUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("uuid", "path", "strfmt.UUID", raw)
	}
	o.UUID = *(value.(*strfmt.UUID))

	if err := o.validateUUID(formats); err != nil {
		return err
	}

	return nil
}

func (o *ReadProjectParams) validateUUID(formats strfmt.Registry) error {

	if err := validate.FormatOf("uuid", "path", "uuid", o.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

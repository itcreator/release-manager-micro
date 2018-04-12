// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
)

// ReadProjectURL generates an URL for the read project operation
type ReadProjectURL struct {
	UUID strfmt.UUID

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ReadProjectURL) WithBasePath(bp string) *ReadProjectURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ReadProjectURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ReadProjectURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/projects/{uuid}"

	uuid := o.UUID.String()
	if uuid != "" {
		_path = strings.Replace(_path, "{uuid}", uuid, -1)
	} else {
		return nil, errors.New("UUID is required on ReadProjectURL")
	}

	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ReadProjectURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ReadProjectURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ReadProjectURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ReadProjectURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ReadProjectURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ReadProjectURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

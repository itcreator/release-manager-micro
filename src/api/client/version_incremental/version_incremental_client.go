// Code generated by go-swagger; DO NOT EDIT.

package version_incremental

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new version incremental API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for version incremental API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
IncrementalDelete deletes incremental version number r e s e t

Delete generated version

*/
func (a *Client) IncrementalDelete(params *IncrementalDeleteParams) (*IncrementalDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIncrementalDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "incrementalDelete",
		Method:             "DELETE",
		PathPattern:        "/increamental_version/{projectName}",
		ProducesMediaTypes: []string{"application/release-manager.v1+json"},
		ConsumesMediaTypes: []string{"application/release-manager.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IncrementalDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IncrementalDeleteNoContent), nil

}

/*
IncrementalGenerate generates new incremental version number

Incremental Versioning

*/
func (a *Client) IncrementalGenerate(params *IncrementalGenerateParams) (*IncrementalGenerateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIncrementalGenerateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "incrementalGenerate",
		Method:             "POST",
		PathPattern:        "/increamental_version/{projectName}",
		ProducesMediaTypes: []string{"application/release-manager.v1+json"},
		ConsumesMediaTypes: []string{"application/release-manager.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IncrementalGenerateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IncrementalGenerateCreated), nil

}

/*
IncrementalUpdate updates incremental version number only for maintenance

Incremental Versioning
Update revision number

*/
func (a *Client) IncrementalUpdate(params *IncrementalUpdateParams) (*IncrementalUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIncrementalUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "incrementalUpdate",
		Method:             "PUT",
		PathPattern:        "/increamental_version/{projectName}",
		ProducesMediaTypes: []string{"application/release-manager.v1+json"},
		ConsumesMediaTypes: []string{"application/release-manager.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IncrementalUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IncrementalUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
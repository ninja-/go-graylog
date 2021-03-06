// Code generated by go-swagger; DO NOT EDIT.

package system_sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system sessions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system sessions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
List lists current status of service manager
*/
func (a *Client) List(params *ListParams, authInfo runtime.ClientAuthInfoWriter) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list",
		Method:             "GET",
		PathPattern:        "/system/serviceManager",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

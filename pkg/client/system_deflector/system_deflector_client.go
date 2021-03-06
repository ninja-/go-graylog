// Code generated by go-swagger; DO NOT EDIT.

package system_deflector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system deflector API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system deflector API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GenerateClusterDebugEvent creates and send a cluster debug event
*/
func (a *Client) GenerateClusterDebugEvent(params *GenerateClusterDebugEventParams, authInfo runtime.ClientAuthInfoWriter) (*GenerateClusterDebugEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateClusterDebugEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "generateClusterDebugEvent",
		Method:             "POST",
		PathPattern:        "/system/debug/events/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenerateClusterDebugEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateClusterDebugEventOK), nil

}

/*
GenerateDebugEvent creates and send a local debug event
*/
func (a *Client) GenerateDebugEvent(params *GenerateDebugEventParams, authInfo runtime.ClientAuthInfoWriter) (*GenerateDebugEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateDebugEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "generateDebugEvent",
		Method:             "POST",
		PathPattern:        "/system/debug/events/local",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenerateDebugEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateDebugEventOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

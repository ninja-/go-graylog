// Code generated by go-swagger; DO NOT EDIT.

package search_universal_keyword

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search universal keyword API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search universal keyword API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAvailable returns all available message decorations
*/
func (a *Client) GetAvailable(params *GetAvailableParams, authInfo runtime.ClientAuthInfoWriter) (*GetAvailableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAvailableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAvailable",
		Method:             "GET",
		PathPattern:        "/search/decorators/available",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAvailableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAvailableOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

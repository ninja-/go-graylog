// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new messages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for messages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Index gets overview of current indexing state for the given index set including deflector config cluster state index ranges and message counts
*/
func (a *Client) Index(params *IndexParams, authInfo runtime.ClientAuthInfoWriter) (*IndexOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "index",
		Method:             "GET",
		PathPattern:        "/system/indexer/overview/{indexSetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IndexOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

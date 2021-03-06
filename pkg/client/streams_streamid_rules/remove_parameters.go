// Code generated by go-swagger; DO NOT EDIT.

package streams_streamid_rules

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

// NewRemoveParams creates a new RemoveParams object
// with the default values initialized.
func NewRemoveParams() *RemoveParams {
	var ()
	return &RemoveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveParamsWithTimeout creates a new RemoveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveParamsWithTimeout(timeout time.Duration) *RemoveParams {
	var ()
	return &RemoveParams{

		timeout: timeout,
	}
}

// NewRemoveParamsWithContext creates a new RemoveParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveParamsWithContext(ctx context.Context) *RemoveParams {
	var ()
	return &RemoveParams{

		Context: ctx,
	}
}

// NewRemoveParamsWithHTTPClient creates a new RemoveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveParamsWithHTTPClient(client *http.Client) *RemoveParams {
	var ()
	return &RemoveParams{
		HTTPClient: client,
	}
}

/*RemoveParams contains all the parameters to send to the API endpoint
for the remove operation typically these are written to a http.Request
*/
type RemoveParams struct {

	/*OutputID
	  The id of the output that should be deleted

	*/
	OutputID string
	/*Streamid
	  The id of the stream whose outputs we want.

	*/
	Streamid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove params
func (o *RemoveParams) WithTimeout(timeout time.Duration) *RemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove params
func (o *RemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove params
func (o *RemoveParams) WithContext(ctx context.Context) *RemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove params
func (o *RemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove params
func (o *RemoveParams) WithHTTPClient(client *http.Client) *RemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove params
func (o *RemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOutputID adds the outputID to the remove params
func (o *RemoveParams) WithOutputID(outputID string) *RemoveParams {
	o.SetOutputID(outputID)
	return o
}

// SetOutputID adds the outputId to the remove params
func (o *RemoveParams) SetOutputID(outputID string) {
	o.OutputID = outputID
}

// WithStreamid adds the streamid to the remove params
func (o *RemoveParams) WithStreamid(streamid string) *RemoveParams {
	o.SetStreamid(streamid)
	return o
}

// SetStreamid adds the streamid to the remove params
func (o *RemoveParams) SetStreamid(streamid string) {
	o.Streamid = streamid
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param outputId
	if err := r.SetPathParam("outputId", o.OutputID); err != nil {
		return err
	}

	// path param streamid
	if err := r.SetPathParam("streamid", o.Streamid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

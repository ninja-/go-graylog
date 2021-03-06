// Code generated by go-swagger; DO NOT EDIT.

package system

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

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// NewCloneStreamParams creates a new CloneStreamParams object
// with the default values initialized.
func NewCloneStreamParams() *CloneStreamParams {
	var ()
	return &CloneStreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloneStreamParamsWithTimeout creates a new CloneStreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloneStreamParamsWithTimeout(timeout time.Duration) *CloneStreamParams {
	var ()
	return &CloneStreamParams{

		timeout: timeout,
	}
}

// NewCloneStreamParamsWithContext creates a new CloneStreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloneStreamParamsWithContext(ctx context.Context) *CloneStreamParams {
	var ()
	return &CloneStreamParams{

		Context: ctx,
	}
}

// NewCloneStreamParamsWithHTTPClient creates a new CloneStreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloneStreamParamsWithHTTPClient(client *http.Client) *CloneStreamParams {
	var ()
	return &CloneStreamParams{
		HTTPClient: client,
	}
}

/*CloneStreamParams contains all the parameters to send to the API endpoint
for the clone stream operation typically these are written to a http.Request
*/
type CloneStreamParams struct {

	/*JSONBody*/
	JSONBody *models.CloneStreamRequest
	/*StreamID*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the clone stream params
func (o *CloneStreamParams) WithTimeout(timeout time.Duration) *CloneStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone stream params
func (o *CloneStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone stream params
func (o *CloneStreamParams) WithContext(ctx context.Context) *CloneStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone stream params
func (o *CloneStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone stream params
func (o *CloneStreamParams) WithHTTPClient(client *http.Client) *CloneStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone stream params
func (o *CloneStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the clone stream params
func (o *CloneStreamParams) WithJSONBody(jSONBody *models.CloneStreamRequest) *CloneStreamParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the clone stream params
func (o *CloneStreamParams) SetJSONBody(jSONBody *models.CloneStreamRequest) {
	o.JSONBody = jSONBody
}

// WithStreamID adds the streamID to the clone stream params
func (o *CloneStreamParams) WithStreamID(streamID string) *CloneStreamParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the clone stream params
func (o *CloneStreamParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *CloneStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param streamId
	if err := r.SetPathParam("streamId", o.StreamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

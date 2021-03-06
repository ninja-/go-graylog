// Code generated by go-swagger; DO NOT EDIT.

package system_indices_index_sets

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

// NewRebuildParams creates a new RebuildParams object
// with the default values initialized.
func NewRebuildParams() *RebuildParams {

	return &RebuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRebuildParamsWithTimeout creates a new RebuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRebuildParamsWithTimeout(timeout time.Duration) *RebuildParams {

	return &RebuildParams{

		timeout: timeout,
	}
}

// NewRebuildParamsWithContext creates a new RebuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewRebuildParamsWithContext(ctx context.Context) *RebuildParams {

	return &RebuildParams{

		Context: ctx,
	}
}

// NewRebuildParamsWithHTTPClient creates a new RebuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRebuildParamsWithHTTPClient(client *http.Client) *RebuildParams {

	return &RebuildParams{
		HTTPClient: client,
	}
}

/*RebuildParams contains all the parameters to send to the API endpoint
for the rebuild operation typically these are written to a http.Request
*/
type RebuildParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rebuild params
func (o *RebuildParams) WithTimeout(timeout time.Duration) *RebuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rebuild params
func (o *RebuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rebuild params
func (o *RebuildParams) WithContext(ctx context.Context) *RebuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rebuild params
func (o *RebuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rebuild params
func (o *RebuildParams) WithHTTPClient(client *http.Client) *RebuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rebuild params
func (o *RebuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RebuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

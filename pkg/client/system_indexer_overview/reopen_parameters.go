// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_overview

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

// NewReopenParams creates a new ReopenParams object
// with the default values initialized.
func NewReopenParams() *ReopenParams {
	var ()
	return &ReopenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReopenParamsWithTimeout creates a new ReopenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReopenParamsWithTimeout(timeout time.Duration) *ReopenParams {
	var ()
	return &ReopenParams{

		timeout: timeout,
	}
}

// NewReopenParamsWithContext creates a new ReopenParams object
// with the default values initialized, and the ability to set a context for a request
func NewReopenParamsWithContext(ctx context.Context) *ReopenParams {
	var ()
	return &ReopenParams{

		Context: ctx,
	}
}

// NewReopenParamsWithHTTPClient creates a new ReopenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReopenParamsWithHTTPClient(client *http.Client) *ReopenParams {
	var ()
	return &ReopenParams{
		HTTPClient: client,
	}
}

/*ReopenParams contains all the parameters to send to the API endpoint
for the reopen operation typically these are written to a http.Request
*/
type ReopenParams struct {

	/*Index*/
	Index string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reopen params
func (o *ReopenParams) WithTimeout(timeout time.Duration) *ReopenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reopen params
func (o *ReopenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reopen params
func (o *ReopenParams) WithContext(ctx context.Context) *ReopenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reopen params
func (o *ReopenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reopen params
func (o *ReopenParams) WithHTTPClient(client *http.Client) *ReopenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reopen params
func (o *ReopenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the reopen params
func (o *ReopenParams) WithIndex(index string) *ReopenParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the reopen params
func (o *ReopenParams) SetIndex(index string) {
	o.Index = index
}

// WriteToRequest writes these params to a swagger request
func (o *ReopenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

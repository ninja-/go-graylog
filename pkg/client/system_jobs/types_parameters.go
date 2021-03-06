// Code generated by go-swagger; DO NOT EDIT.

package system_jobs

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

// NewTypesParams creates a new TypesParams object
// with the default values initialized.
func NewTypesParams() *TypesParams {

	return &TypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTypesParamsWithTimeout creates a new TypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTypesParamsWithTimeout(timeout time.Duration) *TypesParams {

	return &TypesParams{

		timeout: timeout,
	}
}

// NewTypesParamsWithContext creates a new TypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewTypesParamsWithContext(ctx context.Context) *TypesParams {

	return &TypesParams{

		Context: ctx,
	}
}

// NewTypesParamsWithHTTPClient creates a new TypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTypesParamsWithHTTPClient(client *http.Client) *TypesParams {

	return &TypesParams{
		HTTPClient: client,
	}
}

/*TypesParams contains all the parameters to send to the API endpoint
for the types operation typically these are written to a http.Request
*/
type TypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the types params
func (o *TypesParams) WithTimeout(timeout time.Duration) *TypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the types params
func (o *TypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the types params
func (o *TypesParams) WithContext(ctx context.Context) *TypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the types params
func (o *TypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the types params
func (o *TypesParams) WithHTTPClient(client *http.Client) *TypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the types params
func (o *TypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

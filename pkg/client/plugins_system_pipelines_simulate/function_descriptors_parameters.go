// Code generated by go-swagger; DO NOT EDIT.

package plugins_system_pipelines_simulate

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

// NewFunctionDescriptorsParams creates a new FunctionDescriptorsParams object
// with the default values initialized.
func NewFunctionDescriptorsParams() *FunctionDescriptorsParams {

	return &FunctionDescriptorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionDescriptorsParamsWithTimeout creates a new FunctionDescriptorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFunctionDescriptorsParamsWithTimeout(timeout time.Duration) *FunctionDescriptorsParams {

	return &FunctionDescriptorsParams{

		timeout: timeout,
	}
}

// NewFunctionDescriptorsParamsWithContext creates a new FunctionDescriptorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFunctionDescriptorsParamsWithContext(ctx context.Context) *FunctionDescriptorsParams {

	return &FunctionDescriptorsParams{

		Context: ctx,
	}
}

// NewFunctionDescriptorsParamsWithHTTPClient creates a new FunctionDescriptorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFunctionDescriptorsParamsWithHTTPClient(client *http.Client) *FunctionDescriptorsParams {

	return &FunctionDescriptorsParams{
		HTTPClient: client,
	}
}

/*FunctionDescriptorsParams contains all the parameters to send to the API endpoint
for the function descriptors operation typically these are written to a http.Request
*/
type FunctionDescriptorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the function descriptors params
func (o *FunctionDescriptorsParams) WithTimeout(timeout time.Duration) *FunctionDescriptorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the function descriptors params
func (o *FunctionDescriptorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the function descriptors params
func (o *FunctionDescriptorsParams) WithContext(ctx context.Context) *FunctionDescriptorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the function descriptors params
func (o *FunctionDescriptorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the function descriptors params
func (o *FunctionDescriptorsParams) WithHTTPClient(client *http.Client) *FunctionDescriptorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the function descriptors params
func (o *FunctionDescriptorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionDescriptorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package system_configuration

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

// NewGetAllParams creates a new GetAllParams object
// with the default values initialized.
func NewGetAllParams() *GetAllParams {

	return &GetAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllParamsWithTimeout creates a new GetAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllParamsWithTimeout(timeout time.Duration) *GetAllParams {

	return &GetAllParams{

		timeout: timeout,
	}
}

// NewGetAllParamsWithContext creates a new GetAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllParamsWithContext(ctx context.Context) *GetAllParams {

	return &GetAllParams{

		Context: ctx,
	}
}

// NewGetAllParamsWithHTTPClient creates a new GetAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllParamsWithHTTPClient(client *http.Client) *GetAllParams {

	return &GetAllParams{
		HTTPClient: client,
	}
}

/*GetAllParams contains all the parameters to send to the API endpoint
for the get all operation typically these are written to a http.Request
*/
type GetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all params
func (o *GetAllParams) WithTimeout(timeout time.Duration) *GetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all params
func (o *GetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all params
func (o *GetAllParams) WithContext(ctx context.Context) *GetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all params
func (o *GetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all params
func (o *GetAllParams) WithHTTPClient(client *http.Client) *GetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all params
func (o *GetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

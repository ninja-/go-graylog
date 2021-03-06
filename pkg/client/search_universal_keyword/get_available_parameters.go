// Code generated by go-swagger; DO NOT EDIT.

package search_universal_keyword

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

// NewGetAvailableParams creates a new GetAvailableParams object
// with the default values initialized.
func NewGetAvailableParams() *GetAvailableParams {

	return &GetAvailableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvailableParamsWithTimeout creates a new GetAvailableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAvailableParamsWithTimeout(timeout time.Duration) *GetAvailableParams {

	return &GetAvailableParams{

		timeout: timeout,
	}
}

// NewGetAvailableParamsWithContext creates a new GetAvailableParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAvailableParamsWithContext(ctx context.Context) *GetAvailableParams {

	return &GetAvailableParams{

		Context: ctx,
	}
}

// NewGetAvailableParamsWithHTTPClient creates a new GetAvailableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAvailableParamsWithHTTPClient(client *http.Client) *GetAvailableParams {

	return &GetAvailableParams{
		HTTPClient: client,
	}
}

/*GetAvailableParams contains all the parameters to send to the API endpoint
for the get available operation typically these are written to a http.Request
*/
type GetAvailableParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get available params
func (o *GetAvailableParams) WithTimeout(timeout time.Duration) *GetAvailableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get available params
func (o *GetAvailableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get available params
func (o *GetAvailableParams) WithContext(ctx context.Context) *GetAvailableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get available params
func (o *GetAvailableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get available params
func (o *GetAvailableParams) WithHTTPClient(client *http.Client) *GetAvailableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get available params
func (o *GetAvailableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvailableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

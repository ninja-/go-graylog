// Code generated by go-swagger; DO NOT EDIT.

package system_loggers

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

// NewOverrideParams creates a new OverrideParams object
// with the default values initialized.
func NewOverrideParams() *OverrideParams {
	var ()
	return &OverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOverrideParamsWithTimeout creates a new OverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOverrideParamsWithTimeout(timeout time.Duration) *OverrideParams {
	var ()
	return &OverrideParams{

		timeout: timeout,
	}
}

// NewOverrideParamsWithContext creates a new OverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewOverrideParamsWithContext(ctx context.Context) *OverrideParams {
	var ()
	return &OverrideParams{

		Context: ctx,
	}
}

// NewOverrideParamsWithHTTPClient creates a new OverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOverrideParamsWithHTTPClient(client *http.Client) *OverrideParams {
	var ()
	return &OverrideParams{
		HTTPClient: client,
	}
}

/*OverrideParams contains all the parameters to send to the API endpoint
for the override operation typically these are written to a http.Request
*/
type OverrideParams struct {

	/*Status*/
	Status string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the override params
func (o *OverrideParams) WithTimeout(timeout time.Duration) *OverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the override params
func (o *OverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the override params
func (o *OverrideParams) WithContext(ctx context.Context) *OverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the override params
func (o *OverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the override params
func (o *OverrideParams) WithHTTPClient(client *http.Client) *OverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the override params
func (o *OverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStatus adds the status to the override params
func (o *OverrideParams) WithStatus(status string) *OverrideParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the override params
func (o *OverrideParams) SetStatus(status string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *OverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param status
	if err := r.SetPathParam("status", o.Status); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

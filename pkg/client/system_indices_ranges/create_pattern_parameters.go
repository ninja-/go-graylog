// Code generated by go-swagger; DO NOT EDIT.

package system_indices_ranges

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

// NewCreatePatternParams creates a new CreatePatternParams object
// with the default values initialized.
func NewCreatePatternParams() *CreatePatternParams {
	var ()
	return &CreatePatternParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePatternParamsWithTimeout creates a new CreatePatternParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePatternParamsWithTimeout(timeout time.Duration) *CreatePatternParams {
	var ()
	return &CreatePatternParams{

		timeout: timeout,
	}
}

// NewCreatePatternParamsWithContext creates a new CreatePatternParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePatternParamsWithContext(ctx context.Context) *CreatePatternParams {
	var ()
	return &CreatePatternParams{

		Context: ctx,
	}
}

// NewCreatePatternParamsWithHTTPClient creates a new CreatePatternParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePatternParamsWithHTTPClient(client *http.Client) *CreatePatternParams {
	var ()
	return &CreatePatternParams{
		HTTPClient: client,
	}
}

/*CreatePatternParams contains all the parameters to send to the API endpoint
for the create pattern operation typically these are written to a http.Request
*/
type CreatePatternParams struct {

	/*Pattern*/
	Pattern *models.GrokPattern

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create pattern params
func (o *CreatePatternParams) WithTimeout(timeout time.Duration) *CreatePatternParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pattern params
func (o *CreatePatternParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pattern params
func (o *CreatePatternParams) WithContext(ctx context.Context) *CreatePatternParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pattern params
func (o *CreatePatternParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pattern params
func (o *CreatePatternParams) WithHTTPClient(client *http.Client) *CreatePatternParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pattern params
func (o *CreatePatternParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPattern adds the pattern to the create pattern params
func (o *CreatePatternParams) WithPattern(pattern *models.GrokPattern) *CreatePatternParams {
	o.SetPattern(pattern)
	return o
}

// SetPattern adds the pattern to the create pattern params
func (o *CreatePatternParams) SetPattern(pattern *models.GrokPattern) {
	o.Pattern = pattern
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePatternParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Pattern != nil {
		if err := r.SetBodyParam(o.Pattern); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

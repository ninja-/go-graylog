// Code generated by go-swagger; DO NOT EDIT.

package system_journal

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

// NewTriggerParams creates a new TriggerParams object
// with the default values initialized.
func NewTriggerParams() *TriggerParams {
	var ()
	return &TriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerParamsWithTimeout creates a new TriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggerParamsWithTimeout(timeout time.Duration) *TriggerParams {
	var ()
	return &TriggerParams{

		timeout: timeout,
	}
}

// NewTriggerParamsWithContext creates a new TriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggerParamsWithContext(ctx context.Context) *TriggerParams {
	var ()
	return &TriggerParams{

		Context: ctx,
	}
}

// NewTriggerParamsWithHTTPClient creates a new TriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggerParamsWithHTTPClient(client *http.Client) *TriggerParams {
	var ()
	return &TriggerParams{
		HTTPClient: client,
	}
}

/*TriggerParams contains all the parameters to send to the API endpoint
for the trigger operation typically these are written to a http.Request
*/
type TriggerParams struct {

	/*JSONBody*/
	JSONBody *models.TriggerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trigger params
func (o *TriggerParams) WithTimeout(timeout time.Duration) *TriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger params
func (o *TriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger params
func (o *TriggerParams) WithContext(ctx context.Context) *TriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger params
func (o *TriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger params
func (o *TriggerParams) WithHTTPClient(client *http.Client) *TriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger params
func (o *TriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the trigger params
func (o *TriggerParams) WithJSONBody(jSONBody *models.TriggerRequest) *TriggerParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the trigger params
func (o *TriggerParams) SetJSONBody(jSONBody *models.TriggerRequest) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

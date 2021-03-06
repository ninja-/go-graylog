// Code generated by go-swagger; DO NOT EDIT.

package system_shutdown

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

// NewNewSessionParams creates a new NewSessionParams object
// with the default values initialized.
func NewNewSessionParams() *NewSessionParams {
	var ()
	return &NewSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNewSessionParamsWithTimeout creates a new NewSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNewSessionParamsWithTimeout(timeout time.Duration) *NewSessionParams {
	var ()
	return &NewSessionParams{

		timeout: timeout,
	}
}

// NewNewSessionParamsWithContext creates a new NewSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewNewSessionParamsWithContext(ctx context.Context) *NewSessionParams {
	var ()
	return &NewSessionParams{

		Context: ctx,
	}
}

// NewNewSessionParamsWithHTTPClient creates a new NewSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNewSessionParamsWithHTTPClient(client *http.Client) *NewSessionParams {
	var ()
	return &NewSessionParams{
		HTTPClient: client,
	}
}

/*NewSessionParams contains all the parameters to send to the API endpoint
for the new session operation typically these are written to a http.Request
*/
type NewSessionParams struct {

	/*LoginRequest
	  Username and credentials

	*/
	LoginRequest *models.SessionCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the new session params
func (o *NewSessionParams) WithTimeout(timeout time.Duration) *NewSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the new session params
func (o *NewSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the new session params
func (o *NewSessionParams) WithContext(ctx context.Context) *NewSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the new session params
func (o *NewSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the new session params
func (o *NewSessionParams) WithHTTPClient(client *http.Client) *NewSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the new session params
func (o *NewSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoginRequest adds the loginRequest to the new session params
func (o *NewSessionParams) WithLoginRequest(loginRequest *models.SessionCreateRequest) *NewSessionParams {
	o.SetLoginRequest(loginRequest)
	return o
}

// SetLoginRequest adds the loginRequest to the new session params
func (o *NewSessionParams) SetLoginRequest(loginRequest *models.SessionCreateRequest) {
	o.LoginRequest = loginRequest
}

// WriteToRequest writes these params to a swagger request
func (o *NewSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LoginRequest != nil {
		if err := r.SetBodyParam(o.LoginRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

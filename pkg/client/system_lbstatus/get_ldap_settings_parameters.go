// Code generated by go-swagger; DO NOT EDIT.

package system_lbstatus

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

// NewGetLdapSettingsParams creates a new GetLdapSettingsParams object
// with the default values initialized.
func NewGetLdapSettingsParams() *GetLdapSettingsParams {

	return &GetLdapSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapSettingsParamsWithTimeout creates a new GetLdapSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLdapSettingsParamsWithTimeout(timeout time.Duration) *GetLdapSettingsParams {

	return &GetLdapSettingsParams{

		timeout: timeout,
	}
}

// NewGetLdapSettingsParamsWithContext creates a new GetLdapSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLdapSettingsParamsWithContext(ctx context.Context) *GetLdapSettingsParams {

	return &GetLdapSettingsParams{

		Context: ctx,
	}
}

// NewGetLdapSettingsParamsWithHTTPClient creates a new GetLdapSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLdapSettingsParamsWithHTTPClient(client *http.Client) *GetLdapSettingsParams {

	return &GetLdapSettingsParams{
		HTTPClient: client,
	}
}

/*GetLdapSettingsParams contains all the parameters to send to the API endpoint
for the get ldap settings operation typically these are written to a http.Request
*/
type GetLdapSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ldap settings params
func (o *GetLdapSettingsParams) WithTimeout(timeout time.Duration) *GetLdapSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap settings params
func (o *GetLdapSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap settings params
func (o *GetLdapSettingsParams) WithContext(ctx context.Context) *GetLdapSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap settings params
func (o *GetLdapSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap settings params
func (o *GetLdapSettingsParams) WithHTTPClient(client *http.Client) *GetLdapSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap settings params
func (o *GetLdapSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

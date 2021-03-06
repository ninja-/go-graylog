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

// NewDeleteLdapSettingsParams creates a new DeleteLdapSettingsParams object
// with the default values initialized.
func NewDeleteLdapSettingsParams() *DeleteLdapSettingsParams {

	return &DeleteLdapSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLdapSettingsParamsWithTimeout creates a new DeleteLdapSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLdapSettingsParamsWithTimeout(timeout time.Duration) *DeleteLdapSettingsParams {

	return &DeleteLdapSettingsParams{

		timeout: timeout,
	}
}

// NewDeleteLdapSettingsParamsWithContext creates a new DeleteLdapSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLdapSettingsParamsWithContext(ctx context.Context) *DeleteLdapSettingsParams {

	return &DeleteLdapSettingsParams{

		Context: ctx,
	}
}

// NewDeleteLdapSettingsParamsWithHTTPClient creates a new DeleteLdapSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLdapSettingsParamsWithHTTPClient(client *http.Client) *DeleteLdapSettingsParams {

	return &DeleteLdapSettingsParams{
		HTTPClient: client,
	}
}

/*DeleteLdapSettingsParams contains all the parameters to send to the API endpoint
for the delete ldap settings operation typically these are written to a http.Request
*/
type DeleteLdapSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ldap settings params
func (o *DeleteLdapSettingsParams) WithTimeout(timeout time.Duration) *DeleteLdapSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ldap settings params
func (o *DeleteLdapSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ldap settings params
func (o *DeleteLdapSettingsParams) WithContext(ctx context.Context) *DeleteLdapSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ldap settings params
func (o *DeleteLdapSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ldap settings params
func (o *DeleteLdapSettingsParams) WithHTTPClient(client *http.Client) *DeleteLdapSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ldap settings params
func (o *DeleteLdapSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLdapSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

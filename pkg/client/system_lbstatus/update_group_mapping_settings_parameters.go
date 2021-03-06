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

// NewUpdateGroupMappingSettingsParams creates a new UpdateGroupMappingSettingsParams object
// with the default values initialized.
func NewUpdateGroupMappingSettingsParams() *UpdateGroupMappingSettingsParams {
	var ()
	return &UpdateGroupMappingSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGroupMappingSettingsParamsWithTimeout creates a new UpdateGroupMappingSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGroupMappingSettingsParamsWithTimeout(timeout time.Duration) *UpdateGroupMappingSettingsParams {
	var ()
	return &UpdateGroupMappingSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateGroupMappingSettingsParamsWithContext creates a new UpdateGroupMappingSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGroupMappingSettingsParamsWithContext(ctx context.Context) *UpdateGroupMappingSettingsParams {
	var ()
	return &UpdateGroupMappingSettingsParams{

		Context: ctx,
	}
}

// NewUpdateGroupMappingSettingsParamsWithHTTPClient creates a new UpdateGroupMappingSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGroupMappingSettingsParamsWithHTTPClient(client *http.Client) *UpdateGroupMappingSettingsParams {
	var ()
	return &UpdateGroupMappingSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateGroupMappingSettingsParams contains all the parameters to send to the API endpoint
for the update group mapping settings operation typically these are written to a http.Request
*/
type UpdateGroupMappingSettingsParams struct {

	/*JSONBody
	  A hash in which the keys are the LDAP group names and values is the Graylog role name.

	*/
	JSONBody interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) WithTimeout(timeout time.Duration) *UpdateGroupMappingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) WithContext(ctx context.Context) *UpdateGroupMappingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) WithHTTPClient(client *http.Client) *UpdateGroupMappingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) WithJSONBody(jSONBody interface{}) *UpdateGroupMappingSettingsParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the update group mapping settings params
func (o *UpdateGroupMappingSettingsParams) SetJSONBody(jSONBody interface{}) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGroupMappingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

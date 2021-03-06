// Code generated by go-swagger; DO NOT EDIT.

package system_lookup

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

// NewSetSubsystemLoggerLevelParams creates a new SetSubsystemLoggerLevelParams object
// with the default values initialized.
func NewSetSubsystemLoggerLevelParams() *SetSubsystemLoggerLevelParams {
	var ()
	return &SetSubsystemLoggerLevelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetSubsystemLoggerLevelParamsWithTimeout creates a new SetSubsystemLoggerLevelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetSubsystemLoggerLevelParamsWithTimeout(timeout time.Duration) *SetSubsystemLoggerLevelParams {
	var ()
	return &SetSubsystemLoggerLevelParams{

		timeout: timeout,
	}
}

// NewSetSubsystemLoggerLevelParamsWithContext creates a new SetSubsystemLoggerLevelParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetSubsystemLoggerLevelParamsWithContext(ctx context.Context) *SetSubsystemLoggerLevelParams {
	var ()
	return &SetSubsystemLoggerLevelParams{

		Context: ctx,
	}
}

// NewSetSubsystemLoggerLevelParamsWithHTTPClient creates a new SetSubsystemLoggerLevelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetSubsystemLoggerLevelParamsWithHTTPClient(client *http.Client) *SetSubsystemLoggerLevelParams {
	var ()
	return &SetSubsystemLoggerLevelParams{
		HTTPClient: client,
	}
}

/*SetSubsystemLoggerLevelParams contains all the parameters to send to the API endpoint
for the set subsystem logger level operation typically these are written to a http.Request
*/
type SetSubsystemLoggerLevelParams struct {

	/*Level*/
	Level string
	/*Subsystem*/
	Subsystem string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) WithTimeout(timeout time.Duration) *SetSubsystemLoggerLevelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) WithContext(ctx context.Context) *SetSubsystemLoggerLevelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) WithHTTPClient(client *http.Client) *SetSubsystemLoggerLevelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLevel adds the level to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) WithLevel(level string) *SetSubsystemLoggerLevelParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) SetLevel(level string) {
	o.Level = level
}

// WithSubsystem adds the subsystem to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) WithSubsystem(subsystem string) *SetSubsystemLoggerLevelParams {
	o.SetSubsystem(subsystem)
	return o
}

// SetSubsystem adds the subsystem to the set subsystem logger level params
func (o *SetSubsystemLoggerLevelParams) SetSubsystem(subsystem string) {
	o.Subsystem = subsystem
}

// WriteToRequest writes these params to a swagger request
func (o *SetSubsystemLoggerLevelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param level
	if err := r.SetPathParam("level", o.Level); err != nil {
		return err
	}

	// path param subsystem
	if err := r.SetPathParam("subsystem", o.Subsystem); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

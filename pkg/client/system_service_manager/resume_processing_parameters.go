// Code generated by go-swagger; DO NOT EDIT.

package system_service_manager

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

// NewResumeProcessingParams creates a new ResumeProcessingParams object
// with the default values initialized.
func NewResumeProcessingParams() *ResumeProcessingParams {

	return &ResumeProcessingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResumeProcessingParamsWithTimeout creates a new ResumeProcessingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResumeProcessingParamsWithTimeout(timeout time.Duration) *ResumeProcessingParams {

	return &ResumeProcessingParams{

		timeout: timeout,
	}
}

// NewResumeProcessingParamsWithContext creates a new ResumeProcessingParams object
// with the default values initialized, and the ability to set a context for a request
func NewResumeProcessingParamsWithContext(ctx context.Context) *ResumeProcessingParams {

	return &ResumeProcessingParams{

		Context: ctx,
	}
}

// NewResumeProcessingParamsWithHTTPClient creates a new ResumeProcessingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResumeProcessingParamsWithHTTPClient(client *http.Client) *ResumeProcessingParams {

	return &ResumeProcessingParams{
		HTTPClient: client,
	}
}

/*ResumeProcessingParams contains all the parameters to send to the API endpoint
for the resume processing operation typically these are written to a http.Request
*/
type ResumeProcessingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resume processing params
func (o *ResumeProcessingParams) WithTimeout(timeout time.Duration) *ResumeProcessingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume processing params
func (o *ResumeProcessingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume processing params
func (o *ResumeProcessingParams) WithContext(ctx context.Context) *ResumeProcessingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume processing params
func (o *ResumeProcessingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume processing params
func (o *ResumeProcessingParams) WithHTTPClient(client *http.Client) *ResumeProcessingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume processing params
func (o *ResumeProcessingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeProcessingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

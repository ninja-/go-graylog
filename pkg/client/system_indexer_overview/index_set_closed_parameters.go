// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_overview

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

// NewIndexSetClosedParams creates a new IndexSetClosedParams object
// with the default values initialized.
func NewIndexSetClosedParams() *IndexSetClosedParams {
	var ()
	return &IndexSetClosedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexSetClosedParamsWithTimeout creates a new IndexSetClosedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexSetClosedParamsWithTimeout(timeout time.Duration) *IndexSetClosedParams {
	var ()
	return &IndexSetClosedParams{

		timeout: timeout,
	}
}

// NewIndexSetClosedParamsWithContext creates a new IndexSetClosedParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexSetClosedParamsWithContext(ctx context.Context) *IndexSetClosedParams {
	var ()
	return &IndexSetClosedParams{

		Context: ctx,
	}
}

// NewIndexSetClosedParamsWithHTTPClient creates a new IndexSetClosedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexSetClosedParamsWithHTTPClient(client *http.Client) *IndexSetClosedParams {
	var ()
	return &IndexSetClosedParams{
		HTTPClient: client,
	}
}

/*IndexSetClosedParams contains all the parameters to send to the API endpoint
for the index set closed operation typically these are written to a http.Request
*/
type IndexSetClosedParams struct {

	/*IndexSetID*/
	IndexSetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index set closed params
func (o *IndexSetClosedParams) WithTimeout(timeout time.Duration) *IndexSetClosedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index set closed params
func (o *IndexSetClosedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index set closed params
func (o *IndexSetClosedParams) WithContext(ctx context.Context) *IndexSetClosedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index set closed params
func (o *IndexSetClosedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index set closed params
func (o *IndexSetClosedParams) WithHTTPClient(client *http.Client) *IndexSetClosedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index set closed params
func (o *IndexSetClosedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexSetID adds the indexSetID to the index set closed params
func (o *IndexSetClosedParams) WithIndexSetID(indexSetID string) *IndexSetClosedParams {
	o.SetIndexSetID(indexSetID)
	return o
}

// SetIndexSetID adds the indexSetId to the index set closed params
func (o *IndexSetClosedParams) SetIndexSetID(indexSetID string) {
	o.IndexSetID = indexSetID
}

// WriteToRequest writes these params to a swagger request
func (o *IndexSetClosedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param indexSetId
	if err := r.SetPathParam("indexSetId", o.IndexSetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

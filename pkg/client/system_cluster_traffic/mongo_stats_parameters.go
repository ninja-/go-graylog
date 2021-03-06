// Code generated by go-swagger; DO NOT EDIT.

package system_cluster_traffic

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

// NewMongoStatsParams creates a new MongoStatsParams object
// with the default values initialized.
func NewMongoStatsParams() *MongoStatsParams {

	return &MongoStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMongoStatsParamsWithTimeout creates a new MongoStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMongoStatsParamsWithTimeout(timeout time.Duration) *MongoStatsParams {

	return &MongoStatsParams{

		timeout: timeout,
	}
}

// NewMongoStatsParamsWithContext creates a new MongoStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewMongoStatsParamsWithContext(ctx context.Context) *MongoStatsParams {

	return &MongoStatsParams{

		Context: ctx,
	}
}

// NewMongoStatsParamsWithHTTPClient creates a new MongoStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMongoStatsParamsWithHTTPClient(client *http.Client) *MongoStatsParams {

	return &MongoStatsParams{
		HTTPClient: client,
	}
}

/*MongoStatsParams contains all the parameters to send to the API endpoint
for the mongo stats operation typically these are written to a http.Request
*/
type MongoStatsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the mongo stats params
func (o *MongoStatsParams) WithTimeout(timeout time.Duration) *MongoStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mongo stats params
func (o *MongoStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mongo stats params
func (o *MongoStatsParams) WithContext(ctx context.Context) *MongoStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mongo stats params
func (o *MongoStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mongo stats params
func (o *MongoStatsParams) WithHTTPClient(client *http.Client) *MongoStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mongo stats params
func (o *MongoStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MongoStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

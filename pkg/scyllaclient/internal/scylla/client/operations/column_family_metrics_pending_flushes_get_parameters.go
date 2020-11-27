// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsPendingFlushesGetParams creates a new ColumnFamilyMetricsPendingFlushesGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsPendingFlushesGetParams() *ColumnFamilyMetricsPendingFlushesGetParams {

	return &ColumnFamilyMetricsPendingFlushesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsPendingFlushesGetParamsWithTimeout creates a new ColumnFamilyMetricsPendingFlushesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsPendingFlushesGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsPendingFlushesGetParams {

	return &ColumnFamilyMetricsPendingFlushesGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsPendingFlushesGetParamsWithContext creates a new ColumnFamilyMetricsPendingFlushesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsPendingFlushesGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsPendingFlushesGetParams {

	return &ColumnFamilyMetricsPendingFlushesGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsPendingFlushesGetParamsWithHTTPClient creates a new ColumnFamilyMetricsPendingFlushesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsPendingFlushesGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsPendingFlushesGetParams {

	return &ColumnFamilyMetricsPendingFlushesGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsPendingFlushesGetParams contains all the parameters to send to the API endpoint
for the column family metrics pending flushes get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsPendingFlushesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics pending flushes get params
func (o *ColumnFamilyMetricsPendingFlushesGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsPendingFlushesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics pending flushes get params
func (o *ColumnFamilyMetricsPendingFlushesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics pending flushes get params
func (o *ColumnFamilyMetricsPendingFlushesGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsPendingFlushesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics pending flushes get params
func (o *ColumnFamilyMetricsPendingFlushesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics pending flushes get params
func (o *ColumnFamilyMetricsPendingFlushesGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsPendingFlushesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics pending flushes get params
func (o *ColumnFamilyMetricsPendingFlushesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsPendingFlushesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

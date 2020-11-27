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

// NewColumnFamilyGetParams creates a new ColumnFamilyGetParams object
// with the default values initialized.
func NewColumnFamilyGetParams() *ColumnFamilyGetParams {

	return &ColumnFamilyGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyGetParamsWithTimeout creates a new ColumnFamilyGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyGetParams {

	return &ColumnFamilyGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyGetParamsWithContext creates a new ColumnFamilyGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyGetParamsWithContext(ctx context.Context) *ColumnFamilyGetParams {

	return &ColumnFamilyGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyGetParamsWithHTTPClient creates a new ColumnFamilyGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyGetParams {

	return &ColumnFamilyGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyGetParams contains all the parameters to send to the API endpoint
for the column family get operation typically these are written to a http.Request
*/
type ColumnFamilyGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family get params
func (o *ColumnFamilyGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family get params
func (o *ColumnFamilyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family get params
func (o *ColumnFamilyGetParams) WithContext(ctx context.Context) *ColumnFamilyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family get params
func (o *ColumnFamilyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family get params
func (o *ColumnFamilyGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family get params
func (o *ColumnFamilyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

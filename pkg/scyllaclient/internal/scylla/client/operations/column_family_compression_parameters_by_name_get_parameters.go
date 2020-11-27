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

// NewColumnFamilyCompressionParametersByNameGetParams creates a new ColumnFamilyCompressionParametersByNameGetParams object
// with the default values initialized.
func NewColumnFamilyCompressionParametersByNameGetParams() *ColumnFamilyCompressionParametersByNameGetParams {
	var ()
	return &ColumnFamilyCompressionParametersByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyCompressionParametersByNameGetParamsWithTimeout creates a new ColumnFamilyCompressionParametersByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyCompressionParametersByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyCompressionParametersByNameGetParams {
	var ()
	return &ColumnFamilyCompressionParametersByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyCompressionParametersByNameGetParamsWithContext creates a new ColumnFamilyCompressionParametersByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyCompressionParametersByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyCompressionParametersByNameGetParams {
	var ()
	return &ColumnFamilyCompressionParametersByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyCompressionParametersByNameGetParamsWithHTTPClient creates a new ColumnFamilyCompressionParametersByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyCompressionParametersByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyCompressionParametersByNameGetParams {
	var ()
	return &ColumnFamilyCompressionParametersByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyCompressionParametersByNameGetParams contains all the parameters to send to the API endpoint
for the column family compression parameters by name get operation typically these are written to a http.Request
*/
type ColumnFamilyCompressionParametersByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyCompressionParametersByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyCompressionParametersByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyCompressionParametersByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) WithName(name string) *ColumnFamilyCompressionParametersByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family compression parameters by name get params
func (o *ColumnFamilyCompressionParametersByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyCompressionParametersByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

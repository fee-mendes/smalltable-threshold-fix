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
	"github.com/go-openapi/swag"
)

// NewColumnFamilyMinimumCompactionByNamePostParams creates a new ColumnFamilyMinimumCompactionByNamePostParams object
// with the default values initialized.
func NewColumnFamilyMinimumCompactionByNamePostParams() *ColumnFamilyMinimumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMinimumCompactionByNamePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMinimumCompactionByNamePostParamsWithTimeout creates a new ColumnFamilyMinimumCompactionByNamePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMinimumCompactionByNamePostParamsWithTimeout(timeout time.Duration) *ColumnFamilyMinimumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMinimumCompactionByNamePostParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMinimumCompactionByNamePostParamsWithContext creates a new ColumnFamilyMinimumCompactionByNamePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMinimumCompactionByNamePostParamsWithContext(ctx context.Context) *ColumnFamilyMinimumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMinimumCompactionByNamePostParams{

		Context: ctx,
	}
}

// NewColumnFamilyMinimumCompactionByNamePostParamsWithHTTPClient creates a new ColumnFamilyMinimumCompactionByNamePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMinimumCompactionByNamePostParamsWithHTTPClient(client *http.Client) *ColumnFamilyMinimumCompactionByNamePostParams {
	var ()
	return &ColumnFamilyMinimumCompactionByNamePostParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMinimumCompactionByNamePostParams contains all the parameters to send to the API endpoint
for the column family minimum compaction by name post operation typically these are written to a http.Request
*/
type ColumnFamilyMinimumCompactionByNamePostParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string
	/*Value
	  The minimum number of sstables in queue before compaction kicks off

	*/
	Value int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) WithTimeout(timeout time.Duration) *ColumnFamilyMinimumCompactionByNamePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) WithContext(ctx context.Context) *ColumnFamilyMinimumCompactionByNamePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) WithHTTPClient(client *http.Client) *ColumnFamilyMinimumCompactionByNamePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) WithName(name string) *ColumnFamilyMinimumCompactionByNamePostParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) SetName(name string) {
	o.Name = name
}

// WithValue adds the value to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) WithValue(value int32) *ColumnFamilyMinimumCompactionByNamePostParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the column family minimum compaction by name post params
func (o *ColumnFamilyMinimumCompactionByNamePostParams) SetValue(value int32) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMinimumCompactionByNamePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param value
	qrValue := o.Value
	qValue := swag.FormatInt32(qrValue)
	if qValue != "" {
		if err := r.SetQueryParam("value", qValue); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

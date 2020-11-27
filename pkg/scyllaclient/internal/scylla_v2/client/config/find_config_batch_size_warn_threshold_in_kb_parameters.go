// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigBatchSizeWarnThresholdInKbParams creates a new FindConfigBatchSizeWarnThresholdInKbParams object
// with the default values initialized.
func NewFindConfigBatchSizeWarnThresholdInKbParams() *FindConfigBatchSizeWarnThresholdInKbParams {

	return &FindConfigBatchSizeWarnThresholdInKbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigBatchSizeWarnThresholdInKbParamsWithTimeout creates a new FindConfigBatchSizeWarnThresholdInKbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigBatchSizeWarnThresholdInKbParamsWithTimeout(timeout time.Duration) *FindConfigBatchSizeWarnThresholdInKbParams {

	return &FindConfigBatchSizeWarnThresholdInKbParams{

		timeout: timeout,
	}
}

// NewFindConfigBatchSizeWarnThresholdInKbParamsWithContext creates a new FindConfigBatchSizeWarnThresholdInKbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigBatchSizeWarnThresholdInKbParamsWithContext(ctx context.Context) *FindConfigBatchSizeWarnThresholdInKbParams {

	return &FindConfigBatchSizeWarnThresholdInKbParams{

		Context: ctx,
	}
}

// NewFindConfigBatchSizeWarnThresholdInKbParamsWithHTTPClient creates a new FindConfigBatchSizeWarnThresholdInKbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigBatchSizeWarnThresholdInKbParamsWithHTTPClient(client *http.Client) *FindConfigBatchSizeWarnThresholdInKbParams {

	return &FindConfigBatchSizeWarnThresholdInKbParams{
		HTTPClient: client,
	}
}

/*FindConfigBatchSizeWarnThresholdInKbParams contains all the parameters to send to the API endpoint
for the find config batch size warn threshold in kb operation typically these are written to a http.Request
*/
type FindConfigBatchSizeWarnThresholdInKbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config batch size warn threshold in kb params
func (o *FindConfigBatchSizeWarnThresholdInKbParams) WithTimeout(timeout time.Duration) *FindConfigBatchSizeWarnThresholdInKbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config batch size warn threshold in kb params
func (o *FindConfigBatchSizeWarnThresholdInKbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config batch size warn threshold in kb params
func (o *FindConfigBatchSizeWarnThresholdInKbParams) WithContext(ctx context.Context) *FindConfigBatchSizeWarnThresholdInKbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config batch size warn threshold in kb params
func (o *FindConfigBatchSizeWarnThresholdInKbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config batch size warn threshold in kb params
func (o *FindConfigBatchSizeWarnThresholdInKbParams) WithHTTPClient(client *http.Client) *FindConfigBatchSizeWarnThresholdInKbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config batch size warn threshold in kb params
func (o *FindConfigBatchSizeWarnThresholdInKbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigBatchSizeWarnThresholdInKbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

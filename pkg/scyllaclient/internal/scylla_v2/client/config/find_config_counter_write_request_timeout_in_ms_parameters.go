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

// NewFindConfigCounterWriteRequestTimeoutInMsParams creates a new FindConfigCounterWriteRequestTimeoutInMsParams object
// with the default values initialized.
func NewFindConfigCounterWriteRequestTimeoutInMsParams() *FindConfigCounterWriteRequestTimeoutInMsParams {

	return &FindConfigCounterWriteRequestTimeoutInMsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCounterWriteRequestTimeoutInMsParamsWithTimeout creates a new FindConfigCounterWriteRequestTimeoutInMsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCounterWriteRequestTimeoutInMsParamsWithTimeout(timeout time.Duration) *FindConfigCounterWriteRequestTimeoutInMsParams {

	return &FindConfigCounterWriteRequestTimeoutInMsParams{

		timeout: timeout,
	}
}

// NewFindConfigCounterWriteRequestTimeoutInMsParamsWithContext creates a new FindConfigCounterWriteRequestTimeoutInMsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCounterWriteRequestTimeoutInMsParamsWithContext(ctx context.Context) *FindConfigCounterWriteRequestTimeoutInMsParams {

	return &FindConfigCounterWriteRequestTimeoutInMsParams{

		Context: ctx,
	}
}

// NewFindConfigCounterWriteRequestTimeoutInMsParamsWithHTTPClient creates a new FindConfigCounterWriteRequestTimeoutInMsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCounterWriteRequestTimeoutInMsParamsWithHTTPClient(client *http.Client) *FindConfigCounterWriteRequestTimeoutInMsParams {

	return &FindConfigCounterWriteRequestTimeoutInMsParams{
		HTTPClient: client,
	}
}

/*FindConfigCounterWriteRequestTimeoutInMsParams contains all the parameters to send to the API endpoint
for the find config counter write request timeout in ms operation typically these are written to a http.Request
*/
type FindConfigCounterWriteRequestTimeoutInMsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config counter write request timeout in ms params
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) WithTimeout(timeout time.Duration) *FindConfigCounterWriteRequestTimeoutInMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config counter write request timeout in ms params
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config counter write request timeout in ms params
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) WithContext(ctx context.Context) *FindConfigCounterWriteRequestTimeoutInMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config counter write request timeout in ms params
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config counter write request timeout in ms params
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) WithHTTPClient(client *http.Client) *FindConfigCounterWriteRequestTimeoutInMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config counter write request timeout in ms params
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCounterWriteRequestTimeoutInMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

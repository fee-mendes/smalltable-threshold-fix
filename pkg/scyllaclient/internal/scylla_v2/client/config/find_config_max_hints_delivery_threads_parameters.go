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

// NewFindConfigMaxHintsDeliveryThreadsParams creates a new FindConfigMaxHintsDeliveryThreadsParams object
// with the default values initialized.
func NewFindConfigMaxHintsDeliveryThreadsParams() *FindConfigMaxHintsDeliveryThreadsParams {

	return &FindConfigMaxHintsDeliveryThreadsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMaxHintsDeliveryThreadsParamsWithTimeout creates a new FindConfigMaxHintsDeliveryThreadsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigMaxHintsDeliveryThreadsParamsWithTimeout(timeout time.Duration) *FindConfigMaxHintsDeliveryThreadsParams {

	return &FindConfigMaxHintsDeliveryThreadsParams{

		timeout: timeout,
	}
}

// NewFindConfigMaxHintsDeliveryThreadsParamsWithContext creates a new FindConfigMaxHintsDeliveryThreadsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigMaxHintsDeliveryThreadsParamsWithContext(ctx context.Context) *FindConfigMaxHintsDeliveryThreadsParams {

	return &FindConfigMaxHintsDeliveryThreadsParams{

		Context: ctx,
	}
}

// NewFindConfigMaxHintsDeliveryThreadsParamsWithHTTPClient creates a new FindConfigMaxHintsDeliveryThreadsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigMaxHintsDeliveryThreadsParamsWithHTTPClient(client *http.Client) *FindConfigMaxHintsDeliveryThreadsParams {

	return &FindConfigMaxHintsDeliveryThreadsParams{
		HTTPClient: client,
	}
}

/*FindConfigMaxHintsDeliveryThreadsParams contains all the parameters to send to the API endpoint
for the find config max hints delivery threads operation typically these are written to a http.Request
*/
type FindConfigMaxHintsDeliveryThreadsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config max hints delivery threads params
func (o *FindConfigMaxHintsDeliveryThreadsParams) WithTimeout(timeout time.Duration) *FindConfigMaxHintsDeliveryThreadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config max hints delivery threads params
func (o *FindConfigMaxHintsDeliveryThreadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config max hints delivery threads params
func (o *FindConfigMaxHintsDeliveryThreadsParams) WithContext(ctx context.Context) *FindConfigMaxHintsDeliveryThreadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config max hints delivery threads params
func (o *FindConfigMaxHintsDeliveryThreadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config max hints delivery threads params
func (o *FindConfigMaxHintsDeliveryThreadsParams) WithHTTPClient(client *http.Client) *FindConfigMaxHintsDeliveryThreadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config max hints delivery threads params
func (o *FindConfigMaxHintsDeliveryThreadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMaxHintsDeliveryThreadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewFindConfigRPCSendBuffSizeInBytesParams creates a new FindConfigRPCSendBuffSizeInBytesParams object
// with the default values initialized.
func NewFindConfigRPCSendBuffSizeInBytesParams() *FindConfigRPCSendBuffSizeInBytesParams {

	return &FindConfigRPCSendBuffSizeInBytesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigRPCSendBuffSizeInBytesParamsWithTimeout creates a new FindConfigRPCSendBuffSizeInBytesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigRPCSendBuffSizeInBytesParamsWithTimeout(timeout time.Duration) *FindConfigRPCSendBuffSizeInBytesParams {

	return &FindConfigRPCSendBuffSizeInBytesParams{

		timeout: timeout,
	}
}

// NewFindConfigRPCSendBuffSizeInBytesParamsWithContext creates a new FindConfigRPCSendBuffSizeInBytesParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigRPCSendBuffSizeInBytesParamsWithContext(ctx context.Context) *FindConfigRPCSendBuffSizeInBytesParams {

	return &FindConfigRPCSendBuffSizeInBytesParams{

		Context: ctx,
	}
}

// NewFindConfigRPCSendBuffSizeInBytesParamsWithHTTPClient creates a new FindConfigRPCSendBuffSizeInBytesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigRPCSendBuffSizeInBytesParamsWithHTTPClient(client *http.Client) *FindConfigRPCSendBuffSizeInBytesParams {

	return &FindConfigRPCSendBuffSizeInBytesParams{
		HTTPClient: client,
	}
}

/*FindConfigRPCSendBuffSizeInBytesParams contains all the parameters to send to the API endpoint
for the find config rpc send buff size in bytes operation typically these are written to a http.Request
*/
type FindConfigRPCSendBuffSizeInBytesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config rpc send buff size in bytes params
func (o *FindConfigRPCSendBuffSizeInBytesParams) WithTimeout(timeout time.Duration) *FindConfigRPCSendBuffSizeInBytesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config rpc send buff size in bytes params
func (o *FindConfigRPCSendBuffSizeInBytesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config rpc send buff size in bytes params
func (o *FindConfigRPCSendBuffSizeInBytesParams) WithContext(ctx context.Context) *FindConfigRPCSendBuffSizeInBytesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config rpc send buff size in bytes params
func (o *FindConfigRPCSendBuffSizeInBytesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config rpc send buff size in bytes params
func (o *FindConfigRPCSendBuffSizeInBytesParams) WithHTTPClient(client *http.Client) *FindConfigRPCSendBuffSizeInBytesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config rpc send buff size in bytes params
func (o *FindConfigRPCSendBuffSizeInBytesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigRPCSendBuffSizeInBytesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

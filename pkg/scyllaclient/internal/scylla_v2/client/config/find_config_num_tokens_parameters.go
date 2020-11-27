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

// NewFindConfigNumTokensParams creates a new FindConfigNumTokensParams object
// with the default values initialized.
func NewFindConfigNumTokensParams() *FindConfigNumTokensParams {

	return &FindConfigNumTokensParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigNumTokensParamsWithTimeout creates a new FindConfigNumTokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigNumTokensParamsWithTimeout(timeout time.Duration) *FindConfigNumTokensParams {

	return &FindConfigNumTokensParams{

		timeout: timeout,
	}
}

// NewFindConfigNumTokensParamsWithContext creates a new FindConfigNumTokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigNumTokensParamsWithContext(ctx context.Context) *FindConfigNumTokensParams {

	return &FindConfigNumTokensParams{

		Context: ctx,
	}
}

// NewFindConfigNumTokensParamsWithHTTPClient creates a new FindConfigNumTokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigNumTokensParamsWithHTTPClient(client *http.Client) *FindConfigNumTokensParams {

	return &FindConfigNumTokensParams{
		HTTPClient: client,
	}
}

/*FindConfigNumTokensParams contains all the parameters to send to the API endpoint
for the find config num tokens operation typically these are written to a http.Request
*/
type FindConfigNumTokensParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config num tokens params
func (o *FindConfigNumTokensParams) WithTimeout(timeout time.Duration) *FindConfigNumTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config num tokens params
func (o *FindConfigNumTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config num tokens params
func (o *FindConfigNumTokensParams) WithContext(ctx context.Context) *FindConfigNumTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config num tokens params
func (o *FindConfigNumTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config num tokens params
func (o *FindConfigNumTokensParams) WithHTTPClient(client *http.Client) *FindConfigNumTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config num tokens params
func (o *FindConfigNumTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigNumTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

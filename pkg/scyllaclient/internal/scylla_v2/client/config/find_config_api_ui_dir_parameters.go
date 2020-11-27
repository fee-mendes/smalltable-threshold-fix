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

// NewFindConfigAPIUIDirParams creates a new FindConfigAPIUIDirParams object
// with the default values initialized.
func NewFindConfigAPIUIDirParams() *FindConfigAPIUIDirParams {

	return &FindConfigAPIUIDirParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigAPIUIDirParamsWithTimeout creates a new FindConfigAPIUIDirParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigAPIUIDirParamsWithTimeout(timeout time.Duration) *FindConfigAPIUIDirParams {

	return &FindConfigAPIUIDirParams{

		timeout: timeout,
	}
}

// NewFindConfigAPIUIDirParamsWithContext creates a new FindConfigAPIUIDirParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigAPIUIDirParamsWithContext(ctx context.Context) *FindConfigAPIUIDirParams {

	return &FindConfigAPIUIDirParams{

		Context: ctx,
	}
}

// NewFindConfigAPIUIDirParamsWithHTTPClient creates a new FindConfigAPIUIDirParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigAPIUIDirParamsWithHTTPClient(client *http.Client) *FindConfigAPIUIDirParams {

	return &FindConfigAPIUIDirParams{
		HTTPClient: client,
	}
}

/*FindConfigAPIUIDirParams contains all the parameters to send to the API endpoint
for the find config api ui dir operation typically these are written to a http.Request
*/
type FindConfigAPIUIDirParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config api ui dir params
func (o *FindConfigAPIUIDirParams) WithTimeout(timeout time.Duration) *FindConfigAPIUIDirParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config api ui dir params
func (o *FindConfigAPIUIDirParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config api ui dir params
func (o *FindConfigAPIUIDirParams) WithContext(ctx context.Context) *FindConfigAPIUIDirParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config api ui dir params
func (o *FindConfigAPIUIDirParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config api ui dir params
func (o *FindConfigAPIUIDirParams) WithHTTPClient(client *http.Client) *FindConfigAPIUIDirParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config api ui dir params
func (o *FindConfigAPIUIDirParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigAPIUIDirParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

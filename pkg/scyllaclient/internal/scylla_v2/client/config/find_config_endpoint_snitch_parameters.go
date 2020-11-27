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

// NewFindConfigEndpointSnitchParams creates a new FindConfigEndpointSnitchParams object
// with the default values initialized.
func NewFindConfigEndpointSnitchParams() *FindConfigEndpointSnitchParams {

	return &FindConfigEndpointSnitchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigEndpointSnitchParamsWithTimeout creates a new FindConfigEndpointSnitchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigEndpointSnitchParamsWithTimeout(timeout time.Duration) *FindConfigEndpointSnitchParams {

	return &FindConfigEndpointSnitchParams{

		timeout: timeout,
	}
}

// NewFindConfigEndpointSnitchParamsWithContext creates a new FindConfigEndpointSnitchParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigEndpointSnitchParamsWithContext(ctx context.Context) *FindConfigEndpointSnitchParams {

	return &FindConfigEndpointSnitchParams{

		Context: ctx,
	}
}

// NewFindConfigEndpointSnitchParamsWithHTTPClient creates a new FindConfigEndpointSnitchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigEndpointSnitchParamsWithHTTPClient(client *http.Client) *FindConfigEndpointSnitchParams {

	return &FindConfigEndpointSnitchParams{
		HTTPClient: client,
	}
}

/*FindConfigEndpointSnitchParams contains all the parameters to send to the API endpoint
for the find config endpoint snitch operation typically these are written to a http.Request
*/
type FindConfigEndpointSnitchParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config endpoint snitch params
func (o *FindConfigEndpointSnitchParams) WithTimeout(timeout time.Duration) *FindConfigEndpointSnitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config endpoint snitch params
func (o *FindConfigEndpointSnitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config endpoint snitch params
func (o *FindConfigEndpointSnitchParams) WithContext(ctx context.Context) *FindConfigEndpointSnitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config endpoint snitch params
func (o *FindConfigEndpointSnitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config endpoint snitch params
func (o *FindConfigEndpointSnitchParams) WithHTTPClient(client *http.Client) *FindConfigEndpointSnitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config endpoint snitch params
func (o *FindConfigEndpointSnitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigEndpointSnitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

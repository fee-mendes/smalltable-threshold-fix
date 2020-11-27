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

// NewFindConfigAutoSnapshotParams creates a new FindConfigAutoSnapshotParams object
// with the default values initialized.
func NewFindConfigAutoSnapshotParams() *FindConfigAutoSnapshotParams {

	return &FindConfigAutoSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigAutoSnapshotParamsWithTimeout creates a new FindConfigAutoSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigAutoSnapshotParamsWithTimeout(timeout time.Duration) *FindConfigAutoSnapshotParams {

	return &FindConfigAutoSnapshotParams{

		timeout: timeout,
	}
}

// NewFindConfigAutoSnapshotParamsWithContext creates a new FindConfigAutoSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigAutoSnapshotParamsWithContext(ctx context.Context) *FindConfigAutoSnapshotParams {

	return &FindConfigAutoSnapshotParams{

		Context: ctx,
	}
}

// NewFindConfigAutoSnapshotParamsWithHTTPClient creates a new FindConfigAutoSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigAutoSnapshotParamsWithHTTPClient(client *http.Client) *FindConfigAutoSnapshotParams {

	return &FindConfigAutoSnapshotParams{
		HTTPClient: client,
	}
}

/*FindConfigAutoSnapshotParams contains all the parameters to send to the API endpoint
for the find config auto snapshot operation typically these are written to a http.Request
*/
type FindConfigAutoSnapshotParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config auto snapshot params
func (o *FindConfigAutoSnapshotParams) WithTimeout(timeout time.Duration) *FindConfigAutoSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config auto snapshot params
func (o *FindConfigAutoSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config auto snapshot params
func (o *FindConfigAutoSnapshotParams) WithContext(ctx context.Context) *FindConfigAutoSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config auto snapshot params
func (o *FindConfigAutoSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config auto snapshot params
func (o *FindConfigAutoSnapshotParams) WithHTTPClient(client *http.Client) *FindConfigAutoSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config auto snapshot params
func (o *FindConfigAutoSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigAutoSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

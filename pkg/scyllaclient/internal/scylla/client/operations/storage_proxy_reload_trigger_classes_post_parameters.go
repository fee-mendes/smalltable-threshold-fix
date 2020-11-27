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

// NewStorageProxyReloadTriggerClassesPostParams creates a new StorageProxyReloadTriggerClassesPostParams object
// with the default values initialized.
func NewStorageProxyReloadTriggerClassesPostParams() *StorageProxyReloadTriggerClassesPostParams {

	return &StorageProxyReloadTriggerClassesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyReloadTriggerClassesPostParamsWithTimeout creates a new StorageProxyReloadTriggerClassesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyReloadTriggerClassesPostParamsWithTimeout(timeout time.Duration) *StorageProxyReloadTriggerClassesPostParams {

	return &StorageProxyReloadTriggerClassesPostParams{

		timeout: timeout,
	}
}

// NewStorageProxyReloadTriggerClassesPostParamsWithContext creates a new StorageProxyReloadTriggerClassesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyReloadTriggerClassesPostParamsWithContext(ctx context.Context) *StorageProxyReloadTriggerClassesPostParams {

	return &StorageProxyReloadTriggerClassesPostParams{

		Context: ctx,
	}
}

// NewStorageProxyReloadTriggerClassesPostParamsWithHTTPClient creates a new StorageProxyReloadTriggerClassesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyReloadTriggerClassesPostParamsWithHTTPClient(client *http.Client) *StorageProxyReloadTriggerClassesPostParams {

	return &StorageProxyReloadTriggerClassesPostParams{
		HTTPClient: client,
	}
}

/*StorageProxyReloadTriggerClassesPostParams contains all the parameters to send to the API endpoint
for the storage proxy reload trigger classes post operation typically these are written to a http.Request
*/
type StorageProxyReloadTriggerClassesPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy reload trigger classes post params
func (o *StorageProxyReloadTriggerClassesPostParams) WithTimeout(timeout time.Duration) *StorageProxyReloadTriggerClassesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy reload trigger classes post params
func (o *StorageProxyReloadTriggerClassesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy reload trigger classes post params
func (o *StorageProxyReloadTriggerClassesPostParams) WithContext(ctx context.Context) *StorageProxyReloadTriggerClassesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy reload trigger classes post params
func (o *StorageProxyReloadTriggerClassesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy reload trigger classes post params
func (o *StorageProxyReloadTriggerClassesPostParams) WithHTTPClient(client *http.Client) *StorageProxyReloadTriggerClassesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy reload trigger classes post params
func (o *StorageProxyReloadTriggerClassesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyReloadTriggerClassesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

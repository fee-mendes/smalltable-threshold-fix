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

// NewCacheServiceMetricsKeySizeGetParams creates a new CacheServiceMetricsKeySizeGetParams object
// with the default values initialized.
func NewCacheServiceMetricsKeySizeGetParams() *CacheServiceMetricsKeySizeGetParams {

	return &CacheServiceMetricsKeySizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsKeySizeGetParamsWithTimeout creates a new CacheServiceMetricsKeySizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsKeySizeGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsKeySizeGetParams {

	return &CacheServiceMetricsKeySizeGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsKeySizeGetParamsWithContext creates a new CacheServiceMetricsKeySizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsKeySizeGetParamsWithContext(ctx context.Context) *CacheServiceMetricsKeySizeGetParams {

	return &CacheServiceMetricsKeySizeGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsKeySizeGetParamsWithHTTPClient creates a new CacheServiceMetricsKeySizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsKeySizeGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsKeySizeGetParams {

	return &CacheServiceMetricsKeySizeGetParams{
		HTTPClient: client,
	}
}

/*CacheServiceMetricsKeySizeGetParams contains all the parameters to send to the API endpoint
for the cache service metrics key size get operation typically these are written to a http.Request
*/
type CacheServiceMetricsKeySizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics key size get params
func (o *CacheServiceMetricsKeySizeGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsKeySizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics key size get params
func (o *CacheServiceMetricsKeySizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics key size get params
func (o *CacheServiceMetricsKeySizeGetParams) WithContext(ctx context.Context) *CacheServiceMetricsKeySizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics key size get params
func (o *CacheServiceMetricsKeySizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics key size get params
func (o *CacheServiceMetricsKeySizeGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsKeySizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics key size get params
func (o *CacheServiceMetricsKeySizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsKeySizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

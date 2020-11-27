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

// NewStreamManagerMetricsIncomingGetParams creates a new StreamManagerMetricsIncomingGetParams object
// with the default values initialized.
func NewStreamManagerMetricsIncomingGetParams() *StreamManagerMetricsIncomingGetParams {

	return &StreamManagerMetricsIncomingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStreamManagerMetricsIncomingGetParamsWithTimeout creates a new StreamManagerMetricsIncomingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStreamManagerMetricsIncomingGetParamsWithTimeout(timeout time.Duration) *StreamManagerMetricsIncomingGetParams {

	return &StreamManagerMetricsIncomingGetParams{

		timeout: timeout,
	}
}

// NewStreamManagerMetricsIncomingGetParamsWithContext creates a new StreamManagerMetricsIncomingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStreamManagerMetricsIncomingGetParamsWithContext(ctx context.Context) *StreamManagerMetricsIncomingGetParams {

	return &StreamManagerMetricsIncomingGetParams{

		Context: ctx,
	}
}

// NewStreamManagerMetricsIncomingGetParamsWithHTTPClient creates a new StreamManagerMetricsIncomingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStreamManagerMetricsIncomingGetParamsWithHTTPClient(client *http.Client) *StreamManagerMetricsIncomingGetParams {

	return &StreamManagerMetricsIncomingGetParams{
		HTTPClient: client,
	}
}

/*StreamManagerMetricsIncomingGetParams contains all the parameters to send to the API endpoint
for the stream manager metrics incoming get operation typically these are written to a http.Request
*/
type StreamManagerMetricsIncomingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stream manager metrics incoming get params
func (o *StreamManagerMetricsIncomingGetParams) WithTimeout(timeout time.Duration) *StreamManagerMetricsIncomingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stream manager metrics incoming get params
func (o *StreamManagerMetricsIncomingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stream manager metrics incoming get params
func (o *StreamManagerMetricsIncomingGetParams) WithContext(ctx context.Context) *StreamManagerMetricsIncomingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stream manager metrics incoming get params
func (o *StreamManagerMetricsIncomingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stream manager metrics incoming get params
func (o *StreamManagerMetricsIncomingGetParams) WithHTTPClient(client *http.Client) *StreamManagerMetricsIncomingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stream manager metrics incoming get params
func (o *StreamManagerMetricsIncomingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StreamManagerMetricsIncomingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

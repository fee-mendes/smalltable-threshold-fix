// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterClusterIDParams creates a new GetClusterClusterIDParams object
// with the default values initialized.
func NewGetClusterClusterIDParams() *GetClusterClusterIDParams {
	var ()
	return &GetClusterClusterIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDParamsWithTimeout creates a new GetClusterClusterIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDParams {
	var ()
	return &GetClusterClusterIDParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDParamsWithContext creates a new GetClusterClusterIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDParamsWithContext(ctx context.Context) *GetClusterClusterIDParams {
	var ()
	return &GetClusterClusterIDParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDParamsWithHTTPClient creates a new GetClusterClusterIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDParams {
	var ()
	return &GetClusterClusterIDParams{
		HTTPClient: client,
	}
}

/*GetClusterClusterIDParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID operation typically these are written to a http.Request
*/
type GetClusterClusterIDParams struct {

	/*ClusterID*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) WithContext(ctx context.Context) *GetClusterClusterIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) WithClusterID(clusterID string) *GetClusterClusterIDParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID params
func (o *GetClusterClusterIDParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

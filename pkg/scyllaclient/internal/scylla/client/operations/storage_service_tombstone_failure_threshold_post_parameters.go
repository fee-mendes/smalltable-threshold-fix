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
	"github.com/go-openapi/swag"
)

// NewStorageServiceTombstoneFailureThresholdPostParams creates a new StorageServiceTombstoneFailureThresholdPostParams object
// with the default values initialized.
func NewStorageServiceTombstoneFailureThresholdPostParams() *StorageServiceTombstoneFailureThresholdPostParams {
	var ()
	return &StorageServiceTombstoneFailureThresholdPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceTombstoneFailureThresholdPostParamsWithTimeout creates a new StorageServiceTombstoneFailureThresholdPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceTombstoneFailureThresholdPostParamsWithTimeout(timeout time.Duration) *StorageServiceTombstoneFailureThresholdPostParams {
	var ()
	return &StorageServiceTombstoneFailureThresholdPostParams{

		timeout: timeout,
	}
}

// NewStorageServiceTombstoneFailureThresholdPostParamsWithContext creates a new StorageServiceTombstoneFailureThresholdPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceTombstoneFailureThresholdPostParamsWithContext(ctx context.Context) *StorageServiceTombstoneFailureThresholdPostParams {
	var ()
	return &StorageServiceTombstoneFailureThresholdPostParams{

		Context: ctx,
	}
}

// NewStorageServiceTombstoneFailureThresholdPostParamsWithHTTPClient creates a new StorageServiceTombstoneFailureThresholdPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceTombstoneFailureThresholdPostParamsWithHTTPClient(client *http.Client) *StorageServiceTombstoneFailureThresholdPostParams {
	var ()
	return &StorageServiceTombstoneFailureThresholdPostParams{
		HTTPClient: client,
	}
}

/*StorageServiceTombstoneFailureThresholdPostParams contains all the parameters to send to the API endpoint
for the storage service tombstone failure threshold post operation typically these are written to a http.Request
*/
type StorageServiceTombstoneFailureThresholdPostParams struct {

	/*TombstoneDebugThreshold
	  tombstone debug threshold

	*/
	TombstoneDebugThreshold int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) WithTimeout(timeout time.Duration) *StorageServiceTombstoneFailureThresholdPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) WithContext(ctx context.Context) *StorageServiceTombstoneFailureThresholdPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) WithHTTPClient(client *http.Client) *StorageServiceTombstoneFailureThresholdPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTombstoneDebugThreshold adds the tombstoneDebugThreshold to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) WithTombstoneDebugThreshold(tombstoneDebugThreshold int32) *StorageServiceTombstoneFailureThresholdPostParams {
	o.SetTombstoneDebugThreshold(tombstoneDebugThreshold)
	return o
}

// SetTombstoneDebugThreshold adds the tombstoneDebugThreshold to the storage service tombstone failure threshold post params
func (o *StorageServiceTombstoneFailureThresholdPostParams) SetTombstoneDebugThreshold(tombstoneDebugThreshold int32) {
	o.TombstoneDebugThreshold = tombstoneDebugThreshold
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceTombstoneFailureThresholdPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param tombstone_debug_threshold
	qrTombstoneDebugThreshold := o.TombstoneDebugThreshold
	qTombstoneDebugThreshold := swag.FormatInt32(qrTombstoneDebugThreshold)
	if qTombstoneDebugThreshold != "" {
		if err := r.SetQueryParam("tombstone_debug_threshold", qTombstoneDebugThreshold); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

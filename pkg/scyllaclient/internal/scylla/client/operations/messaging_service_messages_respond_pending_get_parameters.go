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

// NewMessagingServiceMessagesRespondPendingGetParams creates a new MessagingServiceMessagesRespondPendingGetParams object
// with the default values initialized.
func NewMessagingServiceMessagesRespondPendingGetParams() *MessagingServiceMessagesRespondPendingGetParams {

	return &MessagingServiceMessagesRespondPendingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesRespondPendingGetParamsWithTimeout creates a new MessagingServiceMessagesRespondPendingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagingServiceMessagesRespondPendingGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesRespondPendingGetParams {

	return &MessagingServiceMessagesRespondPendingGetParams{

		timeout: timeout,
	}
}

// NewMessagingServiceMessagesRespondPendingGetParamsWithContext creates a new MessagingServiceMessagesRespondPendingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagingServiceMessagesRespondPendingGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesRespondPendingGetParams {

	return &MessagingServiceMessagesRespondPendingGetParams{

		Context: ctx,
	}
}

// NewMessagingServiceMessagesRespondPendingGetParamsWithHTTPClient creates a new MessagingServiceMessagesRespondPendingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagingServiceMessagesRespondPendingGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesRespondPendingGetParams {

	return &MessagingServiceMessagesRespondPendingGetParams{
		HTTPClient: client,
	}
}

/*MessagingServiceMessagesRespondPendingGetParams contains all the parameters to send to the API endpoint
for the messaging service messages respond pending get operation typically these are written to a http.Request
*/
type MessagingServiceMessagesRespondPendingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messaging service messages respond pending get params
func (o *MessagingServiceMessagesRespondPendingGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesRespondPendingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages respond pending get params
func (o *MessagingServiceMessagesRespondPendingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages respond pending get params
func (o *MessagingServiceMessagesRespondPendingGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesRespondPendingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages respond pending get params
func (o *MessagingServiceMessagesRespondPendingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages respond pending get params
func (o *MessagingServiceMessagesRespondPendingGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesRespondPendingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages respond pending get params
func (o *MessagingServiceMessagesRespondPendingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesRespondPendingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

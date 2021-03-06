// Code generated by go-swagger; DO NOT EDIT.

package webhook

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

// NewDeleteWebhooksWebhookIDParams creates a new DeleteWebhooksWebhookIDParams object
// with the default values initialized.
func NewDeleteWebhooksWebhookIDParams() *DeleteWebhooksWebhookIDParams {
	var ()
	return &DeleteWebhooksWebhookIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWebhooksWebhookIDParamsWithTimeout creates a new DeleteWebhooksWebhookIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWebhooksWebhookIDParamsWithTimeout(timeout time.Duration) *DeleteWebhooksWebhookIDParams {
	var ()
	return &DeleteWebhooksWebhookIDParams{

		timeout: timeout,
	}
}

// NewDeleteWebhooksWebhookIDParamsWithContext creates a new DeleteWebhooksWebhookIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWebhooksWebhookIDParamsWithContext(ctx context.Context) *DeleteWebhooksWebhookIDParams {
	var ()
	return &DeleteWebhooksWebhookIDParams{

		Context: ctx,
	}
}

// NewDeleteWebhooksWebhookIDParamsWithHTTPClient creates a new DeleteWebhooksWebhookIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWebhooksWebhookIDParamsWithHTTPClient(client *http.Client) *DeleteWebhooksWebhookIDParams {
	var ()
	return &DeleteWebhooksWebhookIDParams{
		HTTPClient: client,
	}
}

/*DeleteWebhooksWebhookIDParams contains all the parameters to send to the API endpoint
for the delete webhooks webhook ID operation typically these are written to a http.Request
*/
type DeleteWebhooksWebhookIDParams struct {

	/*WebhookID
	  Идентификатор хука

	*/
	WebhookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) WithTimeout(timeout time.Duration) *DeleteWebhooksWebhookIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) WithContext(ctx context.Context) *DeleteWebhooksWebhookIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) WithHTTPClient(client *http.Client) *DeleteWebhooksWebhookIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebhookID adds the webhookID to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) WithWebhookID(webhookID string) *DeleteWebhooksWebhookIDParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the delete webhooks webhook ID params
func (o *DeleteWebhooksWebhookIDParams) SetWebhookID(webhookID string) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWebhooksWebhookIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param webhookID
	if err := r.SetPathParam("webhookID", o.WebhookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

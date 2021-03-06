// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPatientsParams creates a new GetPatientsParams object
// with the default values initialized.
func NewGetPatientsParams() *GetPatientsParams {
	var ()
	return &GetPatientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatientsParamsWithTimeout creates a new GetPatientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatientsParamsWithTimeout(timeout time.Duration) *GetPatientsParams {
	var ()
	return &GetPatientsParams{

		timeout: timeout,
	}
}

// NewGetPatientsParamsWithContext creates a new GetPatientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatientsParamsWithContext(ctx context.Context) *GetPatientsParams {
	var ()
	return &GetPatientsParams{

		Context: ctx,
	}
}

// NewGetPatientsParamsWithHTTPClient creates a new GetPatientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatientsParamsWithHTTPClient(client *http.Client) *GetPatientsParams {
	var ()
	return &GetPatientsParams{
		HTTPClient: client,
	}
}

/*GetPatientsParams contains all the parameters to send to the API endpoint
for the get patients operation typically these are written to a http.Request
*/
type GetPatientsParams struct {

	/*Limit
	  Лимит выдачи

	*/
	Limit *int64
	/*Offset
	  Шаг выдачи

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patients params
func (o *GetPatientsParams) WithTimeout(timeout time.Duration) *GetPatientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patients params
func (o *GetPatientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patients params
func (o *GetPatientsParams) WithContext(ctx context.Context) *GetPatientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patients params
func (o *GetPatientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patients params
func (o *GetPatientsParams) WithHTTPClient(client *http.Client) *GetPatientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patients params
func (o *GetPatientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get patients params
func (o *GetPatientsParams) WithLimit(limit *int64) *GetPatientsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get patients params
func (o *GetPatientsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get patients params
func (o *GetPatientsParams) WithOffset(offset *int64) *GetPatientsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get patients params
func (o *GetPatientsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

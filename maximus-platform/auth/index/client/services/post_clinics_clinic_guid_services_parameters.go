// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewPostClinicsClinicGUIDServicesParams creates a new PostClinicsClinicGUIDServicesParams object
// with the default values initialized.
func NewPostClinicsClinicGUIDServicesParams() *PostClinicsClinicGUIDServicesParams {
	var ()
	return &PostClinicsClinicGUIDServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClinicsClinicGUIDServicesParamsWithTimeout creates a new PostClinicsClinicGUIDServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClinicsClinicGUIDServicesParamsWithTimeout(timeout time.Duration) *PostClinicsClinicGUIDServicesParams {
	var ()
	return &PostClinicsClinicGUIDServicesParams{

		timeout: timeout,
	}
}

// NewPostClinicsClinicGUIDServicesParamsWithContext creates a new PostClinicsClinicGUIDServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClinicsClinicGUIDServicesParamsWithContext(ctx context.Context) *PostClinicsClinicGUIDServicesParams {
	var ()
	return &PostClinicsClinicGUIDServicesParams{

		Context: ctx,
	}
}

// NewPostClinicsClinicGUIDServicesParamsWithHTTPClient creates a new PostClinicsClinicGUIDServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClinicsClinicGUIDServicesParamsWithHTTPClient(client *http.Client) *PostClinicsClinicGUIDServicesParams {
	var ()
	return &PostClinicsClinicGUIDServicesParams{
		HTTPClient: client,
	}
}

/*PostClinicsClinicGUIDServicesParams contains all the parameters to send to the API endpoint
for the post clinics clinic GUID services operation typically these are written to a http.Request
*/
type PostClinicsClinicGUIDServicesParams struct {

	/*Body*/
	Body PostClinicsClinicGUIDServicesBody
	/*ClinicGUID
	  GUID клиники

	*/
	ClinicGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) WithTimeout(timeout time.Duration) *PostClinicsClinicGUIDServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) WithContext(ctx context.Context) *PostClinicsClinicGUIDServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) WithHTTPClient(client *http.Client) *PostClinicsClinicGUIDServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) WithBody(body PostClinicsClinicGUIDServicesBody) *PostClinicsClinicGUIDServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) SetBody(body PostClinicsClinicGUIDServicesBody) {
	o.Body = body
}

// WithClinicGUID adds the clinicGUID to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) WithClinicGUID(clinicGUID string) *PostClinicsClinicGUIDServicesParams {
	o.SetClinicGUID(clinicGUID)
	return o
}

// SetClinicGUID adds the clinicGuid to the post clinics clinic GUID services params
func (o *PostClinicsClinicGUIDServicesParams) SetClinicGUID(clinicGUID string) {
	o.ClinicGUID = clinicGUID
}

// WriteToRequest writes these params to a swagger request
func (o *PostClinicsClinicGUIDServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clinicGUID
	if err := r.SetPathParam("clinicGUID", o.ClinicGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
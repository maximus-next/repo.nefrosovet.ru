// Code generated by go-swagger; DO NOT EDIT.

package clinic_employees

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

// NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams creates a new DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized.
func NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams() *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithTimeout creates a new DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithTimeout(timeout time.Duration) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams{

		timeout: timeout,
	}
}

// NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithContext creates a new DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithContext(ctx context.Context) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams{

		Context: ctx,
	}
}

// NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithHTTPClient creates a new DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithHTTPClient(client *http.Client) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams{
		HTTPClient: client,
	}
}

/*DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams contains all the parameters to send to the API endpoint
for the delete clinics clinic GUID employees employee GUID operation typically these are written to a http.Request
*/
type DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams struct {

	/*ClinicGUID
	  GUID клиники

	*/
	ClinicGUID string
	/*EmployeeGUID
	  GUID сотрудника

	*/
	EmployeeGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithTimeout(timeout time.Duration) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithContext(ctx context.Context) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithHTTPClient(client *http.Client) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClinicGUID adds the clinicGUID to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithClinicGUID(clinicGUID string) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetClinicGUID(clinicGUID)
	return o
}

// SetClinicGUID adds the clinicGuid to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetClinicGUID(clinicGUID string) {
	o.ClinicGUID = clinicGUID
}

// WithEmployeeGUID adds the employeeGUID to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithEmployeeGUID(employeeGUID string) *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetEmployeeGUID(employeeGUID)
	return o
}

// SetEmployeeGUID adds the employeeGuid to the delete clinics clinic GUID employees employee GUID params
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetEmployeeGUID(employeeGUID string) {
	o.EmployeeGUID = employeeGUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClinicsClinicGUIDEmployeesEmployeeGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clinicGUID
	if err := r.SetPathParam("clinicGUID", o.ClinicGUID); err != nil {
		return err
	}

	// path param employeeGUID
	if err := r.SetPathParam("employeeGUID", o.EmployeeGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

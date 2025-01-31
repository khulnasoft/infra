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

	"github.com/khulnasoft/infra/packages/shared/pkg/fc/models"
)

// NewPatchGuestDriveByIDParams creates a new PatchGuestDriveByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchGuestDriveByIDParams() *PatchGuestDriveByIDParams {
	return &PatchGuestDriveByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGuestDriveByIDParamsWithTimeout creates a new PatchGuestDriveByIDParams object
// with the ability to set a timeout on a request.
func NewPatchGuestDriveByIDParamsWithTimeout(timeout time.Duration) *PatchGuestDriveByIDParams {
	return &PatchGuestDriveByIDParams{
		timeout: timeout,
	}
}

// NewPatchGuestDriveByIDParamsWithContext creates a new PatchGuestDriveByIDParams object
// with the ability to set a context for a request.
func NewPatchGuestDriveByIDParamsWithContext(ctx context.Context) *PatchGuestDriveByIDParams {
	return &PatchGuestDriveByIDParams{
		Context: ctx,
	}
}

// NewPatchGuestDriveByIDParamsWithHTTPClient creates a new PatchGuestDriveByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchGuestDriveByIDParamsWithHTTPClient(client *http.Client) *PatchGuestDriveByIDParams {
	return &PatchGuestDriveByIDParams{
		HTTPClient: client,
	}
}

/*
PatchGuestDriveByIDParams contains all the parameters to send to the API endpoint

	for the patch guest drive by ID operation.

	Typically these are written to a http.Request.
*/
type PatchGuestDriveByIDParams struct {

	/* Body.

	   Guest drive properties
	*/
	Body *models.PartialDrive

	/* DriveID.

	   The id of the guest drive
	*/
	DriveID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch guest drive by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGuestDriveByIDParams) WithDefaults() *PatchGuestDriveByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch guest drive by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGuestDriveByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) WithTimeout(timeout time.Duration) *PatchGuestDriveByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) WithContext(ctx context.Context) *PatchGuestDriveByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) WithHTTPClient(client *http.Client) *PatchGuestDriveByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) WithBody(body *models.PartialDrive) *PatchGuestDriveByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) SetBody(body *models.PartialDrive) {
	o.Body = body
}

// WithDriveID adds the driveID to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) WithDriveID(driveID string) *PatchGuestDriveByIDParams {
	o.SetDriveID(driveID)
	return o
}

// SetDriveID adds the driveId to the patch guest drive by ID params
func (o *PatchGuestDriveByIDParams) SetDriveID(driveID string) {
	o.DriveID = driveID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGuestDriveByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param drive_id
	if err := r.SetPathParam("drive_id", o.DriveID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

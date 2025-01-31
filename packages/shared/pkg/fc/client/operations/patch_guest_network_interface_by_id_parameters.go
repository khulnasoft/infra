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

// NewPatchGuestNetworkInterfaceByIDParams creates a new PatchGuestNetworkInterfaceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchGuestNetworkInterfaceByIDParams() *PatchGuestNetworkInterfaceByIDParams {
	return &PatchGuestNetworkInterfaceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGuestNetworkInterfaceByIDParamsWithTimeout creates a new PatchGuestNetworkInterfaceByIDParams object
// with the ability to set a timeout on a request.
func NewPatchGuestNetworkInterfaceByIDParamsWithTimeout(timeout time.Duration) *PatchGuestNetworkInterfaceByIDParams {
	return &PatchGuestNetworkInterfaceByIDParams{
		timeout: timeout,
	}
}

// NewPatchGuestNetworkInterfaceByIDParamsWithContext creates a new PatchGuestNetworkInterfaceByIDParams object
// with the ability to set a context for a request.
func NewPatchGuestNetworkInterfaceByIDParamsWithContext(ctx context.Context) *PatchGuestNetworkInterfaceByIDParams {
	return &PatchGuestNetworkInterfaceByIDParams{
		Context: ctx,
	}
}

// NewPatchGuestNetworkInterfaceByIDParamsWithHTTPClient creates a new PatchGuestNetworkInterfaceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchGuestNetworkInterfaceByIDParamsWithHTTPClient(client *http.Client) *PatchGuestNetworkInterfaceByIDParams {
	return &PatchGuestNetworkInterfaceByIDParams{
		HTTPClient: client,
	}
}

/*
PatchGuestNetworkInterfaceByIDParams contains all the parameters to send to the API endpoint

	for the patch guest network interface by ID operation.

	Typically these are written to a http.Request.
*/
type PatchGuestNetworkInterfaceByIDParams struct {

	/* Body.

	   A subset of the guest network interface properties
	*/
	Body *models.PartialNetworkInterface

	/* IfaceID.

	   The id of the guest network interface
	*/
	IfaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch guest network interface by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGuestNetworkInterfaceByIDParams) WithDefaults() *PatchGuestNetworkInterfaceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch guest network interface by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGuestNetworkInterfaceByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithTimeout(timeout time.Duration) *PatchGuestNetworkInterfaceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithContext(ctx context.Context) *PatchGuestNetworkInterfaceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithHTTPClient(client *http.Client) *PatchGuestNetworkInterfaceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithBody(body *models.PartialNetworkInterface) *PatchGuestNetworkInterfaceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetBody(body *models.PartialNetworkInterface) {
	o.Body = body
}

// WithIfaceID adds the ifaceID to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithIfaceID(ifaceID string) *PatchGuestNetworkInterfaceByIDParams {
	o.SetIfaceID(ifaceID)
	return o
}

// SetIfaceID adds the ifaceId to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetIfaceID(ifaceID string) {
	o.IfaceID = ifaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGuestNetworkInterfaceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param iface_id
	if err := r.SetPathParam("iface_id", o.IfaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

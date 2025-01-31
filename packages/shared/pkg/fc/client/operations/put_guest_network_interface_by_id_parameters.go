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

// NewPutGuestNetworkInterfaceByIDParams creates a new PutGuestNetworkInterfaceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutGuestNetworkInterfaceByIDParams() *PutGuestNetworkInterfaceByIDParams {
	return &PutGuestNetworkInterfaceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutGuestNetworkInterfaceByIDParamsWithTimeout creates a new PutGuestNetworkInterfaceByIDParams object
// with the ability to set a timeout on a request.
func NewPutGuestNetworkInterfaceByIDParamsWithTimeout(timeout time.Duration) *PutGuestNetworkInterfaceByIDParams {
	return &PutGuestNetworkInterfaceByIDParams{
		timeout: timeout,
	}
}

// NewPutGuestNetworkInterfaceByIDParamsWithContext creates a new PutGuestNetworkInterfaceByIDParams object
// with the ability to set a context for a request.
func NewPutGuestNetworkInterfaceByIDParamsWithContext(ctx context.Context) *PutGuestNetworkInterfaceByIDParams {
	return &PutGuestNetworkInterfaceByIDParams{
		Context: ctx,
	}
}

// NewPutGuestNetworkInterfaceByIDParamsWithHTTPClient creates a new PutGuestNetworkInterfaceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutGuestNetworkInterfaceByIDParamsWithHTTPClient(client *http.Client) *PutGuestNetworkInterfaceByIDParams {
	return &PutGuestNetworkInterfaceByIDParams{
		HTTPClient: client,
	}
}

/*
PutGuestNetworkInterfaceByIDParams contains all the parameters to send to the API endpoint

	for the put guest network interface by ID operation.

	Typically these are written to a http.Request.
*/
type PutGuestNetworkInterfaceByIDParams struct {

	/* Body.

	   Guest network interface properties
	*/
	Body *models.NetworkInterface

	/* IfaceID.

	   The id of the guest network interface
	*/
	IfaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put guest network interface by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGuestNetworkInterfaceByIDParams) WithDefaults() *PutGuestNetworkInterfaceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put guest network interface by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGuestNetworkInterfaceByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) WithTimeout(timeout time.Duration) *PutGuestNetworkInterfaceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) WithContext(ctx context.Context) *PutGuestNetworkInterfaceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) WithHTTPClient(client *http.Client) *PutGuestNetworkInterfaceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) WithBody(body *models.NetworkInterface) *PutGuestNetworkInterfaceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) SetBody(body *models.NetworkInterface) {
	o.Body = body
}

// WithIfaceID adds the ifaceID to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) WithIfaceID(ifaceID string) *PutGuestNetworkInterfaceByIDParams {
	o.SetIfaceID(ifaceID)
	return o
}

// SetIfaceID adds the ifaceId to the put guest network interface by ID params
func (o *PutGuestNetworkInterfaceByIDParams) SetIfaceID(ifaceID string) {
	o.IfaceID = ifaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutGuestNetworkInterfaceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

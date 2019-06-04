// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPostVpcsVpcIDAddressPrefixesParams creates a new PostVpcsVpcIDAddressPrefixesParams object
// with the default values initialized.
func NewPostVpcsVpcIDAddressPrefixesParams() *PostVpcsVpcIDAddressPrefixesParams {
	var ()
	return &PostVpcsVpcIDAddressPrefixesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVpcsVpcIDAddressPrefixesParamsWithTimeout creates a new PostVpcsVpcIDAddressPrefixesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVpcsVpcIDAddressPrefixesParamsWithTimeout(timeout time.Duration) *PostVpcsVpcIDAddressPrefixesParams {
	var ()
	return &PostVpcsVpcIDAddressPrefixesParams{

		timeout: timeout,
	}
}

// NewPostVpcsVpcIDAddressPrefixesParamsWithContext creates a new PostVpcsVpcIDAddressPrefixesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVpcsVpcIDAddressPrefixesParamsWithContext(ctx context.Context) *PostVpcsVpcIDAddressPrefixesParams {
	var ()
	return &PostVpcsVpcIDAddressPrefixesParams{

		Context: ctx,
	}
}

// NewPostVpcsVpcIDAddressPrefixesParamsWithHTTPClient creates a new PostVpcsVpcIDAddressPrefixesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVpcsVpcIDAddressPrefixesParamsWithHTTPClient(client *http.Client) *PostVpcsVpcIDAddressPrefixesParams {
	var ()
	return &PostVpcsVpcIDAddressPrefixesParams{
		HTTPClient: client,
	}
}

/*PostVpcsVpcIDAddressPrefixesParams contains all the parameters to send to the API endpoint
for the post vpcs vpc ID address prefixes operation typically these are written to a http.Request
*/
type PostVpcsVpcIDAddressPrefixesParams struct {

	/*Body*/
	Body *models.PostVpcsVpcIDAddressPrefixesParamsBody
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) WithTimeout(timeout time.Duration) *PostVpcsVpcIDAddressPrefixesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) WithContext(ctx context.Context) *PostVpcsVpcIDAddressPrefixesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) WithHTTPClient(client *http.Client) *PostVpcsVpcIDAddressPrefixesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) WithBody(body *models.PostVpcsVpcIDAddressPrefixesParamsBody) *PostVpcsVpcIDAddressPrefixesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) SetBody(body *models.PostVpcsVpcIDAddressPrefixesParamsBody) {
	o.Body = body
}

// WithVersion adds the version to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) WithVersion(version string) *PostVpcsVpcIDAddressPrefixesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) WithVpcID(vpcID string) *PostVpcsVpcIDAddressPrefixesParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the post vpcs vpc ID address prefixes params
func (o *PostVpcsVpcIDAddressPrefixesParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *PostVpcsVpcIDAddressPrefixesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
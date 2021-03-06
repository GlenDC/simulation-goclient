package operations

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

// NewGetPartGeometryURLParams creates a new GetPartGeometryURLParams object
// with the default values initialized.
func NewGetPartGeometryURLParams() *GetPartGeometryURLParams {
	var ()
	return &GetPartGeometryURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPartGeometryURLParamsWithTimeout creates a new GetPartGeometryURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPartGeometryURLParamsWithTimeout(timeout time.Duration) *GetPartGeometryURLParams {
	var ()
	return &GetPartGeometryURLParams{

		timeout: timeout,
	}
}

// NewGetPartGeometryURLParamsWithContext creates a new GetPartGeometryURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPartGeometryURLParamsWithContext(ctx context.Context) *GetPartGeometryURLParams {
	var ()
	return &GetPartGeometryURLParams{

		Context: ctx,
	}
}

/*GetPartGeometryURLParams contains all the parameters to send to the API endpoint
for the get part geometry Url operation typically these are written to a http.Request
*/
type GetPartGeometryURLParams struct {

	/*ID
	  ID of part

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get part geometry Url params
func (o *GetPartGeometryURLParams) WithTimeout(timeout time.Duration) *GetPartGeometryURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get part geometry Url params
func (o *GetPartGeometryURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get part geometry Url params
func (o *GetPartGeometryURLParams) WithContext(ctx context.Context) *GetPartGeometryURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get part geometry Url params
func (o *GetPartGeometryURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the get part geometry Url params
func (o *GetPartGeometryURLParams) WithID(id int32) *GetPartGeometryURLParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get part geometry Url params
func (o *GetPartGeometryURLParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPartGeometryURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

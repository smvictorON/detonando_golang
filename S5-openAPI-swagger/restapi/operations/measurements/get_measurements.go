// Code generated by go-swagger; DO NOT EDIT.

package measurements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMeasurementsHandlerFunc turns a function with the right signature into a get measurements handler
type GetMeasurementsHandlerFunc func(GetMeasurementsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMeasurementsHandlerFunc) Handle(params GetMeasurementsParams) middleware.Responder {
	return fn(params)
}

// GetMeasurementsHandler interface for that can handle valid get measurements params
type GetMeasurementsHandler interface {
	Handle(GetMeasurementsParams) middleware.Responder
}

// NewGetMeasurements creates a new http.Handler for the get measurements operation
func NewGetMeasurements(ctx *middleware.Context, handler GetMeasurementsHandler) *GetMeasurements {
	return &GetMeasurements{Context: ctx, Handler: handler}
}

/*
	GetMeasurements swagger:route GET /measurements measurements getMeasurements

GetMeasurements get measurements API
*/
type GetMeasurements struct {
	Context *middleware.Context
	Handler GetMeasurementsHandler
}

func (o *GetMeasurements) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMeasurementsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
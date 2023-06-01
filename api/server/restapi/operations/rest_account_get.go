// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RestAccountGetHandlerFunc turns a function with the right signature into a rest account get handler
type RestAccountGetHandlerFunc func(RestAccountGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RestAccountGetHandlerFunc) Handle(params RestAccountGetParams) middleware.Responder {
	return fn(params)
}

// RestAccountGetHandler interface for that can handle valid rest account get params
type RestAccountGetHandler interface {
	Handle(RestAccountGetParams) middleware.Responder
}

// NewRestAccountGet creates a new http.Handler for the rest account get operation
func NewRestAccountGet(ctx *middleware.Context, handler RestAccountGetHandler) *RestAccountGet {
	return &RestAccountGet{Context: ctx, Handler: handler}
}

/*
	RestAccountGet swagger:route GET /api/accounts/{nickname} restAccountGet

Get the account by nickname given in path, with the option to return the ciphered private key.
*/
type RestAccountGet struct {
	Context *middleware.Context
	Handler RestAccountGetHandler
}

func (o *RestAccountGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRestAccountGetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
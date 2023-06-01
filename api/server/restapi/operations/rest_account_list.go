// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RestAccountListHandlerFunc turns a function with the right signature into a rest account list handler
type RestAccountListHandlerFunc func(RestAccountListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RestAccountListHandlerFunc) Handle(params RestAccountListParams) middleware.Responder {
	return fn(params)
}

// RestAccountListHandler interface for that can handle valid rest account list params
type RestAccountListHandler interface {
	Handle(RestAccountListParams) middleware.Responder
}

// NewRestAccountList creates a new http.Handler for the rest account list operation
func NewRestAccountList(ctx *middleware.Context, handler RestAccountListHandler) *RestAccountList {
	return &RestAccountList{Context: ctx, Handler: handler}
}

/*
	RestAccountList swagger:route GET /api/accounts restAccountList

Get all the accounts.
*/
type RestAccountList struct {
	Context *middleware.Context
	Handler RestAccountListHandler
}

func (o *RestAccountList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRestAccountListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
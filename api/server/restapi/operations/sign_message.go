// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SignMessageHandlerFunc turns a function with the right signature into a sign message handler
type SignMessageHandlerFunc func(SignMessageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SignMessageHandlerFunc) Handle(params SignMessageParams) middleware.Responder {
	return fn(params)
}

// SignMessageHandler interface for that can handle valid sign message params
type SignMessageHandler interface {
	Handle(SignMessageParams) middleware.Responder
}

// NewSignMessage creates a new http.Handler for the sign message operation
func NewSignMessage(ctx *middleware.Context, handler SignMessageHandler) *SignMessage {
	return &SignMessage{Context: ctx, Handler: handler}
}

/*
	SignMessage swagger:route POST /api/accounts/{nickname}/signMessage signMessage

Sign a message using the account associated with the provided nickname in the path.
*/
type SignMessage struct {
	Context *middleware.Context
	Handler SignMessageHandler
}

func (o *SignMessage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSignMessageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
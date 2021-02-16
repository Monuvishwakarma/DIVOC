// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostV1AuthorizeHandlerFunc turns a function with the right signature into a post v1 authorize handler
type PostV1AuthorizeHandlerFunc func(PostV1AuthorizeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostV1AuthorizeHandlerFunc) Handle(params PostV1AuthorizeParams) middleware.Responder {
	return fn(params)
}

// PostV1AuthorizeHandler interface for that can handle valid post v1 authorize params
type PostV1AuthorizeHandler interface {
	Handle(PostV1AuthorizeParams) middleware.Responder
}

// NewPostV1Authorize creates a new http.Handler for the post v1 authorize operation
func NewPostV1Authorize(ctx *middleware.Context, handler PostV1AuthorizeHandler) *PostV1Authorize {
	return &PostV1Authorize{Context: ctx, Handler: handler}
}

/*PostV1Authorize swagger:route POST /v1/authorize login postV1Authorize

Establish token

*/
type PostV1Authorize struct {
	Context *middleware.Context
	Handler PostV1AuthorizeHandler
}

func (o *PostV1Authorize) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostV1AuthorizeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

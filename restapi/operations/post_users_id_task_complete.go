// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUsersIDTaskCompleteHandlerFunc turns a function with the right signature into a post users ID task complete handler
type PostUsersIDTaskCompleteHandlerFunc func(PostUsersIDTaskCompleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUsersIDTaskCompleteHandlerFunc) Handle(params PostUsersIDTaskCompleteParams) middleware.Responder {
	return fn(params)
}

// PostUsersIDTaskCompleteHandler interface for that can handle valid post users ID task complete params
type PostUsersIDTaskCompleteHandler interface {
	Handle(PostUsersIDTaskCompleteParams) middleware.Responder
}

// NewPostUsersIDTaskComplete creates a new http.Handler for the post users ID task complete operation
func NewPostUsersIDTaskComplete(ctx *middleware.Context, handler PostUsersIDTaskCompleteHandler) *PostUsersIDTaskComplete {
	return &PostUsersIDTaskComplete{Context: ctx, Handler: handler}
}

/*
	PostUsersIDTaskComplete swagger:route POST /users/{id}/task/complete postUsersIdTaskComplete

Completion user's task
*/
type PostUsersIDTaskComplete struct {
	Context *middleware.Context
	Handler PostUsersIDTaskCompleteHandler
}

func (o *PostUsersIDTaskComplete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUsersIDTaskCompleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

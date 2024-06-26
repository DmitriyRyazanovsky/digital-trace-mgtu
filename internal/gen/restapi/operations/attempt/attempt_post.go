// Code generated by go-swagger; DO NOT EDIT.

package attempt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttemptPostHandlerFunc turns a function with the right signature into a attempt post handler
type AttemptPostHandlerFunc func(AttemptPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AttemptPostHandlerFunc) Handle(params AttemptPostParams) middleware.Responder {
	return fn(params)
}

// AttemptPostHandler interface for that can handle valid attempt post params
type AttemptPostHandler interface {
	Handle(AttemptPostParams) middleware.Responder
}

// NewAttemptPost creates a new http.Handler for the attempt post operation
func NewAttemptPost(ctx *middleware.Context, handler AttemptPostHandler) *AttemptPost {
	return &AttemptPost{Context: ctx, Handler: handler}
}

/*
AttemptPost swagger:route POST /user/attempt attempt attemptPost

Запрос на создание совой попытки
*/
type AttemptPost struct {
	Context *middleware.Context
	Handler AttemptPostHandler
}

func (o *AttemptPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAttemptPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AttemptPostBody attempt post body
//
// swagger:model AttemptPostBody
type AttemptPostBody struct {

	// test id
	TestID uint64 `json:"test_id,omitempty"`
}

// Validate validates this attempt post body
func (o *AttemptPostBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AttemptPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AttemptPostBody) UnmarshalBinary(b []byte) error {
	var res AttemptPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

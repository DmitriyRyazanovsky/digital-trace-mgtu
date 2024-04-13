// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserAttemptGetHandlerFunc turns a function with the right signature into a user attempt get handler
type UserAttemptGetHandlerFunc func(UserAttemptGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserAttemptGetHandlerFunc) Handle(params UserAttemptGetParams) middleware.Responder {
	return fn(params)
}

// UserAttemptGetHandler interface for that can handle valid user attempt get params
type UserAttemptGetHandler interface {
	Handle(UserAttemptGetParams) middleware.Responder
}

// NewUserAttemptGet creates a new http.Handler for the user attempt get operation
func NewUserAttemptGet(ctx *middleware.Context, handler UserAttemptGetHandler) *UserAttemptGet {
	return &UserAttemptGet{Context: ctx, Handler: handler}
}

/*
	UserAttemptGet swagger:route GET /user/attempt user userAttemptGet

Запрос на поиск информации о попытках прохождения теста пользователем
*/
type UserAttemptGet struct {
	Context *middleware.Context
	Handler UserAttemptGetHandler
}

func (o *UserAttemptGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUserAttemptGetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UserAttemptGetOKBodyItems0 user attempt get o k body items0
//
// swagger:model UserAttemptGetOKBodyItems0
type UserAttemptGetOKBodyItems0 struct {

	// attempt id
	AttemptID uint64 `json:"attempt_id,omitempty"`

	// status id
	StatusID uint64 `json:"status_id,omitempty"`

	// test id
	TestID uint64 `json:"test_id,omitempty"`
}

// Validate validates this user attempt get o k body items0
func (o *UserAttemptGetOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user attempt get o k body items0 based on context it is used
func (o *UserAttemptGetOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UserAttemptGetOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserAttemptGetOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UserAttemptGetOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

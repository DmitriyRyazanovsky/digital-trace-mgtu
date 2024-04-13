// Code generated by go-swagger; DO NOT EDIT.

package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MailUserPostHandlerFunc turns a function with the right signature into a mail user post handler
type MailUserPostHandlerFunc func(MailUserPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MailUserPostHandlerFunc) Handle(params MailUserPostParams) middleware.Responder {
	return fn(params)
}

// MailUserPostHandler interface for that can handle valid mail user post params
type MailUserPostHandler interface {
	Handle(MailUserPostParams) middleware.Responder
}

// NewMailUserPost creates a new http.Handler for the mail user post operation
func NewMailUserPost(ctx *middleware.Context, handler MailUserPostHandler) *MailUserPost {
	return &MailUserPost{Context: ctx, Handler: handler}
}

/*
	MailUserPost swagger:route POST /mail/user mail mailUserPost

Запрос на создание письма верефикации по email для пользователя
*/
type MailUserPost struct {
	Context *middleware.Context
	Handler MailUserPostHandler
}

func (o *MailUserPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMailUserPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// MailUserPostBody mail user post body
//
// swagger:model MailUserPostBody
type MailUserPostBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// login
	// Required: true
	Login *string `json:"login"`

	// name
	// Required: true
	Name *string `json:"name"`

	// password
	// Required: true
	Password *string `json:"password"`

	// role
	// Required: true
	Role *string `json:"role"`

	// surname
	// Required: true
	Surname *string `json:"surname"`
}

// Validate validates this mail user post body
func (o *MailUserPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSurname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MailUserPostBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *MailUserPostBody) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"login", "body", o.Login); err != nil {
		return err
	}

	return nil
}

func (o *MailUserPostBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *MailUserPostBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *MailUserPostBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	return nil
}

func (o *MailUserPostBody) validateSurname(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"surname", "body", o.Surname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mail user post body based on context it is used
func (o *MailUserPostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MailUserPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MailUserPostBody) UnmarshalBinary(b []byte) error {
	var res MailUserPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// MailUserPostOKBody mail user post o k body
//
// swagger:model MailUserPostOKBody
type MailUserPostOKBody struct {

	// success
	Success string `json:"success,omitempty"`
}

// Validate validates this mail user post o k body
func (o *MailUserPostOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mail user post o k body based on context it is used
func (o *MailUserPostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MailUserPostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MailUserPostOKBody) UnmarshalBinary(b []byte) error {
	var res MailUserPostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

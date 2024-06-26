// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserProfileGetHandlerFunc turns a function with the right signature into a user profile get handler
type UserProfileGetHandlerFunc func(UserProfileGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserProfileGetHandlerFunc) Handle(params UserProfileGetParams) middleware.Responder {
	return fn(params)
}

// UserProfileGetHandler interface for that can handle valid user profile get params
type UserProfileGetHandler interface {
	Handle(UserProfileGetParams) middleware.Responder
}

// NewUserProfileGet creates a new http.Handler for the user profile get operation
func NewUserProfileGet(ctx *middleware.Context, handler UserProfileGetHandler) *UserProfileGet {
	return &UserProfileGet{Context: ctx, Handler: handler}
}

/*
UserProfileGet swagger:route GET /user/profile user userProfileGet

Запрос на получение личной информации о пользователе по acces токену
*/
type UserProfileGet struct {
	Context *middleware.Context
	Handler UserProfileGetHandler
}

func (o *UserProfileGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserProfileGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UserProfileGetOKBody user profile get o k body
//
// swagger:model UserProfileGetOKBody
type UserProfileGetOKBody struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// role id
	RoleID uint64 `json:"role_id,omitempty"`

	// surname
	Surname string `json:"surname,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this user profile get o k body
func (o *UserProfileGetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserProfileGetOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("userProfileGetOK"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserProfileGetOKBody) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("userProfileGetOK"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UserProfileGetOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("userProfileGetOK"+"."+"updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserProfileGetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserProfileGetOKBody) UnmarshalBinary(b []byte) error {
	var res UserProfileGetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

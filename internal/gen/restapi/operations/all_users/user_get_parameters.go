// Code generated by go-swagger; DO NOT EDIT.

package all_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUserGetParams creates a new UserGetParams object
// no default values defined in spec.
func NewUserGetParams() UserGetParams {

	return UserGetParams{}
}

// UserGetParams contains all the bound params for the user get operation
// typically these are obtained from a http.Request
//
// swagger:parameters userGet
type UserGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Поиск по email пользователя
	  In: query
	*/
	Email *string
	/*Поиск по Id пользователя
	  In: query
	*/
	ID *uint64
	/*Поиск по логину пользователя
	  In: query
	*/
	Login *string
	/*Поиск по имени пользователя
	  In: query
	*/
	Name *string
	/*Поиск по роли пользователя
	  In: query
	*/
	RoleID *uint64
	/*Поиск по фамилии пользователя
	  In: query
	*/
	Surname *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUserGetParams() beforehand.
func (o *UserGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEmail, qhkEmail, _ := qs.GetOK("email")
	if err := o.bindEmail(qEmail, qhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLogin, qhkLogin, _ := qs.GetOK("login")
	if err := o.bindLogin(qLogin, qhkLogin, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qRoleID, qhkRoleID, _ := qs.GetOK("role_id")
	if err := o.bindRoleID(qRoleID, qhkRoleID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSurname, qhkSurname, _ := qs.GetOK("surname")
	if err := o.bindSurname(qSurname, qhkSurname, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmail binds and validates parameter Email from query.
func (o *UserGetParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Email = &raw

	return nil
}

// bindID binds and validates parameter ID from query.
func (o *UserGetParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("id", "query", "uint64", raw)
	}
	o.ID = &value

	return nil
}

// bindLogin binds and validates parameter Login from query.
func (o *UserGetParams) bindLogin(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Login = &raw

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *UserGetParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	return nil
}

// bindRoleID binds and validates parameter RoleID from query.
func (o *UserGetParams) bindRoleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("role_id", "query", "uint64", raw)
	}
	o.RoleID = &value

	return nil
}

// bindSurname binds and validates parameter Surname from query.
func (o *UserGetParams) bindSurname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Surname = &raw

	return nil
}

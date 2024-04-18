// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewUserProfilePatchParams creates a new UserProfilePatchParams object
//
// There are no default values defined in the spec.
func NewUserProfilePatchParams() UserProfilePatchParams {

	return UserProfilePatchParams{}
}

// UserProfilePatchParams contains all the bound params for the user profile patch operation
// typically these are obtained from a http.Request
//
// swagger:parameters user_profile_patch
type UserProfilePatchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*jwt auth
	  Required: true
	  In: header
	*/
	Authorization string
	/*email
	  In: query
	*/
	Email *strfmt.Email
	/*id
	  In: query
	*/
	ID *uint64
	/*login
	  In: query
	*/
	Login *string
	/*name
	  In: query
	*/
	Name *string
	/*password
	  In: query
	*/
	Password *string
	/*role_id
	  In: query
	*/
	RoleID *uint64
	/*surname
	  In: query
	*/
	Surname *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUserProfilePatchParams() beforehand.
func (o *UserProfilePatchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

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

	qPassword, qhkPassword, _ := qs.GetOK("password")
	if err := o.bindPassword(qPassword, qhkPassword, route.Formats); err != nil {
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

// bindAuthorization binds and validates parameter Authorization from header.
func (o *UserProfilePatchParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Authorization", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Authorization", "header", raw); err != nil {
		return err
	}
	o.Authorization = raw

	return nil
}

// bindEmail binds and validates parameter Email from query.
func (o *UserProfilePatchParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: email
	value, err := formats.Parse("email", raw)
	if err != nil {
		return errors.InvalidType("email", "query", "strfmt.Email", raw)
	}
	o.Email = (value.(*strfmt.Email))

	if err := o.validateEmail(formats); err != nil {
		return err
	}

	return nil
}

// validateEmail carries on validations for parameter Email
func (o *UserProfilePatchParams) validateEmail(formats strfmt.Registry) error {

	if err := validate.FormatOf("email", "query", "email", o.Email.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindID binds and validates parameter ID from query.
func (o *UserProfilePatchParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UserProfilePatchParams) bindLogin(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UserProfilePatchParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindPassword binds and validates parameter Password from query.
func (o *UserProfilePatchParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Password = &raw

	return nil
}

// bindRoleID binds and validates parameter RoleID from query.
func (o *UserProfilePatchParams) bindRoleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UserProfilePatchParams) bindSurname(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

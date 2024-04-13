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

// NewUserAttemptGetParams creates a new UserAttemptGetParams object
//
// There are no default values defined in the spec.
func NewUserAttemptGetParams() UserAttemptGetParams {

	return UserAttemptGetParams{}
}

// UserAttemptGetParams contains all the bound params for the user attempt get operation
// typically these are obtained from a http.Request
//
// swagger:parameters userAttemptGet
type UserAttemptGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*jwt auth
	  Required: true
	  In: header
	*/
	Authorization string
	/*
	  In: query
	*/
	StatusID *uint64
	/*
	  In: query
	*/
	TestID *uint64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUserAttemptGetParams() beforehand.
func (o *UserAttemptGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qStatusID, qhkStatusID, _ := qs.GetOK("status_id")
	if err := o.bindStatusID(qStatusID, qhkStatusID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTestID, qhkTestID, _ := qs.GetOK("test_id")
	if err := o.bindTestID(qTestID, qhkTestID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *UserAttemptGetParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindStatusID binds and validates parameter StatusID from query.
func (o *UserAttemptGetParams) bindStatusID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("status_id", "query", "uint64", raw)
	}
	o.StatusID = &value

	return nil
}

// bindTestID binds and validates parameter TestID from query.
func (o *UserAttemptGetParams) bindTestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("test_id", "query", "uint64", raw)
	}
	o.TestID = &value

	return nil
}

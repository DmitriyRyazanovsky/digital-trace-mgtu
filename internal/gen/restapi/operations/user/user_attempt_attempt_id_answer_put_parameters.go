// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewUserAttemptAttemptIDAnswerPutParams creates a new UserAttemptAttemptIDAnswerPutParams object
//
// There are no default values defined in the spec.
func NewUserAttemptAttemptIDAnswerPutParams() UserAttemptAttemptIDAnswerPutParams {

	return UserAttemptAttemptIDAnswerPutParams{}
}

// UserAttemptAttemptIDAnswerPutParams contains all the bound params for the user attempt attempt Id answer put operation
// typically these are obtained from a http.Request
//
// swagger:parameters userAttemptAttemptIdAnswerPut
type UserAttemptAttemptIDAnswerPutParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*jwt access auth
	  Required: true
	  In: header
	*/
	Authorization string
	/*
	  Required: true
	  In: path
	*/
	AttemptID uint64
	/*
	  Required: true
	  In: body
	*/
	Body UserAttemptAttemptIDAnswerPutBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUserAttemptAttemptIDAnswerPutParams() beforehand.
func (o *UserAttemptAttemptIDAnswerPutParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rAttemptID, rhkAttemptID, _ := route.Params.GetOK("attempt_id")
	if err := o.bindAttemptID(rAttemptID, rhkAttemptID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body UserAttemptAttemptIDAnswerPutBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *UserAttemptAttemptIDAnswerPutParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindAttemptID binds and validates parameter AttemptID from path.
func (o *UserAttemptAttemptIDAnswerPutParams) bindAttemptID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("attempt_id", "path", "uint64", raw)
	}
	o.AttemptID = value

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestTestIDQuestionGetHandlerFunc turns a function with the right signature into a test test Id question get handler
type TestTestIDQuestionGetHandlerFunc func(TestTestIDQuestionGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TestTestIDQuestionGetHandlerFunc) Handle(params TestTestIDQuestionGetParams) middleware.Responder {
	return fn(params)
}

// TestTestIDQuestionGetHandler interface for that can handle valid test test Id question get params
type TestTestIDQuestionGetHandler interface {
	Handle(TestTestIDQuestionGetParams) middleware.Responder
}

// NewTestTestIDQuestionGet creates a new http.Handler for the test test Id question get operation
func NewTestTestIDQuestionGet(ctx *middleware.Context, handler TestTestIDQuestionGetHandler) *TestTestIDQuestionGet {
	return &TestTestIDQuestionGet{Context: ctx, Handler: handler}
}

/*
TestTestIDQuestionGet swagger:route GET /test/{test_id}/questions test testTestIdQuestionGet

Запрос на получение списка всех вопросов по тесту
*/
type TestTestIDQuestionGet struct {
	Context *middleware.Context
	Handler TestTestIDQuestionGetHandler
}

func (o *TestTestIDQuestionGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTestTestIDQuestionGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// TestTestIDQuestionGetOKBodyItems0 test test ID question get o k body items0
//
// swagger:model TestTestIDQuestionGetOKBodyItems0
type TestTestIDQuestionGetOKBodyItems0 struct {

	// answers
	Answers []*TestTestIDQuestionGetOKBodyItems0AnswersItems0 `json:"answers"`

	// button type
	ButtonType uint64 `json:"button_type,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// question id
	QuestionID uint64 `json:"question_id,omitempty"`
}

// Validate validates this test test ID question get o k body items0
func (o *TestTestIDQuestionGetOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestTestIDQuestionGetOKBodyItems0) validateAnswers(formats strfmt.Registry) error {

	if swag.IsZero(o.Answers) { // not required
		return nil
	}

	for i := 0; i < len(o.Answers); i++ {
		if swag.IsZero(o.Answers[i]) { // not required
			continue
		}

		if o.Answers[i] != nil {
			if err := o.Answers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("answers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TestTestIDQuestionGetOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestTestIDQuestionGetOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res TestTestIDQuestionGetOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// TestTestIDQuestionGetOKBodyItems0AnswersItems0 test test ID question get o k body items0 answers items0
//
// swagger:model TestTestIDQuestionGetOKBodyItems0AnswersItems0
type TestTestIDQuestionGetOKBodyItems0AnswersItems0 struct {

	// content
	Content string `json:"content,omitempty"`
}

// Validate validates this test test ID question get o k body items0 answers items0
func (o *TestTestIDQuestionGetOKBodyItems0AnswersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TestTestIDQuestionGetOKBodyItems0AnswersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestTestIDQuestionGetOKBodyItems0AnswersItems0) UnmarshalBinary(b []byte) error {
	var res TestTestIDQuestionGetOKBodyItems0AnswersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

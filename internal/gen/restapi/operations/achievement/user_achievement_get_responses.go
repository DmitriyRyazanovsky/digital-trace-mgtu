// Code generated by go-swagger; DO NOT EDIT.

package achievement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mgtu/digital-trace/main-backend-service/internal/gen/models"
)

// UserAchievementGetOKCode is the HTTP code returned for type UserAchievementGetOK
const UserAchievementGetOKCode int = 200

/*
UserAchievementGetOK информация о достижениях успешно получена

swagger:response userAchievementGetOK
*/
type UserAchievementGetOK struct {

	/*
	  In: Body
	*/
	Payload []*UserAchievementGetOKBodyItems0 `json:"body,omitempty"`
}

// NewUserAchievementGetOK creates UserAchievementGetOK with default headers values
func NewUserAchievementGetOK() *UserAchievementGetOK {

	return &UserAchievementGetOK{}
}

// WithPayload adds the payload to the user achievement get o k response
func (o *UserAchievementGetOK) WithPayload(payload []*UserAchievementGetOKBodyItems0) *UserAchievementGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user achievement get o k response
func (o *UserAchievementGetOK) SetPayload(payload []*UserAchievementGetOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserAchievementGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*UserAchievementGetOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UserAchievementGetInternalServerErrorCode is the HTTP code returned for type UserAchievementGetInternalServerError
const UserAchievementGetInternalServerErrorCode int = 500

/*
UserAchievementGetInternalServerError Ошибка сервера либо запроса

swagger:response userAchievementGetInternalServerError
*/
type UserAchievementGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUserAchievementGetInternalServerError creates UserAchievementGetInternalServerError with default headers values
func NewUserAchievementGetInternalServerError() *UserAchievementGetInternalServerError {

	return &UserAchievementGetInternalServerError{}
}

// WithPayload adds the payload to the user achievement get internal server error response
func (o *UserAchievementGetInternalServerError) WithPayload(payload *models.Error500) *UserAchievementGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user achievement get internal server error response
func (o *UserAchievementGetInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserAchievementGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
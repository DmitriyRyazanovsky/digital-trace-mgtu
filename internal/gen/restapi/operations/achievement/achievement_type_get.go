// Code generated by go-swagger; DO NOT EDIT.

package achievement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AchievementTypeGetHandlerFunc turns a function with the right signature into a achievement type get handler
type AchievementTypeGetHandlerFunc func(AchievementTypeGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AchievementTypeGetHandlerFunc) Handle(params AchievementTypeGetParams) middleware.Responder {
	return fn(params)
}

// AchievementTypeGetHandler interface for that can handle valid achievement type get params
type AchievementTypeGetHandler interface {
	Handle(AchievementTypeGetParams) middleware.Responder
}

// NewAchievementTypeGet creates a new http.Handler for the achievement type get operation
func NewAchievementTypeGet(ctx *middleware.Context, handler AchievementTypeGetHandler) *AchievementTypeGet {
	return &AchievementTypeGet{Context: ctx, Handler: handler}
}

/*
AchievementTypeGet swagger:route GET /achievement/type achievement achievementTypeGet

Запрос на поиск типов достижения
*/
type AchievementTypeGet struct {
	Context *middleware.Context
	Handler AchievementTypeGetHandler
}

func (o *AchievementTypeGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAchievementTypeGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AchievementTypeGetOKBodyItems0 achievement type get o k body items0
//
// swagger:model AchievementTypeGetOKBodyItems0
type AchievementTypeGetOKBodyItems0 struct {

	// achive type id
	AchiveTypeID uint64 `json:"achive_type_id,omitempty"`

	// achive type name
	AchiveTypeName string `json:"achive_type_name,omitempty"`
}

// Validate validates this achievement type get o k body items0
func (o *AchievementTypeGetOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AchievementTypeGetOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AchievementTypeGetOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res AchievementTypeGetOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

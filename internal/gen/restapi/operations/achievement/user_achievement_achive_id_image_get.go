// Code generated by go-swagger; DO NOT EDIT.

package achievement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UserAchievementAchiveIDImageGetHandlerFunc turns a function with the right signature into a user achievement achive Id image get handler
type UserAchievementAchiveIDImageGetHandlerFunc func(UserAchievementAchiveIDImageGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserAchievementAchiveIDImageGetHandlerFunc) Handle(params UserAchievementAchiveIDImageGetParams) middleware.Responder {
	return fn(params)
}

// UserAchievementAchiveIDImageGetHandler interface for that can handle valid user achievement achive Id image get params
type UserAchievementAchiveIDImageGetHandler interface {
	Handle(UserAchievementAchiveIDImageGetParams) middleware.Responder
}

// NewUserAchievementAchiveIDImageGet creates a new http.Handler for the user achievement achive Id image get operation
func NewUserAchievementAchiveIDImageGet(ctx *middleware.Context, handler UserAchievementAchiveIDImageGetHandler) *UserAchievementAchiveIDImageGet {
	return &UserAchievementAchiveIDImageGet{Context: ctx, Handler: handler}
}

/*
UserAchievementAchiveIDImageGet swagger:route GET /user/achievement/{achive_id}/image achievement userAchievementAchiveIdImageGet

Запрос на получение ссылки на скачивание изображения достижения
*/
type UserAchievementAchiveIDImageGet struct {
	Context *middleware.Context
	Handler UserAchievementAchiveIDImageGetHandler
}

func (o *UserAchievementAchiveIDImageGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserAchievementAchiveIDImageGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

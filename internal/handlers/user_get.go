package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/all_users"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userGet500(err error) middleware.Responder {
	err = errors.Wrap(err, "handler error: [userGet]")
	h.log.Error(err.Error())
	return all_users.NewUserGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userGet(params all_users.UserGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	findUserOut, err := h.db.FindUser(tx, database.User{
		Id:      params.ID,
		RoleId:  params.RoleID,
		Email:   params.Email,
		Login:   params.Login,
		Name:    params.Name,
		Surname: params.Surname,
	})
	if err != nil {
		return h.userGet500(errors.Wrap(err, "[h.db.FindUser]"))
	}

	out := []*all_users.UserGetOKBodyItems0{}

	for _, v := range findUserOut.User {
		item := &all_users.UserGetOKBodyItems0{
			ID:      *v.Id,
			Email:   *v.Email,
			Login:   *v.Login,
			Name:    *v.Name,
			Surname: *v.Surname,
		}

		out = append(out, item)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return all_users.NewUserGetOK().WithPayload(out)
}

package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"
	"mgtu/digital-trace/main-backend-service/internal/password_hasher"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userPost500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userPost]")
	h.log.Error(err.Error())
	return user.NewUserPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userPost(params user.UserPostParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userPost500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	hashedPass, err := password_hasher.HashPassword(*params.Body.Password)
	if err != nil {
		return h.userPost500(errors.Wrap(err, "[password_hasher.HashPassword(params.Body.Password)]"))
	}

	adduserOut, err := h.db.AddUser(tx, database.User{
		RoleId:   params.Body.RoleID,
		Email:    params.Body.Email,
		Login:    params.Body.Login,
		Name:     params.Body.Name,
		Surname:  params.Body.Surname,
		Password: &hashedPass,
	})
	if err != nil {
		return h.userPost500(errors.Wrap(err, "[h.db.AddUser]"))
	}

	out := &user.UserPostOKBody{
		UserID: *adduserOut.User.Id,
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userPost500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserPostOK().WithPayload(out)
}

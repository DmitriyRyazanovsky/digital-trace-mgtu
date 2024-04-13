package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

func (h *Handler) userProfileGet500(err error) middleware.Responder {
	err = errors.Wrap(err, "handler error: [userProfileGet]")
	h.log.Error(err.Error())
	return user.NewUserPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userProfileGet(params user.UserProfileGetParams) middleware.Responder {
	claims, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		err = errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]")
		return h.userProfileGet500(err)
	}
	tx, err := h.db.OpenTransaction()
	if err != nil {
		err = errors.Wrap(err, "[h.db.OpenTransaction()]")
		return h.userProfileGet500(err)
	}
	findUserOut, err := h.db.FindUser(tx, database.User{
		Id: &claims.UserId,
	})
	if err != nil {
		err = errors.Wrap(err, "[h.db.FindUser()]")
		return h.userProfileGet500(err)
	}
	if !findUserOut.IsFound {
		err = errors.New("unable find user by id")
		return h.userProfileGet500(err)
	}
	if len(findUserOut.User) != 1 {
		err = errors.New("len(findUserOut.User) != 1")
		return h.userProfileGet500(err)
	}
	userInfo := findUserOut.User[0]
	out := &user.UserProfileGetOKBody{
		ID:        *userInfo.Id,
		CreatedAt: *userInfo.CreatedAt,
		Email:     strfmt.Email(*userInfo.Email),
		Login:     *userInfo.Login,
		Name:      *userInfo.Name,
		Surname:   *userInfo.Surname,
		Password:  *userInfo.Password,
		RoleID:    *userInfo.RoleId,
		UpdatedAt: *userInfo.UpdatedAt,
	}
	err = h.db.CommitTransaction(tx)
	if err != nil {
		err = errors.Wrap(err, "[h.db.CommitTransaction(tx)]")
		return h.userProfileGet500(err)
	}
	return user.NewUserProfileGetOK().WithPayload(out)
}

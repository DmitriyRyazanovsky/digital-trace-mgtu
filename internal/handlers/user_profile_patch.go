package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userProfilePatch500(err error) middleware.Responder {
	err = errors.Wrap(err, "handler error: [userProfilePatch]")
	h.log.Error(err.Error())
	return user.NewUserPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userProfilePatch(params user.UserProfilePatchParams) middleware.Responder {
	claims, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		err = errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]")
		return h.userProfilePatch500(err)
	}

	tx, err := h.db.OpenTransaction()
	if err != nil {
		err = errors.Wrap(err, "[h.db.OpenTransaction()]")
		return h.userProfilePatch500(err)
	}

	_, err = h.db.ChangeUser(tx, database.User{
		Login:   params.Login,
		Name:    params.Name,
		Surname: params.Surname,
	}, database.User{
		Id: &claims.UserId,
	})
	err = h.db.CommitTransaction(tx)
	if err != nil {
		err = errors.Wrap(err, "[h.db.CommitTransaction(tx)]")
		return h.userProfilePatch500(err)
	}

	return user.NewUserProfilePatchOK()
}

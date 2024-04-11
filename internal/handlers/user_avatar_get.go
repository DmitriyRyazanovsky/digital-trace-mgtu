package handlers

import (
	"fmt"
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) UserAvatarGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAvatarPut]")
	h.log.Error(err.Error())
	return user.NewUserAvatarGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAvatarGet(params user.UserAvatarGetParams) middleware.Responder {
	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.UserAvatarPut500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	findUserAvatar, err := h.db.FindUserAvatar(tx, database.UserAvatar{
		UserId: &accessToken.UserId,
	})
	if err != nil {
		err = errors.Wrap(err, "[h.db.FindUserAvatar(tx, {...})]")
		return h.UserAvatarPut500(err)
	}
	if !findUserAvatar.IsFound {
		err = errors.New("[!findUserAvatar.IsFound]")
		return h.UserAvatarPut500(err)
	}
	if len(findUserAvatar.UserAvatar) != 1 {
		err = errors.New("[len(findUserAvatar.UserAvatar) != 1]")
		return h.UserAvatarPut500(err)
	}

	findUserAvatarItem := findUserAvatar.UserAvatar[0]

	data, _, err := h.fileWorker.UserAvatarRead(accessToken.UserId, *findUserAvatarItem.Prefix)
	if err != nil {
		err = errors.Wrap(err, "[h.fileWorker.UserAvatarWrite(accessToken.UserId, data)]")
		return h.UserAvatarPut500(err)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.UserAvatarPut500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserAvatarGetOK().
		WithPayload(string(data)).
		WithContentDisposition(fmt.Sprintf("attachment; filename=\"image.%s\"", *findUserAvatarItem.Prefix))
}

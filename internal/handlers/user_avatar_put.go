package handlers

import (
	"encoding/base64"
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) UserAvatarPut500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAvatarPut]")
	h.log.Error(err.Error())
	return user.NewUserAvatarPutInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAvatarPut(params user.UserAvatarPutParams) middleware.Responder {
	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	//* Раскодируем base64 файл
	data, err := base64.StdEncoding.DecodeString(params.Body.ImageBase64)
	if err != nil {
		err = errors.Wrap(err, "[base64.StdEncoding.DecodeString(base64String)]")
		return h.UserAvatarPut500(err)
	}

	//* узнаём путь до сохранения файла
	path := h.fileWorker.GenUserAvatarPath(accessToken.UserId)

	//* открываем транзакцию
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.UserAvatarPut500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	//* ищем аватарку пользователя по его id из access токена
	findUserAvatar, err := h.db.FindUserAvatar(tx, database.UserAvatar{
		UserId: &accessToken.UserId,
	})
	if err != nil {
		err = errors.Wrap(err, "[h.db.FindUserAvatar(tx, {...})]")
		return h.UserAvatarPut500(err)
	}
	if !findUserAvatar.IsFound {
		//* если пользователь ещё не заводил аватарку, то добавляем аватарку пользователя
		_, err = h.db.AddUserAvatar(tx, database.UserAvatar{
			UserId: &accessToken.UserId,
			Prefix: &params.Body.Prefix,
			Path:   &path,
		})
		if err != nil {
			err = errors.Wrap(err, "[h.db.AddUserAvatar(tx, {...}]")
			return h.UserAvatarPut500(err)
		}
	} else {
		//* если же аватарка существует, то мы должны убедиться, что длина не больше одной аватарки
		if len(findUserAvatar.UserAvatar) > 1 {
			err = errors.New("[len(findUserAvatar.UserAvatar) >= 1]")
			return h.UserAvatarPut500(err)
		}
		//* когда всё ОК, то мы должны изменить прошлую аватарку на новую
		_, err = h.db.ChangeUserAvatar(
			tx,
			database.UserAvatar{
				Prefix: &params.Body.Prefix,
			},
			database.UserAvatar{
				UserId: &accessToken.UserId,
			},
		)
		if err != nil {
			err = errors.Wrap(err, "[h.db.ChangeUserAvatar(tx, {...}, {...}]")
			return h.UserAvatarPut500(err)
		}
	}

	//* перезаписываем аватарку на новую
	err = h.fileWorker.UserAvatarWrite(accessToken.UserId, data)
	if err != nil {
		err = errors.Wrap(err, "[h.fileWorker.UserAvatarWrite(accessToken.UserId, data)]")
		return h.UserAvatarPut500(err)
	}

	//* коммитим транзакцию
	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.UserAvatarPut500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserAvatarGetOK()
}

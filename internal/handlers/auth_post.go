package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/attempt"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/auth"
	"mgtu/digital-trace/main-backend-service/internal/password_hasher"
	"mgtu/digital-trace/main-backend-service/internal/sequrity/jwt_service.go"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) authPost500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [authPost]")
	h.log.Error(err.Error())
	return attempt.NewAttemptPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) authPost(params auth.AuthPostParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.authPost500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	findUserOut, err := h.db.FindUser(tx, database.User{
		Email: &params.Body.Email,
	})
	if err != nil {
		return h.authPost500(errors.Wrap(err, "[h.db.FindUser()]"))
	}
	if !findUserOut.IsFound {
		return h.authPost500(errors.New("unable to find user [!findUserOut.IsFound]"))
	}
	user := findUserOut.User[0]

	if !password_hasher.CheckPassword(params.Body.Password, *user.Password) {
		return h.authPost500(errors.New("wrong password"))
	}

	createTokenOut, err := h.jwt.CreateSession(jwt_service.CreateSessionIn{
		UserId: *user.Id,
	})
	if err != nil {
		return h.authPost500(errors.Wrap(err, "[h.jwt.CreateSession()]"))
	}

	out := &auth.AuthPostOKBody{
		AccessToken:  createTokenOut.AccessToken,
		RefreshToken: createTokenOut.RefreshToken,
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.authPost500(errors.Wrap(err, "[h.db.CommitTransaction(tx)]"))
	}

	return auth.NewAuthPostOK().WithPayload(out)
}

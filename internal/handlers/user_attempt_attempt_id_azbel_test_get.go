package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/attempt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userAttemptAttemptIdAzbelTestGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAttemptAttemptIdAzbelTestGet]")
	h.log.Error(err.Error())
	return attempt.NewUserAttemptAttemptIDAzbelTestGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAttemptAttemptIdAzbelTestGet(params attempt.UserAttemptAttemptIDAzbelTestGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAttemptAttemptIdAzbelTestGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	findTestAzbelOut, err := h.db.FindTestAzbel(tx, database.TestAzbel{
		UserId:    &accessToken.UserId,
		AttemptId: &params.AttemptID,
	})
	if err != nil {
		return h.userAttemptAttemptIdAzbelTestGet500(errors.Wrap(err, "[h.db.FindTestAzbel()]"))
	}

	if !findTestAzbelOut.IsFound {
		return h.userAttemptAttemptIdAzbelTestGet500(errors.Wrap(err, "[!findTestAzbelOut.IsFound]"))
	}

	out := &models.TestAzbel{}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAttemptAttemptIdAzbelTestGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return attempt.NewUserAttemptAttemptIDAzbelTestGetOK().WithPayload(out)
}

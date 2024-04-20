package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/attempt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) attemptPost500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [attemptPost]")
	h.log.Error(err.Error())
	return attempt.NewAttemptPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) attemptPost(params attempt.AttemptPostParams) middleware.Responder {
	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.attemptPost500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.attemptPost500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}
	openStatus := uint64(1)

	findAttemptOut, err := h.db.FindAttempt(tx, database.Attempt{
		UserId:   &accessToken.UserId,
		TestId:   &params.Body.TestID,
		StatusId: &openStatus,
	})
	if err != nil {
		return h.attemptPost500(errors.Wrap(err, "[h.db.FindAttempt]"))
	}
	if findAttemptOut.IsFound {
		return h.attemptPost500(errors.New("close your last attempt"))
	}

	_, err = h.db.AddAttempt(tx, database.Attempt{
		UserId:   &accessToken.UserId,
		TestId:   &params.Body.TestID,
		StatusId: &openStatus,
	})
	if err != nil {
		return h.attemptPost500(errors.Wrap(err, "[h.db.AddAttempt]"))
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.attemptPost500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return attempt.NewAttemptPostOK()
}

package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/attempt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userAttemptGetError500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userUserIDAttemptGet]")
	h.log.Error(err.Error())
	return attempt.NewUserAttemptGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAttemptGet(params attempt.UserAttemptGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	findAttemptOut, err := h.db.FindAttempt(tx, database.Attempt{
		UserId:   &accessToken.UserId,
		TestId:   params.TestID,
		StatusId: params.StatusID,
	})
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.db.FindAttempt()]"))
	}

	out := []*attempt.UserAttemptGetOKBodyItems0{}

	for _, v := range findAttemptOut.Attempt {
		item := &attempt.UserAttemptGetOKBodyItems0{
			AttemptID: *v.Id,
			StatusID:  *v.StatusId,
			TestID:    *v.TestId,
		}
		out = append(out, item)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAttemptAttemptIdKlimovTestGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return attempt.NewUserAttemptGetOK().WithPayload(out)
}

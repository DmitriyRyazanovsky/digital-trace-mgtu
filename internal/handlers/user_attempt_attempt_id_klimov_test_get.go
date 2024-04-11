package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userAttemptAttemptIdKlimovTestGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userUserIdAttemptAttemptIdKlimovTestGet]")
	h.log.Error(err.Error())
	return user.NewUserAttemptAttemptIDKlimovTestGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAttemptAttemptIdKlimovTestGet(params user.UserAttemptAttemptIDKlimovTestGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAttemptAttemptIdKlimovTestGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	findTestKlimovOut, err := h.db.FindTestKlimov(tx, database.TestKlimov{
		UserId:    &accessToken.UserId,
		AttemptId: &params.AttemptID,
	})
	if err != nil {
		return h.userAttemptAttemptIdKlimovTestGet500(errors.Wrap(err, "[h.db.FindTestKlimov()]"))
	}

	if !findTestKlimovOut.IsFound {
		return h.userAttemptAttemptIdKlimovTestGet500(errors.Wrap(err, "[!findTestKlimovOut.IsFound]"))
	}

	test := findTestKlimovOut.TestKlimov[0]

	out := &models.TestKlimov{
		HumanSign:       test.HumanSign,
		HumanHuman:      test.HumanHuman,
		HumanNature:     test.HumanNature,
		HumanTechnic:    test.HumanTechnic,
		HumanSignSystem: test.HumanSignSystem,
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAttemptAttemptIdKlimovTestGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserAttemptAttemptIDKlimovTestGetOK().WithPayload(out)
}

package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

// * UserUserIDAttemptAttemptIDAnswerGet

func (h *Handler) userAttemptAttemptIdAnswerGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAttemptAttemptIDAnswerGet]")
	h.log.Error(err.Error())
	return user.NewUserAttemptAttemptIDAnswerGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAttemptAttemptIDAnswerGet(params user.UserAttemptAttemptIDAnswerGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAttemptAttemptIdAnswerGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	findAttemptOut, err := h.db.FindAttempt(tx, database.Attempt{
		Id:     &params.AttemptID,
		UserId: &accessToken.UserId,
	})
	if err != nil {
		err = errors.Wrap(err, "[h.db.FindAttempt()]")
		return h.userAttemptAttemptIdAnswerGet500(err)
	}
	if !findAttemptOut.IsFound {
		err = errors.New("this is not your attempt")
		return h.userAttemptAttemptIdAnswerGet500(err)
	}

	findUserAnswerOut, err := h.db.FindUserAnswer(tx, database.UserAnswer{
		AttemptId: &params.AttemptID,
	})
	if err != nil {
		return h.userAttemptAttemptIdAnswerGet500(errors.Wrap(err, "[h.db.FindUserAnswer()]"))
	}

	out := []*user.UserAttemptAttemptIDAnswerGetOKBodyItems0{}

	for _, v := range findUserAnswerOut.UserAnswer {
		userAnswer := []*int64{}
		for _, v := range *v.Answer {
			userAnswer = append(userAnswer, &v)
		}
		item := &user.UserAttemptAttemptIDAnswerGetOKBodyItems0{
			QuestionID: *v.QuestionId,
			UserAnswer: userAnswer,
		}

		out = append(out, item)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAttemptAttemptIdAnswerGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserAttemptAttemptIDAnswerGetOK().WithPayload(out)
}

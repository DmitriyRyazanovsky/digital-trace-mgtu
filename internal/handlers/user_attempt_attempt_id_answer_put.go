package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// * UserUserIDAttemptAttemptIDAnswerPatch

func (h *Handler) userAttemptAttemptIdAnswerPut500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAttemptAttemptIDAnswerPut]")
	h.log.Error(err.Error())
	return user.NewUserAttemptAttemptIDAnswerPutInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAttemptAttemptIDAnswerPut(params user.UserAttemptAttemptIDAnswerPutParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
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

	radio := uint64(1)
	checkbox := uint64(2)

	var findQuestion *database.FindQuestionOut
	if len(params.Body.UserAnswer) > 1 {
		findQuestion, err = h.db.FindQuestion(tx, database.Question{
			Id:           &params.Body.QuestionID,
			ButtonTypeId: &checkbox,
		})
		if err != nil {
			return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.FindQuestion]"))
		}
	} else {
		findQuestion, err = h.db.FindQuestion(tx, database.Question{
			Id:           &params.Body.QuestionID,
			ButtonTypeId: &radio,
		})
		if err != nil {
			return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.FindQuestion]"))
		}
	}
	if !findQuestion.IsFound {
		return h.userAttemptAttemptIdAnswerPut500(errors.New("[!findQuestion.IsFound]"))
	}

	var openStatus uint64 = 1

	findAttemptOut, err = h.db.FindAttempt(tx, database.Attempt{
		Id:       &params.AttemptID,
		StatusId: &openStatus,
	})
	if err != nil {
		return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.FindAttempt]"))
	}
	if !findAttemptOut.IsFound {
		return h.userAttemptAttemptIdAnswerPut500(errors.New("[!findAttemptOut.IsFound]"))
	}
	if len(findAttemptOut.Attempt) != 1 {
		return h.userAttemptAttemptIdAnswerPut500(errors.New("[len(findAttemptOut.Attempt) != 1]"))
	}

	findUserAnswer, err := h.db.FindUserAnswer(tx, database.UserAnswer{
		QuestionId: &params.Body.QuestionID,
		AttemptId:  &params.AttemptID,
	})
	if err != nil {
		return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.FindUserAnswer]"))
	}

	//* Если ответ пользователя не найден, то нужно добавить новый
	if !findUserAnswer.IsFound {
		answer := pq.Int64Array{}
		for _, v := range params.Body.UserAnswer {
			answer = append(answer, *v)
		}
		_, err = h.db.AddUserAnswer(tx, database.UserAnswer{
			AttemptId:  &params.AttemptID,
			QuestionId: &params.Body.QuestionID,
			Answer:     &answer,
		})
		if err != nil {
			return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.AddUserAnswer]"))
		}
		//* Иначе мы должны исправить старый ответ
	} else {
		if len(findUserAnswer.UserAnswer) != 1 {
			return h.userAttemptAttemptIdAnswerPut500(errors.New("[len(findUserAnswer.UserAnswer) != 1]"))
		}

		//* Берём ответ пользователя
		userAnswer := findUserAnswer.UserAnswer[0]

		answer := pq.Int64Array{}
		for _, v := range params.Body.UserAnswer {
			answer = append(answer, *v)
		}
		_, err = h.db.ChangeUserAnswer(
			tx,
			database.UserAnswer{
				Answer: &answer,
			},
			database.UserAnswer{
				Id: userAnswer.Id,
			},
		)
		if err != nil {
			return h.userAttemptAttemptIdAnswerPut500(errors.New("h.db.ChangeUserAnswer()]"))
		}
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAttemptAttemptIdAnswerPut500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserAttemptAttemptIDAnswerPutOK()
}

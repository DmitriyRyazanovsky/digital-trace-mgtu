package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/test"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) testTestIDQuestionGet500(err error) middleware.Responder {
	err = errors.Wrap(err, "handler error: [testGet]")
	h.log.Error(err.Error())
	return test.NewTestGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) testTestIDQuestionGet(params test.TestTestIDQuestionGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.testTestIDQuestionGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	sql, err := h.db.FindQuestion(tx, database.Question{
		TestId: &params.TestID,
	})
	if err != nil {
		return h.testTestIDQuestionGet500(errors.Wrap(err, "[h.db.GetTest]"))
	}

	out := []*test.TestTestIDQuestionGetOKBodyItems0{}

	for _, v := range sql.Question {
		answers := []*test.TestTestIDQuestionGetOKBodyItems0AnswersItems0{}

		for _, v := range *v.Answer {
			elem := &test.TestTestIDQuestionGetOKBodyItems0AnswersItems0{
				Content: v,
			}
			answers = append(answers, elem)
		}

		item := &test.TestTestIDQuestionGetOKBodyItems0{
			QuestionID: *v.Id,
			Content:    *v.Content,
			ButtonType: *v.ButtonTypeId,
			Answers:    answers,
		}

		out = append(out, item)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.testTestIDQuestionGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return test.NewTestTestIDQuestionGetOK().WithPayload(out)
}

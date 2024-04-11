package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/attempt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

const (
	testKlimov = "test_klimov"
	testAzbel  = "test_azbel"
)

type Answer struct {
	QuestionNumber uint64
	Answer         int64
}

type AzbelCalculate struct {
	HumanSign       []Answer
	HumanHuman      []Answer
	HumanNature     []Answer
	HumanTechnic    []Answer
	HumanSignSystem []Answer
}

var (
	azbelCalc = AzbelCalculate{
		HumanSign: []Answer{
			Answer{
				QuestionNumber: 0,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 2,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 5,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 9,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 10,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 12,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 15,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 19,
				Answer:         1,
			},
		},
		HumanHuman: []Answer{
			Answer{
				QuestionNumber: 1,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 3,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 1,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 5,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 7,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 11,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 13,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 15,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 17,
				Answer:         1,
			},
		},
		HumanNature: []Answer{
			Answer{
				QuestionNumber: 2,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 4,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 6,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 7,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 12,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 14,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 16,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 17,
				Answer:         2,
			},
		},
		HumanTechnic: []Answer{
			Answer{
				QuestionNumber: 0,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 3,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 6,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 8,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 10,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 13,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 16,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 18,
				Answer:         1,
			},
		},
		HumanSignSystem: []Answer{
			Answer{
				QuestionNumber: 1,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 4,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 8,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 9,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 11,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 14,
				Answer:         1,
			},
			Answer{
				QuestionNumber: 18,
				Answer:         2,
			},
			Answer{
				QuestionNumber: 19,
				Answer:         2,
			},
		},
	}
)

func (h *Handler) attemptAttemptIdClosePatchError(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [attemptAttemptIdClosePatch]")
	h.log.Error(err.Error())
	return attempt.NewAttemptAttemptIDClosePatchInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) attemptAttemptIdClosePatch(params attempt.AttemptAttemptIDClosePatchParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	findAttemptOut, err := h.db.FindAttempt(tx, database.Attempt{
		Id: &params.AttemptID,
	})
	if err != nil {
		return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindAttempt()]"))
	}
	if !findAttemptOut.IsFound {
		return h.attemptAttemptIdClosePatchError(errors.New("unable to find attempt [!findAttemptOut.IsFound]"))
	}
	if len(findAttemptOut.Attempt) != 1 {
		return h.attemptAttemptIdClosePatchError(errors.New("[len(findAttemptOut.Attempt) != 1]"))
	}

	userAttempt := findAttemptOut.Attempt[0]

	findTestOut, err := h.db.FindTest(tx, database.Test{
		Id: userAttempt.TestId,
	})
	if err != nil {
		return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindTest()]"))
	}
	if !findTestOut.IsFound {
		return h.attemptAttemptIdClosePatchError(errors.New("unable to find test [!findTestOut.IsFound]"))
	}
	if len(findTestOut.Test) != 1 {
		return h.attemptAttemptIdClosePatchError(errors.New("[len(findTestOut.Test) != 1]"))
	}

	test := findTestOut.Test[0]

	switch *test.PgName {
	case testKlimov:
		humanSign := uint64(0)
		humanHuman := uint64(0)
		humanNature := uint64(0)
		humanTechnic := uint64(0)
		humanSignSystem := uint64(0)

		for _, v := range azbelCalc.HumanSign {
			findQuestionOut, err := h.db.FindQuestion(tx, database.Question{
				TestId: test.Id,
				Number: &v.QuestionNumber,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindQuestion()]"))
			}
			if !findQuestionOut.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findQuestionOut.IsFound]"))
			}
			if len(findQuestionOut.Question) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findQuestionOut.Question) != 1]"))
			}

			question := findQuestionOut.Question[0]

			findUserAnswer, err := h.db.FindUserAnswer(tx, database.UserAnswer{
				AttemptId:  &params.AttemptID,
				QuestionId: question.Id,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindUserAnswer()]"))
			}
			if !findUserAnswer.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findUserAnswer.IsFound]"))
			}
			if len(findUserAnswer.UserAnswer) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findUserAnswer.UserAnswer) != 1]"))
			}

			answer := findUserAnswer.UserAnswer[0]

			if (*answer.Answer)[0] == v.Answer {
				humanSign++
			}
		}

		for _, v := range azbelCalc.HumanHuman {
			findQuestionOut, err := h.db.FindQuestion(tx, database.Question{
				TestId: test.Id,
				Number: &v.QuestionNumber,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindQuestion()]"))
			}
			if !findQuestionOut.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findQuestionOut.IsFound]"))
			}
			if len(findQuestionOut.Question) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findQuestionOut.Question) != 1]"))
			}

			question := findQuestionOut.Question[0]

			findUserAnswer, err := h.db.FindUserAnswer(tx, database.UserAnswer{
				AttemptId:  &params.AttemptID,
				QuestionId: question.Id,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindUserAnswer()]"))
			}
			if !findUserAnswer.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findUserAnswer.IsFound]"))
			}
			if len(findUserAnswer.UserAnswer) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findUserAnswer.UserAnswer) != 1]"))
			}

			answer := findUserAnswer.UserAnswer[0]

			if (*answer.Answer)[0] == v.Answer {
				humanHuman++
			}
		}

		for _, v := range azbelCalc.HumanNature {
			findQuestionOut, err := h.db.FindQuestion(tx, database.Question{
				TestId: test.Id,
				Number: &v.QuestionNumber,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindQuestion()]"))
			}
			if !findQuestionOut.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findQuestionOut.IsFound]"))
			}
			if len(findQuestionOut.Question) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findQuestionOut.Question) != 1]"))
			}

			question := findQuestionOut.Question[0]

			findUserAnswer, err := h.db.FindUserAnswer(tx, database.UserAnswer{
				AttemptId:  &params.AttemptID,
				QuestionId: question.Id,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindUserAnswer()]"))
			}
			if !findUserAnswer.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findUserAnswer.IsFound]"))
			}
			if len(findUserAnswer.UserAnswer) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findUserAnswer.UserAnswer) != 1]"))
			}

			answer := findUserAnswer.UserAnswer[0]

			if (*answer.Answer)[0] == v.Answer {
				humanNature++
			}
		}

		for _, v := range azbelCalc.HumanTechnic {
			findQuestionOut, err := h.db.FindQuestion(tx, database.Question{
				TestId: test.Id,
				Number: &v.QuestionNumber,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindQuestion()]"))
			}
			if !findQuestionOut.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findQuestionOut.IsFound]"))
			}
			if len(findQuestionOut.Question) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findQuestionOut.Question) != 1]"))
			}

			question := findQuestionOut.Question[0]

			findUserAnswer, err := h.db.FindUserAnswer(tx, database.UserAnswer{
				AttemptId:  &params.AttemptID,
				QuestionId: question.Id,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindUserAnswer()]"))
			}
			if !findUserAnswer.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findUserAnswer.IsFound]"))
			}
			if len(findUserAnswer.UserAnswer) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findUserAnswer.UserAnswer) != 1]"))
			}

			answer := findUserAnswer.UserAnswer[0]

			if (*answer.Answer)[0] == v.Answer {
				humanTechnic++
			}
		}

		for _, v := range azbelCalc.HumanSignSystem {
			findQuestionOut, err := h.db.FindQuestion(tx, database.Question{
				TestId: test.Id,
				Number: &v.QuestionNumber,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindQuestion()]"))
			}
			if !findQuestionOut.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findQuestionOut.IsFound]"))
			}
			if len(findQuestionOut.Question) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findQuestionOut.Question) != 1]"))
			}

			question := findQuestionOut.Question[0]

			findUserAnswer, err := h.db.FindUserAnswer(tx, database.UserAnswer{
				AttemptId:  &params.AttemptID,
				QuestionId: question.Id,
			})
			if err != nil {
				return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.FindUserAnswer()]"))
			}
			if !findUserAnswer.IsFound {
				return h.attemptAttemptIdClosePatchError(errors.New("[!findUserAnswer.IsFound]"))
			}
			if len(findUserAnswer.UserAnswer) != 1 {
				return h.attemptAttemptIdClosePatchError(errors.New("[len(findUserAnswer.UserAnswer) != 1]"))
			}

			answer := findUserAnswer.UserAnswer[0]

			if (*answer.Answer)[0] == v.Answer {
				humanSignSystem++
			}
		}

		_, err = h.db.AddTestKlimov(tx, database.TestKlimov{
			UserId:          userAttempt.UserId,
			AttemptId:       &params.AttemptID,
			HumanSign:       &humanSign,
			HumanHuman:      &humanHuman,
			HumanNature:     &humanNature,
			HumanTechnic:    &humanTechnic,
			HumanSignSystem: &humanSignSystem,
		})
		if err != nil {
			return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.AddTestKlimov()]"))
		}

		var closedStatus uint64 = 2

		changeAttemptOut, err := h.db.ChangeAttempt(tx,
			database.Attempt{
				StatusId: &closedStatus,
			},
			database.Attempt{
				Id: &params.AttemptID,
			},
		)
		if err != nil {
			return h.attemptAttemptIdClosePatchError(errors.Wrap(err, "[h.db.ChangeAttempt()]"))
		}
		if !changeAttemptOut.IsFound {
			return h.attemptAttemptIdClosePatchError(errors.New("[!changeAttemptOut.IsFound]"))
		}
	default:
		return h.attemptAttemptIdClosePatchError(errors.New("unable to find test"))
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userPost500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return attempt.NewAttemptPostOK()
}

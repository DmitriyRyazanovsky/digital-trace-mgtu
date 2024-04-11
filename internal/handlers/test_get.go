package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/test"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) testGet500(err error) middleware.Responder {
	err = errors.Wrap(err, "handler error: [testGet]")
	h.log.Error(err.Error())
	return test.NewTestGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) testGet(params test.TestGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.testGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	sql, err := h.db.FindTest(tx, database.Test{})
	if err != nil {
		return h.testGet500(errors.Wrap(err, "[h.db.GetTest]"))
	}

	out := []*test.TestGetOKBodyItems0{}

	for _, v := range sql.Test {
		item := &test.TestGetOKBodyItems0{
			Content:     *v.Content,
			Description: *v.Description,
			Name:        *v.Name,
			TestID:      *v.Id,
		}

		out = append(out, item)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.testGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return test.NewTestGetOK().WithPayload(out)
}

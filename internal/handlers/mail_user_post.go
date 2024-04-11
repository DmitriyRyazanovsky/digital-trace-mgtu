package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/mail"
	"mgtu/digital-trace/main-backend-service/internal/mail_service"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) mailPost500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [mailUserPost]")
	h.log.Error(err.Error())
	return mail.NewMailUserPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func generateRandomToken(len int64) (string, error) {
	buffer := make([]byte, len)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", errors.Wrap(err, "unable to make rand buffer: [rand.Read]")
	}
	randomString := base64.URLEncoding.EncodeToString(buffer)
	return randomString[:len], nil
}

func (h *Handler) mailUserPost(params mail.MailUserPostParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.mailPost500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	token, err := generateRandomToken(64)
	if err != nil {
		return h.mailPost500(errors.Wrap(err, "[generateRandomToken]"))
	}

	err = h.mail.SendMail(&mail_service.SendMailIn{
		To:      *params.Body.Email,
		Subject: "token",
		Body:    fmt.Sprintf("http://79.174.95.104:5500/js/mail/?accept_token=%s", token),
	})
	if err != nil {
		return h.mailPost500(errors.Wrap(err, "[h.mail.SendMail]"))
	}

	_, err = h.db.AddMailUserVerefication(tx, database.MailUserVerefication{
		Token:    &token,
		Login:    params.Body.Login,
		Name:     params.Body.Name,
		Surname:  params.Body.Surname,
		Email:    params.Body.Email,
		Password: params.Body.Password,
	})
	if err != nil {
		return h.mailPost500(errors.Wrap(err, "[h.db.AddMailUserVerification]"))
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.mailPost500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	out := mail.MailUserPostOKBody{
		Success: "success",
	}

	return mail.NewMailUserPostOK().WithPayload(&out)
}

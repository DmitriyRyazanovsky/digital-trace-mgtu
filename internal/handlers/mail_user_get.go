package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/mail"
	"mgtu/digital-trace/main-backend-service/internal/password_hasher"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) mailUserGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [mailUserGet]")
	h.log.Error(err.Error())
	return mail.NewMailUserGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) mailUserGet(params mail.MailUserGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.mailUserGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	user := "user"

	findMailUserVereficationOut, err := h.db.FindMailUserVerefication(tx, database.MailUserVerefication{
		Token: &params.AcceptToken,
	})
	if err != nil {
		return h.mailUserGet500(errors.Wrap(err, "[h.db.GetMailUserVerificationByToken]"))
	}

	mailVer := findMailUserVereficationOut.MailUserVerefication[0]

	findRoleOut, err := h.db.FindRole(tx, database.Role{
		Name: &user,
	})
	if err != nil {
		return h.mailUserGet500(errors.Wrap(err, "[h.db.GetRoleIdByName]"))
	}

	role := findRoleOut.Role[0]

	hashPass, err := password_hasher.HashPassword(*mailVer.Password)
	if err != nil {
		err = errors.Wrap(err, "[password_hasher.HashPassword(mailVer.Password)]")
		return h.mailUserGet500(err)
	}

	adduserOut, err := h.db.AddUser(tx, database.User{
		RoleId:   role.Id,
		Email:    mailVer.Email,
		Login:    mailVer.Login,
		Name:     mailVer.Name,
		Surname:  mailVer.Surname,
		Password: &hashPass,
	})
	if err != nil {
		return h.mailUserGet500(errors.Wrap(err, "h.db.AddUser]"))
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.mailUserGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return mail.NewMailUserGetOK().WithPayload(&mail.MailUserGetOKBody{
		UserID: *adduserOut.User.Id,
	})
}

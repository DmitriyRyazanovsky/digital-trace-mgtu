package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/session"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) sessionPatch500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [sessionPatch]")
	h.log.Error(err.Error())
	return session.NewSessionPatchInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) sessionPatch(params session.SessionPatchParams) middleware.Responder {
	updatedSession, err := h.jwt.ReloadSession(params.Body.RefreshToken)
	if err != nil {
		err = errors.Wrap(err, "[h.jwt.ReloadSession(params.Body.RefreshToken)]")
		return h.sessionPatch500(err)
	}

	out := &session.SessionPatchOKBody{
		AccessToken:  updatedSession.AccessToken,
		RefreshToken: updatedSession.RefreshToken,
	}

	return session.NewSessionPatchOK().WithPayload(out)
}

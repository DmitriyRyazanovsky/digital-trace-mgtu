package handlers

import (
	fileworker "mgtu/digital-trace/main-backend-service/internal/file_worker"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/achievement"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userAchievementAchiveIDImageGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAchievementAchiveIDImageGet]")
	h.log.Error(err.Error())
	return achievement.NewUserAchievementAchiveIDImageGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAchievementAchiveIDImageGet(params achievement.UserAchievementAchiveIDImageGetParams) middleware.Responder {
	claims, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAchievementAchiveIDImageGet500(errors.Wrap(err, "[ValidateAccessToken(params.Authorization)]"))
	}

	bytes, _, err := h.fileWorker.GetAchievement(fileworker.GetAchievementIn{
		UserId:        claims.UserId,
		AchievementId: params.AchiveID,
	})
	if err != nil {
		return h.userAchievementAchiveIDImageGet500(errors.Wrap(err, "[h.fileWorker.GetAchievement()]"))
	}
	return achievement.NewUserAchievementAchiveIDImageGetOK().
		WithPayload(string(bytes)).
		WithContentDisposition(`attachment; filename="image.png"`)
}

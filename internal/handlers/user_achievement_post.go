package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/features"
	fileworker "mgtu/digital-trace/main-backend-service/internal/file_worker"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/achievement"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userAchievementPost500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAchievementPost]")
	h.log.Error(err.Error())
	return achievement.NewUserAchievementPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAchievementPost(params achievement.UserAchievementPostParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAchievementPost500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	addAchievementOut, err := h.db.AddAchievement(tx, database.Achievement{
		UserId:           &accessToken.UserId,
		AchievementTypes: features.ConvertInToPqInt64Array(params.Body.AchiveTypes),
	})
	if err != nil {
		return h.userAchievementPost500(errors.Wrap(err, "[h.db.AddAchievement()]"))
	}

	err = h.fileWorker.AddAchievement(fileworker.AddAchievementIn{
		AchievementId: *addAchievementOut.Achievement.Id,
		UserId:        accessToken.UserId,
		FileContent:   []byte(params.Body.Image),
	})
	if err != nil {
		return h.userAchievementPost500(errors.Wrapf(err, "[h.fileWorker.AddAchievement()]"))
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAchievementPost500(errors.Wrap(err, "[h.db.CommitTransaction(tx)]"))
	}

	return achievement.NewUserAchievementPostOK()
}

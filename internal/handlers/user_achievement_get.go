package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) userAchievementGet500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [userAchievementGet]")
	h.log.Error(err.Error())
	return user.NewUserAchievementGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) userAchievementGet(params user.UserAchievementGetParams) middleware.Responder {
	tx, err := h.db.OpenTransaction()
	if err != nil {
		return h.userAchievementGet500(errors.Wrap(err, "[h.db.OpenTransaction()]"))
	}

	accessToken, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		return h.userAttemptGetError500(errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]"))
	}

	getAchievementOut, err := h.db.FindAchievement(tx, database.Achievement{
		UserId: &accessToken.UserId,
	})
	if err != nil {
		return h.userAchievementGet500(errors.Wrap(err, "[h.db.GetAchievement()]"))
	}

	out := []*user.UserAchievementGetOKBodyItems0{}

	for _, v := range getAchievementOut.Achievement {
		achiveTypeList := []*user.UserAchievementGetOKBodyItems0AchievementTypesItems0{}

		for _, v := range *v.AchievementTypes {
			item := &user.UserAchievementGetOKBodyItems0AchievementTypesItems0{
				AchievementTypeID: &v,
			}

			achiveTypeList = append(achiveTypeList, item)
		}

		elem := &user.UserAchievementGetOKBodyItems0{
			AchievementID:    *v.Id,
			AchievementTypes: achiveTypeList,
		}
		out = append(out, elem)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		return h.userAchievementGet500(errors.Wrap(err, "[h.db.CommitTransaction()]"))
	}

	return user.NewUserAchievementGetOK().WithPayload(out)
}

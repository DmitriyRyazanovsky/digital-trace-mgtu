package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

func (h *Handler) userProfilePatch500(err error) middleware.Responder {
	err = errors.Wrap(err, "handler error: [userProfileGet]")
	h.log.Error(err.Error())
	return user.NewUserPostInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

type SetUser struct {
	Email    *strfmt.Email
	ID       *uint64
	Login    *string
	Name     *string
	Password *string
	RoleID   *uint64
	Surname  *string
}

func (h *Handler) userProfilePatch(params user.UserProfilePatchParams) middleware.Responder {
	setuser := database.User{
		Email:    (*string)(params.Email),
		Id:       params.ID,
		Login:    params.Login,
		Name:     params.Name,
		Password: params.Password,
		RoleId:   params.RoleID,
		Surname:  params.Surname,
	}

	claims, err := h.jwt.ValidateAccessToken(params.Authorization)
	if err != nil {
		err = errors.Wrap(err, "[h.jwt.ValidateAccessToken(params.Authorization)]")
		return h.userProfilePatch500(err)
	}

	tx, err := h.db.OpenTransaction()
	if err != nil {
		err = errors.Wrap(err, "[h.db.OpenTransaction()]")
		return h.userProfilePatch500(err)
	}

	findUserOut, err := h.db.ChangeUser(tx, setuser, database.User{
		Id: &claims.UserId,
	})
	if err != nil {
		err = errors.Wrap(err, "[h.db.FindUser()]")
		return h.userProfilePatch500(err)
	}
	if !findUserOut.IsFound {
		err = errors.New("unable find user by id")
		return h.userProfilePatch500(err)
	}
	if len(findUserOut.User) != 1 {
		err = errors.New("len(findUserOut.User) != 1")
		return h.userProfilePatch500(err)
	}

	err = h.db.CommitTransaction(tx)
	if err != nil {
		err = errors.Wrap(err, "[h.db.CommitTransaction(tx)]")
		return h.userProfilePatch500(err)
	}

	return user.NewUserProfilePatchOK()
}

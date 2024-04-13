package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/test"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func (h *Handler) getAllUsersGet(params test.GetAllUsersGetParams) middleware.Responder {
	tx, _ := h.db.OpenTransaction()
	FindUserOut, _ := h.db.FindUser(tx, database.User{
		Id: params.ID,
	})
	out := []*test.GetAllUsersGetOKBodyItems0{}
	for _, v := range FindUserOut.User {
		item := &test.GetAllUsersGetOKBodyItems0{
			ID:        *v.Id,
			CreatedAt: *v.CreatedAt,
			Email:     strfmt.Email(*v.Email),
			Login:     *v.Login,
			Name:      *v.Name,
			Surname:   *v.Surname,
			Password:  *v.Password,
			RoleID:    *v.RoleId,
			UpdatedAt: *v.UpdatedAt,
		}
		out = append(out, item)
	}
	h.db.CommitTransaction(tx)
	return test.NewGetAllUsersGetOK().WithPayload(out)
}

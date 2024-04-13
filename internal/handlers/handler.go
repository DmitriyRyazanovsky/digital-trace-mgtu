package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/features/logging"
	fileworker "mgtu/digital-trace/main-backend-service/internal/file_worker"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/all_users"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/attempt"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/auth"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/logs"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/mail"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/session"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/test"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/user"
	"mgtu/digital-trace/main-backend-service/internal/mail_service"
	"mgtu/digital-trace/main-backend-service/internal/sequrity/jwt_service.go"
)

type Handler struct {
	db         *database.Database
	fileWorker *fileworker.FileWorker
	mail       *mail_service.Mail
	jwt        *jwt_service.JWT
	log        logging.Logger
}

func NewHandler(db *database.Database, fileWorker *fileworker.FileWorker, mail *mail_service.Mail, jwt *jwt_service.JWT, log logging.Logger) *Handler {
	return &Handler{
		db:         db,
		fileWorker: fileWorker,
		mail:       mail,
		jwt:        jwt,
		log:        log,
	}
}

func (h *Handler) Register(api *operations.BackendServiceAPI) {
	api.AttemptAttemptPostHandler = attempt.AttemptPostHandlerFunc(h.attemptPost)
	api.AuthAuthPostHandler = auth.AuthPostHandlerFunc(h.authPost)
	api.SessionSessionPatchHandler = session.SessionPatchHandlerFunc(h.sessionPatch)
	api.LogsLogsGetHandler = logs.LogsGetHandlerFunc(h.logsGet)
	api.MailMailUserGetHandler = mail.MailUserGetHandlerFunc(h.mailUserGet)
	api.MailMailUserPostHandler = mail.MailUserPostHandlerFunc(h.mailUserPost)
	api.TestTestGetHandler = test.TestGetHandlerFunc(h.testGet)
	api.TestTestTestIDQuestionGetHandler = test.TestTestIDQuestionGetHandlerFunc(h.testTestIDQuestionGet)
	api.AllUsersUserGetHandler = all_users.UserGetHandlerFunc(h.userGet)
	api.UserUserPostHandler = user.UserPostHandlerFunc(h.userPost)
	api.UserUserAchievementPostHandler = user.UserAchievementPostHandlerFunc(h.userAchievementPost)
	api.UserUserAchievementAchiveIDImageGetHandler = user.UserAchievementAchiveIDImageGetHandlerFunc(h.userAchievementAchiveIDImageGet)
	api.UserUserAchievementGetHandler = user.UserAchievementGetHandlerFunc(h.userAchievementGet)
	api.UserUserAttemptAttemptIDAnswerGetHandler = user.UserAttemptAttemptIDAnswerGetHandlerFunc(h.userAttemptAttemptIDAnswerGet)
	api.UserUserAttemptAttemptIDAnswerPutHandler = user.UserAttemptAttemptIDAnswerPutHandlerFunc(h.userAttemptAttemptIDAnswerPut)
	api.UserUserAttemptAttemptIDAzbelTestGetHandler = user.UserAttemptAttemptIDAzbelTestGetHandlerFunc(h.userAttemptAttemptIdAzbelTestGet)
	api.UserUserAttemptAttemptIDKlimovTestGetHandler = user.UserAttemptAttemptIDKlimovTestGetHandlerFunc(h.userAttemptAttemptIdKlimovTestGet)
	api.AttemptAttemptAttemptIDClosePatchHandler = attempt.AttemptAttemptIDClosePatchHandlerFunc(h.attemptAttemptIdClosePatch)
	api.UserUserAttemptGetHandler = user.UserAttemptGetHandlerFunc(h.userAttemptGet)
	api.UserUserAvatarPutHandler = user.UserAvatarPutHandlerFunc(h.userAvatarPut)
	api.UserUserAvatarGetHandler = user.UserAvatarGetHandlerFunc(h.userAvatarGet)
	api.TestGetAllUsersGetHandler = test.GetAllUsersGetHandlerFunc(h.getAllUsersGet)
	api.UserUserProfileGetHandler = user.UserProfileGetHandlerFunc(h.userProfileGet)
}

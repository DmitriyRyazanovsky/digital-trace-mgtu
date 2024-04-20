package handlers

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/features/logging"
	fileworker "mgtu/digital-trace/main-backend-service/internal/file_worker"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/achievement"
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
	//*: Запросы ориентированные на попытки прохождения теста
	api.AttemptAttemptPostHandler = attempt.AttemptPostHandlerFunc(h.attemptPost)
	api.AttemptUserAttemptAttemptIDAnswerGetHandler = attempt.UserAttemptAttemptIDAnswerGetHandlerFunc(h.userAttemptAttemptIDAnswerGet)
	api.AttemptUserAttemptAttemptIDAnswerPutHandler = attempt.UserAttemptAttemptIDAnswerPutHandlerFunc(h.userAttemptAttemptIDAnswerPut)
	api.AttemptUserAttemptAttemptIDAzbelTestGetHandler = attempt.UserAttemptAttemptIDAzbelTestGetHandlerFunc(h.userAttemptAttemptIdAzbelTestGet)
	api.AttemptUserAttemptAttemptIDKlimovTestGetHandler = attempt.UserAttemptAttemptIDKlimovTestGetHandlerFunc(h.userAttemptAttemptIdKlimovTestGet)
	api.AttemptAttemptAttemptIDClosePatchHandler = attempt.AttemptAttemptIDClosePatchHandlerFunc(h.attemptAttemptIdClosePatch)
	api.AttemptUserAttemptGetHandler = attempt.UserAttemptGetHandlerFunc(h.userAttemptGet)

	//*: Запросы связанные с авторизацией
	api.AuthAuthPostHandler = auth.AuthPostHandlerFunc(h.authPost)

	//*: Запросы завязанные на завершении сессий
	api.SessionSessionPatchHandler = session.SessionPatchHandlerFunc(h.sessionPatch)

	//*: Запросы созданные для логирования информации
	//TODO: Сделать только доступ для root
	api.LogsLogsGetHandler = logs.LogsGetHandlerFunc(h.logsGet)

	//*: Запросы завязанные на mail регистрации
	api.MailMailUserGetHandler = mail.MailUserGetHandlerFunc(h.mailUserGet)
	api.MailMailUserPostHandler = mail.MailUserPostHandlerFunc(h.mailUserPost)

	//*: Запросы завязанные на тестировании
	api.TestTestGetHandler = test.TestGetHandlerFunc(h.testGet)
	api.TestTestTestIDQuestionGetHandler = test.TestTestIDQuestionGetHandlerFunc(h.testTestIDQuestionGet)

	//*: Запрос на получение всех пользователей
	api.AllUsersUserGetHandler = all_users.UserGetHandlerFunc(h.userGet)

	//*: Запросы завязанные на достижениях
	api.AchievementUserAchievementPostHandler = achievement.UserAchievementPostHandlerFunc(h.userAchievementPost)
	api.AchievementUserAchievementAchiveIDImageGetHandler = achievement.UserAchievementAchiveIDImageGetHandlerFunc(h.userAchievementAchiveIDImageGet)
	api.AchievementUserAchievementGetHandler = achievement.UserAchievementGetHandlerFunc(h.userAchievementGet)

	//*: Запросы на пользователя
	api.UserUserAvatarPutHandler = user.UserAvatarPutHandlerFunc(h.userAvatarPut)
	api.UserUserAvatarGetHandler = user.UserAvatarGetHandlerFunc(h.userAvatarGet)
	api.UserUserProfileGetHandler = user.UserProfileGetHandlerFunc(h.userProfileGet)

	//TODO: Сделать только admin и root доступ
	api.UserUserPostHandler = user.UserPostHandlerFunc(h.userPost)
}

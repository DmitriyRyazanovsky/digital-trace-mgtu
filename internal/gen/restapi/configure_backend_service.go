// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

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
)

//go:generate swagger generate server --target ../../gen --name BackendService --spec ../../../api/swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.BackendServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BackendServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()

	if api.AchievementAchievementTypeGetHandler == nil {
		api.AchievementAchievementTypeGetHandler = achievement.AchievementTypeGetHandlerFunc(func(params achievement.AchievementTypeGetParams) middleware.Responder {
			return middleware.NotImplemented("operation achievement.AchievementTypeGet has not yet been implemented")
		})
	}
	if api.AttemptAttemptAttemptIDClosePatchHandler == nil {
		api.AttemptAttemptAttemptIDClosePatchHandler = attempt.AttemptAttemptIDClosePatchHandlerFunc(func(params attempt.AttemptAttemptIDClosePatchParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.AttemptAttemptIDClosePatch has not yet been implemented")
		})
	}
	if api.AttemptAttemptPostHandler == nil {
		api.AttemptAttemptPostHandler = attempt.AttemptPostHandlerFunc(func(params attempt.AttemptPostParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.AttemptPost has not yet been implemented")
		})
	}
	if api.AuthAuthPostHandler == nil {
		api.AuthAuthPostHandler = auth.AuthPostHandlerFunc(func(params auth.AuthPostParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthPost has not yet been implemented")
		})
	}
	if api.LogsLogsGetHandler == nil {
		api.LogsLogsGetHandler = logs.LogsGetHandlerFunc(func(params logs.LogsGetParams) middleware.Responder {
			return middleware.NotImplemented("operation logs.LogsGet has not yet been implemented")
		})
	}
	if api.MailMailUserGetHandler == nil {
		api.MailMailUserGetHandler = mail.MailUserGetHandlerFunc(func(params mail.MailUserGetParams) middleware.Responder {
			return middleware.NotImplemented("operation mail.MailUserGet has not yet been implemented")
		})
	}
	if api.MailMailUserPostHandler == nil {
		api.MailMailUserPostHandler = mail.MailUserPostHandlerFunc(func(params mail.MailUserPostParams) middleware.Responder {
			return middleware.NotImplemented("operation mail.MailUserPost has not yet been implemented")
		})
	}
	if api.SessionSessionDeleteHandler == nil {
		api.SessionSessionDeleteHandler = session.SessionDeleteHandlerFunc(func(params session.SessionDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation session.SessionDelete has not yet been implemented")
		})
	}
	if api.SessionSessionPatchHandler == nil {
		api.SessionSessionPatchHandler = session.SessionPatchHandlerFunc(func(params session.SessionPatchParams) middleware.Responder {
			return middleware.NotImplemented("operation session.SessionPatch has not yet been implemented")
		})
	}
	if api.TestTestGetHandler == nil {
		api.TestTestGetHandler = test.TestGetHandlerFunc(func(params test.TestGetParams) middleware.Responder {
			return middleware.NotImplemented("operation test.TestGet has not yet been implemented")
		})
	}
	if api.TestTestTestIDQuestionGetHandler == nil {
		api.TestTestTestIDQuestionGetHandler = test.TestTestIDQuestionGetHandlerFunc(func(params test.TestTestIDQuestionGetParams) middleware.Responder {
			return middleware.NotImplemented("operation test.TestTestIDQuestionGet has not yet been implemented")
		})
	}
	if api.AttemptUserAttemptAttemptIDAnswerGetHandler == nil {
		api.AttemptUserAttemptAttemptIDAnswerGetHandler = attempt.UserAttemptAttemptIDAnswerGetHandlerFunc(func(params attempt.UserAttemptAttemptIDAnswerGetParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.UserAttemptAttemptIDAnswerGet has not yet been implemented")
		})
	}
	if api.AttemptUserAttemptAttemptIDAnswerPutHandler == nil {
		api.AttemptUserAttemptAttemptIDAnswerPutHandler = attempt.UserAttemptAttemptIDAnswerPutHandlerFunc(func(params attempt.UserAttemptAttemptIDAnswerPutParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.UserAttemptAttemptIDAnswerPut has not yet been implemented")
		})
	}
	if api.AllUsersUserGetHandler == nil {
		api.AllUsersUserGetHandler = all_users.UserGetHandlerFunc(func(params all_users.UserGetParams) middleware.Responder {
			return middleware.NotImplemented("operation all_users.UserGet has not yet been implemented")
		})
	}
	if api.UserUserPostHandler == nil {
		api.UserUserPostHandler = user.UserPostHandlerFunc(func(params user.UserPostParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserPost has not yet been implemented")
		})
	}
	if api.AchievementUserAchievementAchiveIDImageGetHandler == nil {
		api.AchievementUserAchievementAchiveIDImageGetHandler = achievement.UserAchievementAchiveIDImageGetHandlerFunc(func(params achievement.UserAchievementAchiveIDImageGetParams) middleware.Responder {
			return middleware.NotImplemented("operation achievement.UserAchievementAchiveIDImageGet has not yet been implemented")
		})
	}
	if api.AchievementUserAchievementGetHandler == nil {
		api.AchievementUserAchievementGetHandler = achievement.UserAchievementGetHandlerFunc(func(params achievement.UserAchievementGetParams) middleware.Responder {
			return middleware.NotImplemented("operation achievement.UserAchievementGet has not yet been implemented")
		})
	}
	if api.AchievementUserAchievementPostHandler == nil {
		api.AchievementUserAchievementPostHandler = achievement.UserAchievementPostHandlerFunc(func(params achievement.UserAchievementPostParams) middleware.Responder {
			return middleware.NotImplemented("operation achievement.UserAchievementPost has not yet been implemented")
		})
	}
	if api.AttemptUserAttemptAttemptIDAzbelTestGetHandler == nil {
		api.AttemptUserAttemptAttemptIDAzbelTestGetHandler = attempt.UserAttemptAttemptIDAzbelTestGetHandlerFunc(func(params attempt.UserAttemptAttemptIDAzbelTestGetParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.UserAttemptAttemptIDAzbelTestGet has not yet been implemented")
		})
	}
	if api.AttemptUserAttemptAttemptIDKlimovTestGetHandler == nil {
		api.AttemptUserAttemptAttemptIDKlimovTestGetHandler = attempt.UserAttemptAttemptIDKlimovTestGetHandlerFunc(func(params attempt.UserAttemptAttemptIDKlimovTestGetParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.UserAttemptAttemptIDKlimovTestGet has not yet been implemented")
		})
	}
	if api.AttemptUserAttemptGetHandler == nil {
		api.AttemptUserAttemptGetHandler = attempt.UserAttemptGetHandlerFunc(func(params attempt.UserAttemptGetParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.UserAttemptGet has not yet been implemented")
		})
	}
	if api.UserUserAvatarGetHandler == nil {
		api.UserUserAvatarGetHandler = user.UserAvatarGetHandlerFunc(func(params user.UserAvatarGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAvatarGet has not yet been implemented")
		})
	}
	if api.UserUserAvatarPutHandler == nil {
		api.UserUserAvatarPutHandler = user.UserAvatarPutHandlerFunc(func(params user.UserAvatarPutParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAvatarPut has not yet been implemented")
		})
	}
	if api.UserUserProfileGetHandler == nil {
		api.UserUserProfileGetHandler = user.UserProfileGetHandlerFunc(func(params user.UserProfileGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserProfileGet has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

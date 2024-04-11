// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

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

// NewBackendServiceAPI creates a new BackendService instance
func NewBackendServiceAPI(spec *loads.Document) *BackendServiceAPI {
	return &BackendServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),

		AchievementAchievementTypeGetHandler: achievement.AchievementTypeGetHandlerFunc(func(params achievement.AchievementTypeGetParams) middleware.Responder {
			return middleware.NotImplemented("operation achievement.AchievementTypeGet has not yet been implemented")
		}),
		AttemptAttemptAttemptIDClosePatchHandler: attempt.AttemptAttemptIDClosePatchHandlerFunc(func(params attempt.AttemptAttemptIDClosePatchParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.AttemptAttemptIDClosePatch has not yet been implemented")
		}),
		AttemptAttemptPostHandler: attempt.AttemptPostHandlerFunc(func(params attempt.AttemptPostParams) middleware.Responder {
			return middleware.NotImplemented("operation attempt.AttemptPost has not yet been implemented")
		}),
		AuthAuthPostHandler: auth.AuthPostHandlerFunc(func(params auth.AuthPostParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthPost has not yet been implemented")
		}),
		LogsLogsGetHandler: logs.LogsGetHandlerFunc(func(params logs.LogsGetParams) middleware.Responder {
			return middleware.NotImplemented("operation logs.LogsGet has not yet been implemented")
		}),
		MailMailUserGetHandler: mail.MailUserGetHandlerFunc(func(params mail.MailUserGetParams) middleware.Responder {
			return middleware.NotImplemented("operation mail.MailUserGet has not yet been implemented")
		}),
		MailMailUserPostHandler: mail.MailUserPostHandlerFunc(func(params mail.MailUserPostParams) middleware.Responder {
			return middleware.NotImplemented("operation mail.MailUserPost has not yet been implemented")
		}),
		SessionSessionDeleteHandler: session.SessionDeleteHandlerFunc(func(params session.SessionDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation session.SessionDelete has not yet been implemented")
		}),
		SessionSessionPatchHandler: session.SessionPatchHandlerFunc(func(params session.SessionPatchParams) middleware.Responder {
			return middleware.NotImplemented("operation session.SessionPatch has not yet been implemented")
		}),
		TestTestGetHandler: test.TestGetHandlerFunc(func(params test.TestGetParams) middleware.Responder {
			return middleware.NotImplemented("operation test.TestGet has not yet been implemented")
		}),
		TestTestTestIDQuestionGetHandler: test.TestTestIDQuestionGetHandlerFunc(func(params test.TestTestIDQuestionGetParams) middleware.Responder {
			return middleware.NotImplemented("operation test.TestTestIDQuestionGet has not yet been implemented")
		}),
		UserUserAchievementGetHandler: user.UserAchievementGetHandlerFunc(func(params user.UserAchievementGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAchievementGet has not yet been implemented")
		}),
		UserUserAchievementPostHandler: user.UserAchievementPostHandlerFunc(func(params user.UserAchievementPostParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAchievementPost has not yet been implemented")
		}),
		UserUserAttemptAttemptIDAnswerGetHandler: user.UserAttemptAttemptIDAnswerGetHandlerFunc(func(params user.UserAttemptAttemptIDAnswerGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAttemptAttemptIDAnswerGet has not yet been implemented")
		}),
		UserUserAttemptAttemptIDAnswerPutHandler: user.UserAttemptAttemptIDAnswerPutHandlerFunc(func(params user.UserAttemptAttemptIDAnswerPutParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAttemptAttemptIDAnswerPut has not yet been implemented")
		}),
		UserUserAttemptAttemptIDAzbelTestGetHandler: user.UserAttemptAttemptIDAzbelTestGetHandlerFunc(func(params user.UserAttemptAttemptIDAzbelTestGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAttemptAttemptIDAzbelTestGet has not yet been implemented")
		}),
		UserUserAttemptAttemptIDKlimovTestGetHandler: user.UserAttemptAttemptIDKlimovTestGetHandlerFunc(func(params user.UserAttemptAttemptIDKlimovTestGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAttemptAttemptIDKlimovTestGet has not yet been implemented")
		}),
		UserUserAttemptGetHandler: user.UserAttemptGetHandlerFunc(func(params user.UserAttemptGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAttemptGet has not yet been implemented")
		}),
		AllUsersUserGetHandler: all_users.UserGetHandlerFunc(func(params all_users.UserGetParams) middleware.Responder {
			return middleware.NotImplemented("operation all_users.UserGet has not yet been implemented")
		}),
		UserUserPostHandler: user.UserPostHandlerFunc(func(params user.UserPostParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserPost has not yet been implemented")
		}),
		UserUserProfileHandler: user.UserProfileHandlerFunc(func(params user.UserProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserProfile has not yet been implemented")
		}),
		UserUserAchievementAchiveIDImageGetHandler: user.UserAchievementAchiveIDImageGetHandlerFunc(func(params user.UserAchievementAchiveIDImageGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAchievementAchiveIDImageGet has not yet been implemented")
		}),
		UserUserAvatarGetHandler: user.UserAvatarGetHandlerFunc(func(params user.UserAvatarGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAvatarGet has not yet been implemented")
		}),
		UserUserAvatarPutHandler: user.UserAvatarPutHandlerFunc(func(params user.UserAvatarPutParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserAvatarPut has not yet been implemented")
		}),
	}
}

/*BackendServiceAPI Главный backend обработчик для проекта digital trace */
type BackendServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - image/png
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AchievementAchievementTypeGetHandler sets the operation handler for the achievement type get operation
	AchievementAchievementTypeGetHandler achievement.AchievementTypeGetHandler
	// AttemptAttemptAttemptIDClosePatchHandler sets the operation handler for the attempt attempt Id close patch operation
	AttemptAttemptAttemptIDClosePatchHandler attempt.AttemptAttemptIDClosePatchHandler
	// AttemptAttemptPostHandler sets the operation handler for the attempt post operation
	AttemptAttemptPostHandler attempt.AttemptPostHandler
	// AuthAuthPostHandler sets the operation handler for the auth post operation
	AuthAuthPostHandler auth.AuthPostHandler
	// LogsLogsGetHandler sets the operation handler for the logs get operation
	LogsLogsGetHandler logs.LogsGetHandler
	// MailMailUserGetHandler sets the operation handler for the mail user get operation
	MailMailUserGetHandler mail.MailUserGetHandler
	// MailMailUserPostHandler sets the operation handler for the mail user post operation
	MailMailUserPostHandler mail.MailUserPostHandler
	// SessionSessionDeleteHandler sets the operation handler for the session delete operation
	SessionSessionDeleteHandler session.SessionDeleteHandler
	// SessionSessionPatchHandler sets the operation handler for the session patch operation
	SessionSessionPatchHandler session.SessionPatchHandler
	// TestTestGetHandler sets the operation handler for the test get operation
	TestTestGetHandler test.TestGetHandler
	// TestTestTestIDQuestionGetHandler sets the operation handler for the test test Id question get operation
	TestTestTestIDQuestionGetHandler test.TestTestIDQuestionGetHandler
	// UserUserAchievementGetHandler sets the operation handler for the user achievement get operation
	UserUserAchievementGetHandler user.UserAchievementGetHandler
	// UserUserAchievementPostHandler sets the operation handler for the user achievement post operation
	UserUserAchievementPostHandler user.UserAchievementPostHandler
	// UserUserAttemptAttemptIDAnswerGetHandler sets the operation handler for the user attempt attempt Id answer get operation
	UserUserAttemptAttemptIDAnswerGetHandler user.UserAttemptAttemptIDAnswerGetHandler
	// UserUserAttemptAttemptIDAnswerPutHandler sets the operation handler for the user attempt attempt Id answer put operation
	UserUserAttemptAttemptIDAnswerPutHandler user.UserAttemptAttemptIDAnswerPutHandler
	// UserUserAttemptAttemptIDAzbelTestGetHandler sets the operation handler for the user attempt attempt Id azbel test get operation
	UserUserAttemptAttemptIDAzbelTestGetHandler user.UserAttemptAttemptIDAzbelTestGetHandler
	// UserUserAttemptAttemptIDKlimovTestGetHandler sets the operation handler for the user attempt attempt Id klimov test get operation
	UserUserAttemptAttemptIDKlimovTestGetHandler user.UserAttemptAttemptIDKlimovTestGetHandler
	// UserUserAttemptGetHandler sets the operation handler for the user attempt get operation
	UserUserAttemptGetHandler user.UserAttemptGetHandler
	// AllUsersUserGetHandler sets the operation handler for the user get operation
	AllUsersUserGetHandler all_users.UserGetHandler
	// UserUserPostHandler sets the operation handler for the user post operation
	UserUserPostHandler user.UserPostHandler
	// UserUserProfileHandler sets the operation handler for the user profile operation
	UserUserProfileHandler user.UserProfileHandler
	// UserUserAchievementAchiveIDImageGetHandler sets the operation handler for the user achievement achive Id image get operation
	UserUserAchievementAchiveIDImageGetHandler user.UserAchievementAchiveIDImageGetHandler
	// UserUserAvatarGetHandler sets the operation handler for the user avatar get operation
	UserUserAvatarGetHandler user.UserAvatarGetHandler
	// UserUserAvatarPutHandler sets the operation handler for the user avatar put operation
	UserUserAvatarPutHandler user.UserAvatarPutHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BackendServiceAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BackendServiceAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BackendServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BackendServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BackendServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BackendServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BackendServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BackendServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BackendServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BackendServiceAPI
func (o *BackendServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AchievementAchievementTypeGetHandler == nil {
		unregistered = append(unregistered, "achievement.AchievementTypeGetHandler")
	}
	if o.AttemptAttemptAttemptIDClosePatchHandler == nil {
		unregistered = append(unregistered, "attempt.AttemptAttemptIDClosePatchHandler")
	}
	if o.AttemptAttemptPostHandler == nil {
		unregistered = append(unregistered, "attempt.AttemptPostHandler")
	}
	if o.AuthAuthPostHandler == nil {
		unregistered = append(unregistered, "auth.AuthPostHandler")
	}
	if o.LogsLogsGetHandler == nil {
		unregistered = append(unregistered, "logs.LogsGetHandler")
	}
	if o.MailMailUserGetHandler == nil {
		unregistered = append(unregistered, "mail.MailUserGetHandler")
	}
	if o.MailMailUserPostHandler == nil {
		unregistered = append(unregistered, "mail.MailUserPostHandler")
	}
	if o.SessionSessionDeleteHandler == nil {
		unregistered = append(unregistered, "session.SessionDeleteHandler")
	}
	if o.SessionSessionPatchHandler == nil {
		unregistered = append(unregistered, "session.SessionPatchHandler")
	}
	if o.TestTestGetHandler == nil {
		unregistered = append(unregistered, "test.TestGetHandler")
	}
	if o.TestTestTestIDQuestionGetHandler == nil {
		unregistered = append(unregistered, "test.TestTestIDQuestionGetHandler")
	}
	if o.UserUserAchievementGetHandler == nil {
		unregistered = append(unregistered, "user.UserAchievementGetHandler")
	}
	if o.UserUserAchievementPostHandler == nil {
		unregistered = append(unregistered, "user.UserAchievementPostHandler")
	}
	if o.UserUserAttemptAttemptIDAnswerGetHandler == nil {
		unregistered = append(unregistered, "user.UserAttemptAttemptIDAnswerGetHandler")
	}
	if o.UserUserAttemptAttemptIDAnswerPutHandler == nil {
		unregistered = append(unregistered, "user.UserAttemptAttemptIDAnswerPutHandler")
	}
	if o.UserUserAttemptAttemptIDAzbelTestGetHandler == nil {
		unregistered = append(unregistered, "user.UserAttemptAttemptIDAzbelTestGetHandler")
	}
	if o.UserUserAttemptAttemptIDKlimovTestGetHandler == nil {
		unregistered = append(unregistered, "user.UserAttemptAttemptIDKlimovTestGetHandler")
	}
	if o.UserUserAttemptGetHandler == nil {
		unregistered = append(unregistered, "user.UserAttemptGetHandler")
	}
	if o.AllUsersUserGetHandler == nil {
		unregistered = append(unregistered, "all_users.UserGetHandler")
	}
	if o.UserUserPostHandler == nil {
		unregistered = append(unregistered, "user.UserPostHandler")
	}
	if o.UserUserProfileHandler == nil {
		unregistered = append(unregistered, "user.UserProfileHandler")
	}
	if o.UserUserAchievementAchiveIDImageGetHandler == nil {
		unregistered = append(unregistered, "user.UserAchievementAchiveIDImageGetHandler")
	}
	if o.UserUserAvatarGetHandler == nil {
		unregistered = append(unregistered, "user.UserAvatarGetHandler")
	}
	if o.UserUserAvatarPutHandler == nil {
		unregistered = append(unregistered, "user.UserAvatarPutHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BackendServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BackendServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *BackendServiceAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BackendServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BackendServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "image/png":
			result["image/png"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BackendServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the backend service API
func (o *BackendServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BackendServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/achievement/type"] = achievement.NewAchievementTypeGet(o.context, o.AchievementAchievementTypeGetHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/attempt/{attempt_id}/close"] = attempt.NewAttemptAttemptIDClosePatch(o.context, o.AttemptAttemptAttemptIDClosePatchHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/attempt"] = attempt.NewAttemptPost(o.context, o.AttemptAttemptPostHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth"] = auth.NewAuthPost(o.context, o.AuthAuthPostHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/logs"] = logs.NewLogsGet(o.context, o.LogsLogsGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mail/user"] = mail.NewMailUserGet(o.context, o.MailMailUserGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mail/user"] = mail.NewMailUserPost(o.context, o.MailMailUserPostHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/session"] = session.NewSessionDelete(o.context, o.SessionSessionDeleteHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/session"] = session.NewSessionPatch(o.context, o.SessionSessionPatchHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/test"] = test.NewTestGet(o.context, o.TestTestGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/test/{test_id}/questions"] = test.NewTestTestIDQuestionGet(o.context, o.TestTestTestIDQuestionGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/achievement"] = user.NewUserAchievementGet(o.context, o.UserUserAchievementGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/achievement"] = user.NewUserAchievementPost(o.context, o.UserUserAchievementPostHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/attempt/{attempt_id}/answer"] = user.NewUserAttemptAttemptIDAnswerGet(o.context, o.UserUserAttemptAttemptIDAnswerGetHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/attempt/{attempt_id}/answer"] = user.NewUserAttemptAttemptIDAnswerPut(o.context, o.UserUserAttemptAttemptIDAnswerPutHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/attempt/{attempt_id}/azbel_test"] = user.NewUserAttemptAttemptIDAzbelTestGet(o.context, o.UserUserAttemptAttemptIDAzbelTestGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/attempt/{attempt_id}/klimov_test"] = user.NewUserAttemptAttemptIDKlimovTestGet(o.context, o.UserUserAttemptAttemptIDKlimovTestGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/attempt"] = user.NewUserAttemptGet(o.context, o.UserUserAttemptGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user"] = all_users.NewUserGet(o.context, o.AllUsersUserGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = user.NewUserPost(o.context, o.UserUserPostHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/profile"] = user.NewUserProfile(o.context, o.UserUserProfileHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/achievement/{achive_id}/image"] = user.NewUserAchievementAchiveIDImageGet(o.context, o.UserUserAchievementAchiveIDImageGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/avatar"] = user.NewUserAvatarGet(o.context, o.UserUserAvatarGetHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/avatar"] = user.NewUserAvatarPut(o.context, o.UserUserAvatarPutHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BackendServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BackendServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BackendServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BackendServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BackendServiceAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}

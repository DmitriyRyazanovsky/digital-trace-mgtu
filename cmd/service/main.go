package main

import (
	"fmt"
	"math/rand"
	"mgtu/digital-trace/main-backend-service/internal/config"
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/features/logging"
	fileworker "mgtu/digital-trace/main-backend-service/internal/file_worker"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations"
	"mgtu/digital-trace/main-backend-service/internal/handlers"
	"mgtu/digital-trace/main-backend-service/internal/mail_service"
	"mgtu/digital-trace/main-backend-service/internal/sequrity/jwt_service.go"
	"time"

	"os"
	"os/signal"
	"syscall"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
)

func mainError(err error, log *logging.Logger) {
	err = errors.Wrap(err, "[main()]")
	if log != nil {
		(*log).Error(err.Error())
	} else {
		fmt.Println("ERROR:", err.Error())
	}

	panic(err)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Заводим канал для правного завершения работы сервиса
	quit := make(chan os.Signal, 1)

	// До тех пор, пока не будет прокинут системый shutdown, лочим канал
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Подготавливаем конфиг
	conf, err := config.New()
	if err != nil {
		err = errors.Wrap(err, "unable to create config: [config.New]")
		panic(err)
	}

	// Подготавливаем логи
	logConfig := logging.NewConfig()
	logConfig.SetInFile(conf.Logger.LogInFile)
	logConfig.SetOutputDir(conf.Logger.OutputDir)
	logConfig.SetLevel(logging.InfoLevel)
	log, err := logging.New(logConfig)
	if err != nil {
		err = errors.Wrap(err, "unable to create logging: [logging.New()]")
		panic(err)
	}

	//* Считываем сваггер спецификацию
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		err = errors.Wrap(err, "unable to create load swagger spec: [loads.Analyzed()]")
		log.Error(errors.Wrap(err, "[main()]").Error())
		panic(err)
	}

	//* Подготавливаем почту
	mail, err := mail_service.NewMail(conf.Mail.ServerName, conf.Mail.Username, conf.Mail.Password)
	if err != nil {
		err = errors.Wrap(err, "unable to create mail service: [mail_service.NewMail()]")
		log.Error(errors.Wrap(err, "[main()]").Error())
		panic(err)
	}

	fileWorker := fileworker.NewFileWorker(log)

	//* Подготавливаем БД
	db, err := database.NewDatabase(database.NewDatabaseIn{
		Database:         conf.Database.Database,
		PasswordFilePath: conf.Database.PasswordFilePath,
		UserName:         conf.Database.UserName,
		Url:              conf.Database.Url,
		SslMode:          conf.Database.SslMode,
		DriverName:       conf.Database.DriverName,
	})
	if err != nil {
		err = errors.Wrap(err, "database.NewDatabase(database.NewDatabaseIn{...})")
		mainError(err, &log)
	}

	//* создаём JWT объект для работы с сессиями
	jwtService, err := jwt_service.NewJWtService(jwt_service.NewJWtServiceIn{
		Db:                db,
		Iss:               conf.Jwt.Iss,
		Log:               log,
		SigningFilePath:   conf.Jwt.SigningFilePath,
		SessionTokenLen:   conf.Jwt.SessionTokenLen,
		SessionSigningLen: conf.Jwt.SessionSigningLen,
		AccessTokenExp:    conf.Jwt.AccessTokenExp,
		RefreshTokenExp:   conf.Jwt.RefreshTokenExp,
	})
	if err != nil {
		err = errors.Wrap(err, "[jwt_service.NewJWtService()]")
		log.Error(errors.Wrap(err, "[main()]").Error())
		panic(err)
	}

	// Подготавливаем пакет handlers для реализации запросов
	handlers := handlers.NewHandler(
		db,
		fileWorker,
		mail,
		jwtService,
		log,
	)

	// Запускаем сервис
	api := operations.NewBackendServiceAPI(swaggerSpec)
	handlers.Register(api)
	server := restapi.NewServer(api)
	server.Port = conf.Service.Port
	server.Host = conf.Service.Host

	log.Info("start service")

	// Сёрвим сервис на наличие panic ошибок
	if err = server.Serve(); err != nil {
		err = errors.Wrap(err, "unable to serve service: [server.Serve()]")
		log.Error(errors.Wrap(err, "[main()]").Error())
		panic(err)
	}

	<-quit

	log.Info("shutdown database")
	// Закрытие БД
	err = db.Close()
	if err != nil {
		err = errors.Wrap(err, "unable to close database: [db.Close()]")
		log.Error(errors.Wrap(err, "[main()]").Error())
		panic(err)
	}

	log.Info("shutdown service")

	// Завершаем сеанс
	err = server.Shutdown()
	if err != nil {
		err = errors.Wrap(err, "unable to shutdown service: [server.Shutdown()]")
		log.Error(errors.Wrap(err, "[main()]").Error())
		panic(err)
	}
}

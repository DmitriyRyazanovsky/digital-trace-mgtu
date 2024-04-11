package handlers

import (
	"bufio"
	"encoding/json"
	"mgtu/digital-trace/main-backend-service/internal/gen/models"
	"mgtu/digital-trace/main-backend-service/internal/gen/restapi/operations/logs"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

func (h *Handler) logsGen500(err error) middleware.Responder {
	err = errors.Wrapf(err, "handler error: [logsGet]")
	h.log.Error(err.Error())
	return logs.NewLogsGetInternalServerError().WithPayload(
		&models.Error500{
			Error: err.Error(),
		},
	)
}

func (h *Handler) logsGet(params logs.LogsGetParams) middleware.Responder {
	file, err := os.Open("./logs/logs/console.log")
	if err != nil {
		return h.logsGen500(errors.Wrap(err, `unable to open file: [os.Open("./logs/logs/console.log")]`))
	}
	defer file.Close()

	out := []*logs.LogsGetOKBodyItems0{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		var log *logs.LogsGetOKBodyItems0
		if err := json.Unmarshal(line, &log); err != nil {
			return h.logsGen500(errors.Wrap(err, "error when unmarshal line: [json.Unmarshal(line, &log)]"))
		}
		out = append(out, log)
	}

	if err := scanner.Err(); err != nil {
		return h.logsGen500(errors.Wrap(err, "error when scan file: [scanner.Err()]"))
	}

	return logs.NewLogsGetOK().WithPayload(out)
}

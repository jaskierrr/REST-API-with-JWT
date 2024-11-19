package handlers

import (
	"main/internal/controller"
	"main/restapi/operations"
	"log/slog"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
)

type handlers struct {
	logger     *slog.Logger
	controller controller.Controller
}

var validate *validator.Validate

type Handlers interface {
	GetUsersID(params operations.GetUsersIDStatusParams) middleware.Responder
	PostUsers(params operations.PostUsersParams) middleware.Responder
	DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder
	GetUsersLeader(params operations.GetUsersLeaderboardParams) middleware.Responder

	Link(api *operations.CryptoAPI)
}

func New(controller controller.Controller, validator *validator.Validate, logger *slog.Logger) Handlers {
	validate = validator
	return &handlers{
		logger:     logger,
		controller: controller,
	}
}

func (h *handlers) Link(api *operations.CryptoAPI) {
	api.GetUsersLeaderboardHandler = operations.GetUsersLeaderboardHandlerFunc(h.GetUsersLeader)
	api.GetUsersIDStatusHandler = operations.GetUsersIDStatusHandlerFunc(h.GetUsersID)
	api.PostUsersHandler = operations.PostUsersHandlerFunc(h.PostUsers)
	api.DeleteUsersIDHandler = operations.DeleteUsersIDHandlerFunc(h.DeleteUsersID)
}

func convertI64tStr(integer int64) string {
	return strconv.FormatInt(integer, 10)
}

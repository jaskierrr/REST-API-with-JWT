package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/controller"
	"main/internal/lib/jwt"
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
	PostUsers(params operations.PostUsersParams) middleware.Responder
	GetUsersID(params operations.GetUsersIDStatusParams, principal interface{}) middleware.Responder
	DeleteUsersID(params operations.DeleteUsersIDParams, principal interface{}) middleware.Responder
	GetUsersLeader(params operations.GetUsersLeaderboardParams, principal interface{}) middleware.Responder
	PostTask(params operations.PostUsersUserIDTaskCompleteParams, principal interface{}) middleware.Responder
	PostRef(params operations.PostUsersUserIDReferrerParams, principal interface{}) middleware.Responder

	Login(params operations.PostUsersLoginParams) middleware.Responder

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
	api.PostUsersLoginHandler = operations.PostUsersLoginHandlerFunc(h.Login)
	api.PostUsersUserIDTaskCompleteHandler = operations.PostUsersUserIDTaskCompleteHandlerFunc(h.PostTask)
	api.PostUsersUserIDReferrerHandler = operations.PostUsersUserIDReferrerHandlerFunc(h.PostRef)

	//! Объявление Middleware
	api.BearerAuth = jwt.ValidateToken
}

func convertI64tStr(integer int64) string {
	return strconv.FormatInt(integer, 10)
}

package handlers

import (
	"errors"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) Login(params operations.PostUsersLoginParams) middleware.Responder {
	h.logger.Info(
		"Trying to Login",
		slog.Any("user email", params.User.Email),
	)

	err := validate.Struct(params.User)
	if err != nil {
		h.logger.Error(
			"Failed to Login",
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersLoginDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to Login: " + errors.New("all fields must be filled").Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	token, err := h.controller.Login(ctx, *params.User)

	if err != nil {
		h.logger.Error(
			"Failed to Login",
			slog.Any("user email", params.User.Email),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersLoginDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to Login: " + errors.New("incorrect username or password").Error(),
			},
		})
	}

	return operations.NewPostUsersLoginCreated().WithPayload(token)
}

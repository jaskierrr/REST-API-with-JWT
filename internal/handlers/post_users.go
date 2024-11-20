package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) PostUsers(params operations.PostUsersParams) middleware.Responder {
	h.logger.Info(
		"Trying to POST user in storage",
		slog.Any("user", params.User),
	)

	err := validate.Struct(params.User)
	if err != nil {
		h.logger.Error(
			"Failed to POST user in storage",
			slog.Any("user", params.User),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST user in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	user, err := h.controller.PostUser(ctx, *params.User)

	if err != nil {
		h.logger.Error(
			"Failed to POST user in storage",
			slog.Any("user", params.User),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST user in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostUsersCreated().WithPayload(&user)
}

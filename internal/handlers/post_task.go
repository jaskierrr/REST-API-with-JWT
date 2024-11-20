package handlers

import (
	"errors"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) PostTask(params operations.PostUsersUserIDTaskCompleteParams, principal interface{}) middleware.Responder {
	h.logger.Info(
		"Trying to POST task in storage",
		slog.Any("user", params.UserID),
	)

	if params.UserID == 0 {
		err := errors.New("user_id cant be = 0")
		h.logger.Error(
			"Failed to POST task in storage",
			slog.Any("user", params.UserID),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersUserIDTaskCompleteDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST task in storage " + err.Error(),
			},
		})
	}

	err := validate.Struct(params.Task)
	if err != nil {
		h.logger.Error(
			"Failed to POST task in storage",
			slog.Any("user", params.UserID),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersUserIDTaskCompleteDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST task in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	task, err := h.controller.PostTask(ctx, *params.Task, params.UserID)

	if err != nil {
		h.logger.Error(
			"Failed to POST task in storage",
			slog.Any("user", params.UserID),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersUserIDTaskCompleteDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST task in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostUsersUserIDTaskCompleteCreated().WithPayload(&task)
}

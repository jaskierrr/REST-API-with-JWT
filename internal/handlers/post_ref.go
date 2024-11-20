package handlers

import (
	"errors"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) PostRef(params operations.PostUsersUserIDReferrerParams, principal interface{}) middleware.Responder {
	h.logger.Info(
		"Trying to POST referrer in storage",
		slog.Any("user", params.UserID),
	)

	if params.UserID == 0 {
		err := errors.New("user_id cant be = 0")
		h.logger.Error(
			"Failed to POST referrer in storage",
			slog.Any("user", params.UserID),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersUserIDReferrerDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST referrer in storage " + err.Error(),
			},
		})
	}

	err := validate.Struct(params.Refferer)
	if err != nil {
		h.logger.Error(
			"Failed to POST referrer in storage",
			slog.Any("user", params.UserID),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersUserIDReferrerDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST referrer in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	ref, err := h.controller.PostRef(ctx, *params.Refferer, params.UserID)

	if err != nil {
		h.logger.Error(
			"Failed to POST referrer in storage",
			slog.Any("user", params.UserID),
			slog.String("error", err.Error()),
		)
		return operations.NewPostUsersUserIDReferrerDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST referrer in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostUsersUserIDReferrerCreated().WithPayload(&ref)
}

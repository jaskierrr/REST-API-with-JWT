package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsersLeader(params operations.GetUsersLeaderboardParams, principal interface{}) middleware.Responder {
	h.logger.Info("Trying to GET users from storage")

	ctx := params.HTTPRequest.Context()
	users, err := h.controller.GetUsers(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to GET users from storage",
			slog.String("error", err.Error()),
		)
		return operations.NewGetUsersLeaderboardDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET users in storage " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersLeaderboardOK().WithPayload(users)
}

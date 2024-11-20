package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) UpdateBalance(params operations.PatchUsersIDBalanceParams, principal interface{}) middleware.Responder {
	h.logger.Info(
		"Trying to UPDATE user balance in storage",
		slog.Any("user", params.ID),
	)

	ctx := params.HTTPRequest.Context()
	balance, err := h.controller.UpdateBalance(ctx, params.ID, params.Balance.Amount)


	if err != nil {
		h.logger.Error(
			"Failed to UPDATE user balance in storage",
			slog.Any("user", params.ID),
			slog.String("error", err.Error()),
		)
		return operations.NewPatchUsersIDBalanceDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to UPDATE user balance in storage " + err.Error(),
			},
		})
	}

	return operations.NewPatchUsersIDBalanceCreated().WithPayload(&balance)
}

package handlers

import (
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) DeleteUsersID(params operations.DeleteUsersIDParams, principal interface{}) middleware.Responder {
	h.logger.Info("Trying to DELETE user from storage, user id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewDeleteUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE user from storage, user id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteUserID(ctx, params)

	if err != nil {
		return operations.NewDeleteUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE user from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteUsersIDNoContent()
}

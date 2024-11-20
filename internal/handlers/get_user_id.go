package handlers

import (
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsersID(params operations.GetUsersIDStatusParams, principal interface{}) middleware.Responder {
	h.logger.Info("Trying to GET user from storage, user id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewGetUsersIDStatusDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET user from storage, user id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	user, err := h.controller.GetUserID(ctx, params)

	if err != nil {
		return operations.NewGetUsersIDStatusDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET user from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersIDStatusOK().WithPayload(&user)
}

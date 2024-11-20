package service

import (
	"context"
	"errors"
	"main/api/restapi/operations"
	"main/internal/lib/jwt"
	"main/internal/models"
	"strings"
)

func (s service) GetUserID(ctx context.Context, params operations.GetUsersIDStatusParams) (models.User, error) {
	token := params.HTTPRequest.Header.Get("Authorization")
	if token == "" {
		return models.User{}, errors.New("Unauthorized")
	}

	token = strings.TrimPrefix(token, "Bearer ")

	err := jwt.ValidateToken(token, s.secret)
	if err != nil {
		return models.User{}, err
	}

	// s.logger.Info("GetUserID authorization complete" + token)

	return s.repo.GetUserID(ctx, int(params.ID))
}

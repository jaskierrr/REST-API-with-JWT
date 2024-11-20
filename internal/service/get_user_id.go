package service

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (s service) GetUserID(ctx context.Context, params operations.GetUsersIDStatusParams) (models.User, error) {


	user, err := s.repo.GetUserID(ctx, int(params.ID))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

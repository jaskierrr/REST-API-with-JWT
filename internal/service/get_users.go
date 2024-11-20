package service

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (s service) GetUsers(ctx context.Context, params operations.GetUsersLeaderboardParams) ([]*models.User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

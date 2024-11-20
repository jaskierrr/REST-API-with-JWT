package service

import (
	"context"
	"main/internal/models"
)

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetUsers(ctx)
}

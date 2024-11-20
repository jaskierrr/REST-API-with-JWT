package service

import (
	"context"
	"main/internal/models"
)

func (s service) UpdateBalance(ctx context.Context, id int64, amount int64) (models.User, error) {
	user, err := s.repo.UpdateBalance(ctx, id, amount)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

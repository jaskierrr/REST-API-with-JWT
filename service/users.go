package service

import (
	"main/models"
	"context"

	"github.com/google/uuid"
)

func (s service) GetUserID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetUserID(ctx, id)
}

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	id, _ := uuid.NewUUID()

	user := models.User{
		ID:        int64(id.ID()),
		Name: userData.Name,
	}
	user, err := s.repo.PostUser(ctx, user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s service) DeleteUserID(ctx context.Context, id int) error {
	_, err := s.repo.DeleteUser(ctx, id)
	return err
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetUsers(ctx)
}

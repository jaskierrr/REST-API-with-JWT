//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"main/models"
	repo "main/repositories"
	"context"
	"log/slog"
)

type service struct {
	logger *slog.Logger

	repo repo.Repository
}

type Service interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func New(repo repo.Repository, logger *slog.Logger) Service {
	return &service{
		logger: logger,
		repo: repo,
	}
}

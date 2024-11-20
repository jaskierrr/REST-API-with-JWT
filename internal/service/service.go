//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"
	repo "main/internal/repositories"
)

type service struct {
	logger   *slog.Logger
	repo     repo.Repository
	secret   string
	tokenTTL int
}

type Service interface {
	GetUserID(ctx context.Context, params operations.GetUsersIDStatusParams) (models.User, error)
	PostUser(ctx context.Context, userData models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, params operations.DeleteUsersIDParams) error
	GetUsers(ctx context.Context, params operations.GetUsersLeaderboardParams) ([]*models.User, error)
	PostTask(ctx context.Context, taskData models.NewTask, userID int64) (models.Task, error)
	PostRef(ctx context.Context, refData models.NewReferrer, userID int64) (models.Referrer, error)
	
	Login(ctx context.Context, user models.NewUser) (string, error)
}

func New(repo repo.Repository, logger *slog.Logger, secret string, tokenTTL int) Service {
	return &service{
		logger:   logger,
		repo:     repo,
		secret:   secret,
		tokenTTL: tokenTTL,
	}
}

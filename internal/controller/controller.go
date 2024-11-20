package controller

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"
	"main/internal/service"
)

type controller struct {
	logger  *slog.Logger
	service service.Service
}

type Controller interface {
	GetUserID(ctx context.Context, params operations.GetUsersIDStatusParams) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, params operations.DeleteUsersIDParams) error
	GetUsers(ctx context.Context, params operations.GetUsersLeaderboardParams) ([]*models.User, error)
	PostTask(ctx context.Context, taskData models.NewTask, userID int64) (models.Task, error)
	PostRef(ctx context.Context, refData models.NewReferrer, userID int64) (models.Referrer, error)

	Login(ctx context.Context, userData models.NewUser) (string, error)
}

func New(service service.Service, logger *slog.Logger) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}

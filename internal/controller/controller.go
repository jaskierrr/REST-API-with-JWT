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
	DeleteUserID(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]*models.User, error)
	Login(ctx context.Context, userData models.NewUser) (string, error)
}

func New(service service.Service, logger *slog.Logger) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}

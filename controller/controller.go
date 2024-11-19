package controller

import (
	"context"
	"log/slog"
	"main/models"
	"main/service"
)

type controller struct {
	logger  *slog.Logger
	service service.Service
}

type Controller interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func New(service service.Service, logger *slog.Logger) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}

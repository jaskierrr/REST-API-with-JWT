package controller

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (c controller) GetUserID(ctx context.Context, params operations.GetUsersIDStatusParams) (models.User, error) {
	return c.service.GetUserID(ctx, params)
}

func (c controller) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	return c.service.PostUser(ctx, userData)
}

func (c controller) DeleteUserID(ctx context.Context, params operations.DeleteUsersIDParams) error {
	return c.service.DeleteUserID(ctx, params)
}

func (c controller) GetUsers(ctx context.Context, params operations.GetUsersLeaderboardParams) ([]*models.User, error) {
	return c.service.GetUsers(ctx, params)
}

func (c controller) Login(ctx context.Context, userData models.NewUser) (string, error) {
	return c.service.Login(ctx, userData)
}

func (c controller) UpdateBalance(ctx context.Context, id int64, amount int64) (models.User, error) {
	return c.service.UpdateBalance(ctx, id, amount)
}

package controller

import (
	"context"
	"main/internal/models"
)

func (c controller) PostTask(ctx context.Context, taskData models.NewTask, userID int64) (models.Task, error) {
	return c.service.PostTask(ctx, taskData, userID)
}

package service

import (
	"context"
	"main/internal/models"

	"github.com/google/uuid"
)

func (s service) PostTask(ctx context.Context, taskData models.NewTask, userID int64) (models.Task, error) {
	id, _ := uuid.NewUUID()

	task := models.Task{
		ID:   int64(id.ID()),
		UserID: userID,
		Name: taskData.Name,
	}

	task, err := s.repo.PostTask(ctx, task)

	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

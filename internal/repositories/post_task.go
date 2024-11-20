package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)


func (repo *repository) PostTask(ctx context.Context, taskData models.Task) (models.Task, error) {
	sql, args, err := sq.Insert("tasks").
										Columns("id", "user_id", "name").
										Values(taskData.ID, taskData.UserID, taskData.Name).
										Suffix("returning *").
										PlaceholderFormat(sq.Dollar).
										ToSql()
	if err != nil {
		return models.Task{}, err
	}

	task := models.Task{}
	err = repo.db.
		GetConn().
		QueryRow(ctx, sql, args...).
		Scan(&task.ID, &task.UserID, &task.Name)

	if err != nil {
		return models.Task{}, err
	}

	repo.logger.Info(
		"Success POST task in storage",
		slog.Any("ID", task.ID),
	)

	return task, nil
}

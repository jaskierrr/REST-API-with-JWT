//go:generate mockgen -source=./repository.go -destination=../../mocks/repo_mock.go -package=mock
package repo

import (
	"context"
	"log/slog"
	"main/internal/database"
	"main/internal/models"
)

type repository struct {
	db     database.DB
	logger *slog.Logger
}

type Repository interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.User, passHash []byte) (models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]*models.User, error)
	UpdateBalance(ctx context.Context, id int64, amount int64) (models.User, error)
	PostTask(ctx context.Context, taskData models.Task) (models.Task, error)
	PostRef(ctx context.Context, refData models.Referrer) (models.Referrer, error)

	GetUserEmail(ctx context.Context, id string) (models.User, error)
}

func NewUserRepo(db database.DB, logger *slog.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

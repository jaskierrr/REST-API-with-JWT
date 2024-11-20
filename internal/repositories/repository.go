//go:generate mockgen -source=./repository.go -destination=../../mocks/repo_mock.go -package=mock
package repo

import (
	"main/internal/database"
	"main/internal/models"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	deleteUserIDQuery = `delete from users where id = @userID`
)

type repository struct {
	db     database.DB
	logger *slog.Logger
}

type Repository interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.User, passHash []byte) (models.User, error)
	DeleteUser(ctx context.Context, id int) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserEmail(ctx context.Context, id string) (models.User, error)
}

func NewUserRepo(db database.DB, logger *slog.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

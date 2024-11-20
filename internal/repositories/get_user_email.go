package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) GetUserEmail(ctx context.Context, email string) (models.User, error) {
	sql, args, err := sq.Select("*").
											From("users").
											Where(sq.Eq{"email": email}).
											PlaceholderFormat(sq.Dollar).
											ToSql()

	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	err = repo.db.
		GetConn().
		QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Balance, &user.Email, &user.Password)

	if err != nil {
		return models.User{}, err
	}

	repo.logger.Info(
		"Success GET user from storage",
		slog.Any("ID", user.ID),
	)

	return user, nil
}

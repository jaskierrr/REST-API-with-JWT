package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) GetUserID(ctx context.Context, id int) (models.User, error) {
	sql, args, err := sq.Select("id, email, balance").
											From("users").
											Where(sq.Eq{"id": id}).
											PlaceholderFormat(sq.Dollar).
											ToSql()

	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	err = repo.db.
		GetConn().
		QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Email, &user.Balance)

	if err != nil {
		return models.User{}, err
	}

	repo.logger.Info(
		"Success GET user from storage",
		slog.Any("ID", user.ID),
	)

	return user, nil
}

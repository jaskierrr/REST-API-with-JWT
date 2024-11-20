package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)


func (repo *repository) PostUser(ctx context.Context, userData models.User, passHash []byte) (models.User, error) {
	sql, args, err := sq.Insert("users").
										Columns("id", "email", "password_hash").
										Values(userData.ID, userData.Email, passHash).
										Suffix("returning id, email, balance").
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
		"Success POST user from storage",
		slog.Any("ID", user.ID),
	)

	return user, nil
}

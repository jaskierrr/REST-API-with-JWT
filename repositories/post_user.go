package repo

import (
	"context"
	"log/slog"
	"main/models"

	sq "github.com/Masterminds/squirrel"
)


func (repo *repository) PostUser(ctx context.Context, userData models.User) (models.User, error) {
	sql, args, err := sq.Insert("users").
										Columns("id", "name").
										Values(userData.ID, userData.Name).
										Suffix("returning *").
										PlaceholderFormat(sq.Dollar).
										ToSql()
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	err = repo.db.
		GetConn().
		QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Name, &user.Balance)

	if err != nil {
		return models.User{}, err
	}

	repo.logger.Info(
		"Success POST user from storage",
		slog.Any("ID", user.ID),
	)

	return user, nil
}

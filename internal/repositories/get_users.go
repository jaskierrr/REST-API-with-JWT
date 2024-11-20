package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) GetUsers(ctx context.Context) ([]*models.User, error) {
	sql, args, err := sq.Select("id, email, balance").
												From("users").
												Where("balance > 0").
												OrderBy("balance DESC").
												Limit(5).
												PlaceholderFormat(sq.Dollar).
												ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repo.db.GetConn().Query(ctx, sql, args...)

	users := []*models.User{}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.ID, &user.Email, &user.Balance)

		if err != nil {
			repo.logger.Error("Error",
				slog.Any("error", err),
			)
			return nil, err
		}

		users = append(users, &user)
	}

	repo.logger.Info("Success GET users from storage")

	return users, err
}

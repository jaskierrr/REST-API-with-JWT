package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

// SELECT id, name, balance
// FROM users
// ORDER BY balance DESC
// LIMIT 5;

// getUsersQuery     = `select * from users`

func (repo *repository) GetUsers(ctx context.Context) ([]*models.User, error) {
	sql, args, err := sq.Select("*").
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

		err := rows.Scan(&user.ID, &user.Name, &user.Balance)

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

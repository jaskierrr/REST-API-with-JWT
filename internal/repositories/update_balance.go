package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)


func (repo *repository) UpdateBalance(ctx context.Context, id int64, amount int64) (models.User, error) {
	repo.logger.Info(
		"balance in data",
		slog.Any("amount", amount),
	)
	sql, args, err := sq.Update("users").
										Set("balance", sq.Expr("balance + ?", amount)).
										Where(sq.Eq{"id": id}).
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
		"Success UPDATE user balance from storage",
		slog.Any("ID", user.ID),
	)

	return user, nil
}

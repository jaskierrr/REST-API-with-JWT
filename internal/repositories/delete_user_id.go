package repo

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *repository) DeleteUser(ctx context.Context, id int) (pgconn.CommandTag, error) {
	args := pgx.NamedArgs{
		"userID": id,
	}
	commandTag, err := repo.db.
		GetConn().
		Exec(ctx, deleteUserIDQuery, args)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	repo.logger.Info(
		"Success DELETE user from storage",
		slog.Any("ID", id),
	)

	return commandTag, nil
}

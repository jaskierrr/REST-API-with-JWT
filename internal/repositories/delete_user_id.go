package repo

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) DeleteUser(ctx context.Context, id int) error {
	sql, args, err := sq.Delete("*").
												From("users").
												Where(sq.Eq{"id": id}).
												PlaceholderFormat(sq.Dollar).
												ToSql()

	if err != nil {
		return err
	}
	_, err = repo.db.
					GetConn().
					Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	repo.logger.Info(
		"Success DELETE user from storage",
		slog.Any("ID", id),
	)

	return nil
}

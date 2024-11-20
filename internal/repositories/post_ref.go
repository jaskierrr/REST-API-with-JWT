package repo

import (
	"context"
	"log/slog"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)


func (repo *repository) PostRef(ctx context.Context, refData models.Referrer) (models.Referrer, error) {
	sql, args, err := sq.Insert("referrers").
										Columns("id", "user_id", "referrer").
										Values(refData.ID, refData.UserID, refData.Referrer).
										Suffix("returning *").
										PlaceholderFormat(sq.Dollar).
										ToSql()
	if err != nil {
		return models.Referrer{}, err
	}

	ref := models.Referrer{}
	err = repo.db.
		GetConn().
		QueryRow(ctx, sql, args...).
		Scan(&ref.ID, &ref.UserID, &ref.Referrer)

	if err != nil {
		return models.Referrer{}, err
	}

	repo.logger.Info(
		"Success POST task in storage",
		slog.Any("ID", ref.ID),
	)

	return ref, nil
}

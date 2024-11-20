package service

import (
	"context"
	"main/internal/models"

	"github.com/google/uuid"
)

func (s service) PostRef(ctx context.Context, refData models.NewReferrer, userID int64) (models.Referrer, error) {
	id, _ := uuid.NewUUID()

	ref := models.Referrer{
		ID:   int64(id.ID()),
		UserID: userID,
		Referrer: refData.Referrer,
	}

	ref, err := s.repo.PostRef(ctx, ref)

	if err != nil {
		return models.Referrer{}, err
	}
	return ref, nil
}

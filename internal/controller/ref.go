package controller

import (
	"context"
	"main/internal/models"
)

func (c controller) PostRef(ctx context.Context, refData models.NewReferrer, userID int64) (models.Referrer, error) {
	return c.service.PostRef(ctx, refData, userID)
}

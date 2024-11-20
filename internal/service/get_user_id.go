package service

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (s service) GetUserID(ctx context.Context, params operations.GetUsersIDStatusParams) (models.User, error) {
	// token := params.HTTPRequest.Header.Get("Authorization")
	// if token == "" {
	// 	return models.User{}, errors.New("Unauthorized")
	// }

	// token = strings.TrimPrefix(token, "Bearer ")

	// err := jwt.ValidateToken(token, s.secret)
	// if err != nil {
	// 	return models.User{}, err
	// }

	user, err := s.repo.GetUserID(ctx, int(params.ID))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

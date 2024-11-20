package service

import (
	"context"
	"main/api/restapi/operations"
)

func (s service) DeleteUserID(ctx context.Context, params operations.DeleteUsersIDParams) error {
	// token := params.HTTPRequest.Header.Get("Authorization")
	// if token == "" {
	// 	return errors.New("Unauthorized")
	// }

	// token = strings.TrimPrefix(token, "Bearer ")

	// err := jwt.ValidateToken(token, s.secret)
	// if err != nil {
	// 	return err
	// }

	err := s.repo.DeleteUser(ctx, int(params.ID))
	if err != nil {
		return err
	}
	return nil
}

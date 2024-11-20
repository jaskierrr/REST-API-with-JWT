package service

import (
	"context"
)

func (s service) DeleteUserID(ctx context.Context, id int) error {
	_, err := s.repo.DeleteUser(ctx, id)
	return err
}

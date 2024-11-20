package service

import (
	"context"
	"errors"
	"main/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	id, _ := uuid.NewUUID()

	passHash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Join(errors.New("failed to generate password hash"), err)
		return models.User{}, err
	}

	user := models.User{
		ID:   int64(id.ID()),
		Email: userData.Email,
	}
	
	user, err = s.repo.PostUser(ctx, user, passHash)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

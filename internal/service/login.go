package service

import (
	"context"
	"main/internal/lib/jwt"
	"main/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s service) Login(ctx context.Context, userData models.NewUser) (string, error) {
	user, err := s.repo.GetUserEmail(ctx, userData.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		s.logger.Error("invalid credentials: " + err.Error())
		return "", err
	}

	token, err := jwt.NewToken(user, s.secret, time.Duration(s.tokenTTL) * time.Hour)
	if err != nil {
		s.logger.Error("failed to generate token" + err.Error())
		return "", err
	}

	return token, nil
}

package jwt

import (
	"fmt"
	"main/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user models.User, secret string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(duration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// slog.Info("NEW TOKEN tokenStr: " + tokenString)

	return tokenString, nil
}

func ValidateToken(tokenStr string, secret string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// slog.Info("VALIDATE TOKEN tokenStr: " + tokenStr)

		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("token invalid")
	}

	return nil
}

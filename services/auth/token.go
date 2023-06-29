package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jongseokleedev/sibsi-web-backend/server/responses"
	"time"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

func NewClaim(userID string) Claims {
	now := time.Now()
	claims := &Claims{

		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Minute * 15).Unix(),
		},
	}
	return *claims
}

func GenerateToken(claim Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))
}

func ValidateToken(token string, secret string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf(responses.ErrInvalidToken)
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	mapClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(responses.ErrInvalidToken)
	}

	return mapClaims, nil
}

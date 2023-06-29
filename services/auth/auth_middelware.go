package auth

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/responses"
	"net/http"
	"os"
	"strings"
)

type authenticationMiddleware struct {
	secret string
}

func NewAuthentication(secret string) *authenticationMiddleware {
	return &authenticationMiddleware{secret: secret}
}

var secret = os.Getenv("SECRET")

func TokenAuthMiddleware(c *gin.Context) {
	token, err := c.Request.Cookie("access-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed"})
		c.Abort()
		return
	}
	tknStr := token.Value
	if tknStr == "" {
		c.JSON(http.StatusUnauthorized,
			gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed"})
		c.Abort()
		return
	}
	claims := &Claims{}

	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized, "error": "token is expired"})
			c.Abort()
			return
		}
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden, "error": "Authentication failed"})
		c.Abort()
		return
	} else {
		c.Next()
	}
}

func (a *authenticationMiddleware) StripTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), responses.ErrStatusCode(err))
			return
		}

		claim, err := ValidateToken(token, a.secret)
		if err != nil {
			http.Error(w, err.Error(), responses.ErrStatusCode(err))
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claim["sub"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

const (
	BearerSchema string = "BEARER "
)

func getTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New(responses.ErrAuthorizationHeaderRequired)
	}

	bearerLength := len(BearerSchema)
	if len(authHeader) > bearerLength && strings.ToUpper(authHeader[0:bearerLength]) == BearerSchema {
		return authHeader[bearerLength:], nil
	}

	return "", errors.New(responses.ErrInvalidBearerScheme)
}

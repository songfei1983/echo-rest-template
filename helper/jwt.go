package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/songfei1983/go-api-server/internal/model"
	"strings"
	"time"
)

var DefaultJWTConfig = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: []byte("secret"),
	Skipper: func(c echo.Context) bool {
		isSkip := func(prefix string) bool { return strings.HasPrefix(c.Request().URL.Path, prefix) }
		for _, prefix := range []string{"/login", "/register"} {
			if isSkip(prefix) {
				return true
			}
		}
		return false
	},
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	model.User
	jwt.StandardClaims
}

func GenToken(u model.User) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		u,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

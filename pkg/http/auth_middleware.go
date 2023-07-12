package http

import (
	"net/http"
	"strings"

	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	token base.Token
}

func NewAuthMiddleware(token base.Token) *AuthMiddleware {
	return &AuthMiddleware{token}
}

func (m *AuthMiddleware) CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := m.getToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, base.Resp{Message: err.Error()})
		}

		user, err := m.token.Parse(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, base.Resp{Message: err.Error()})
		}

		c.Set("user", user)
		return next(c)
	}
}

func (m *AuthMiddleware) getToken(c echo.Context) (token string, err error) {
	a := c.Request().Header["Authorization"]
	if len(a) == 0 {
		return "", base.ErrTokenNotProvided
	}

	token = strings.Split(a[0], " ")[1]
	if token == "" {
		return "", base.ErrTokenNotProvided
	}

	return token, nil
}

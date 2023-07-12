package route

import (
	"github.com/Najwan160/go-experiment-3/code/auth/delivery/http/handler"
	"github.com/Najwan160/go-experiment-3/pkg/http"
	"github.com/labstack/echo/v4"
)

type AuthRouteParam struct {
	RegisterHandler *handler.RegisterHandler
	LoginHandler    *handler.LoginHandler
	AuthMiddleware  *http.AuthMiddleware
}

type AuthRoute struct {
	registerHandler *handler.RegisterHandler
	loginHandler    *handler.LoginHandler

	authMiddleware *http.AuthMiddleware
}

func NewAuthRoute(p AuthRouteParam) *AuthRoute {
	return &AuthRoute{
		registerHandler: p.RegisterHandler,
		loginHandler:    p.LoginHandler,
		authMiddleware:  p.AuthMiddleware,
	}
}

func (r *AuthRoute) Routes(e *echo.Group) {
	g := e.Group("/auth")

	g.POST("/register", r.registerHandler.Register)
	g.POST("/login", r.loginHandler.Login)

}

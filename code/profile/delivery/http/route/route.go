package route

import (
	"github.com/Najwan160/go-experiment-3/code/profile/delivery/http/handler"
	"github.com/Najwan160/go-experiment-3/pkg/http"
	"github.com/labstack/echo/v4"
)

type ProfileRouteParam struct {
	ProfileHandler *handler.ProfileHandler
	AuthMiddleware *http.AuthMiddleware
}

type ProfileRoute struct {
	profileHandler *handler.ProfileHandler
	AuthMiddleware *http.AuthMiddleware
}

func NewProfileRoute(p ProfileRouteParam) *ProfileRoute {
	return &ProfileRoute{
		profileHandler: p.ProfileHandler,
		AuthMiddleware: p.AuthMiddleware,
	}
}

func (r *ProfileRoute) Routes(e *echo.Group) {
	g := e.Group("/profile")

	g.GET("/:id", r.profileHandler.GetProfile, r.AuthMiddleware.CheckLogin)

}

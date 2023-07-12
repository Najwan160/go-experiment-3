package code

import (
	"github.com/labstack/echo/v4"

	"github.com/Najwan160/go-experiment-3/code/auth"
	auth_route "github.com/Najwan160/go-experiment-3/code/auth/delivery/http/route"
	"github.com/Najwan160/go-experiment-3/code/profile"
	profile_route "github.com/Najwan160/go-experiment-3/code/profile/delivery/http/route"
)

func Routes(e *echo.Echo) {
	g := e.Group("/api")

	auth_route.NewAuthRoute(auth_route.AuthRouteParam{
		RegisterHandler: auth.ProvideRegisterHandler(),
		LoginHandler:    auth.ProvideLoginHandler(),
	}).Routes(g)

	profile_route.NewProfileRoute(profile_route.ProfileRouteParam{
		ProfileHandler: profile.ProvideProfileHandler(),
		AuthMiddleware: auth.ProvideAuthMiddleware(),
	}).Routes(g)

}

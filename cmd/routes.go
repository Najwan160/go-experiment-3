package cmd

import (
	"github.com/Najwan160/go-experiment-3/code"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	code.Routes(e)
}

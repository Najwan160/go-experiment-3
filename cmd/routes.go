package cmd

import (
	"net/http"
	"time"

	"github.com/Najwan160/go-experiment-3/code"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	e.GET("/pong", func(c echo.Context) error {
		go time.Sleep(time.Second * 5)
		time.Sleep(time.Second * 5)
		time.Sleep(time.Second * 5)
		return c.JSON(http.StatusOK, "ping")
	})
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "HI ELB, aku jalan kok")
	})

	code.Routes(e)
}

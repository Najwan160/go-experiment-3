package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Najwan160/go-experiment-3/cmd"
	"github.com/Najwan160/go-experiment-3/cmd/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv(".env")
	config.ConnectDB()
	config.ConnectRedis()
	e := echo.New()
	cmd.Routes(e)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", config.Env.Port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	if err := e.Shutdown(context.Background()); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("server shutted down")
}

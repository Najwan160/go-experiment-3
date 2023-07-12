package handler

import (
	"net/http"

	"github.com/Najwan160/go-experiment-3/code/domain/auth"
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/labstack/echo/v4"
)

type LoginHandlerParam struct {
	LoginUC auth.LoginUsecase
}

type LoginHandler struct {
	loginUC auth.LoginUsecase
}

func NewLoginHandler(p LoginHandlerParam) *LoginHandler {
	return &LoginHandler{
		loginUC: p.LoginUC,
	}
}

func (h *LoginHandler) Login(c echo.Context) error {
	var req auth.LoginReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, base.Resp{Message: err.Error()})
	}

	resp, err := h.loginUC.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(base.GetStatusCode(err), base.GetRespErr(err))
	}

	return c.JSON(http.StatusOK, base.Resp{
		Message: "Berhasil login",
		Data:    resp,
	})
}

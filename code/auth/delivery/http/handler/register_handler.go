package handler

import (
	"net/http"

	"github.com/Najwan160/go-experiment-3/code/domain/auth"
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/labstack/echo/v4"
)

type RegisterHandlerParam struct {
	RegisterUC auth.RegisterUsecase
}

type RegisterHandler struct {
	registerUC auth.RegisterUsecase
}

func NewRegisterHandler(p RegisterHandlerParam) *RegisterHandler {
	return &RegisterHandler{
		registerUC: p.RegisterUC,
	}
}

func (h *RegisterHandler) Register(c echo.Context) error {
	var req auth.RegisterReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, base.Resp{Message: err.Error()})
	}

	resp, err := h.registerUC.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(base.GetStatusCode(err), base.GetRespErr(err))
	}

	return c.JSON(http.StatusCreated, base.Resp{
		Message: "register telah berhasil",
		Data:    resp,
	})
}

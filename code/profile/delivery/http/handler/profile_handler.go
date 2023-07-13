package handler

import (
	"net/http"

	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/Najwan160/go-experiment-3/code/domain/profile"
	"github.com/labstack/echo/v4"
)

type ProfileHandlerParam struct {
	ProfileUC profile.ProfileUsecase
}

type ProfileHandler struct {
	profileUC profile.ProfileUsecase
}

func NewProfileHandler(p ProfileHandlerParam) *ProfileHandler {
	return &ProfileHandler{
		profileUC: p.ProfileUC,
	}
}

func (h *ProfileHandler) GetProfile(c echo.Context) error {
	var req profile.GetProfileReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, base.Resp{Message: err.Error()})
	}

	resp, err := h.profileUC.GetProfile(c.Request().Context(), req)
	if err != nil {
		return c.JSON(base.GetStatusCode(err), base.GetRespErr(err))
	}

	return c.JSON(http.StatusOK, base.Resp{
		Message: "Berhasil Mendapatkan Profile",
		Data:    resp,
	})
}

func (h *ProfileHandler) GetProfiles(c echo.Context) error {
	var req profile.GetProfilesReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, base.Resp{Message: err.Error()})
	}

	resp, err := h.profileUC.GetProfiles(c.Request().Context(), req)
	if err != nil {
		return c.JSON(base.GetStatusCode(err), base.GetRespErr(err))
	}

	return c.JSON(http.StatusOK, base.Resp{
		Message: "Berhasil Mendapatkan Profile",
		Data:    resp,
	})
}

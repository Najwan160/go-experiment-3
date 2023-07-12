package profile

import (
	"context"
)

type GetProfileReq struct {
	ID string `param:"id"`
}

type GetProfileResp struct {
	Name string `json:"name"`
}

type ProfileUsecase interface {
	GetProfile(echo context.Context, req GetProfileReq) (resp GetProfileResp, err error)
}

package profile

import (
	"context"
)

type GetProfileReq struct {
	ID string `param:"id"`
}

type GetProfilesReq struct {
	Name string `query:"name"`
}

type GetProfilesResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetProfileResp struct {
	Name string `json:"name"`
}

type ProfileUsecase interface {
	GetProfile(echo context.Context, req GetProfileReq) (resp GetProfileResp, err error)
	GetProfiles(echo context.Context, req GetProfilesReq) (resp []GetProfilesResp, err error)
}

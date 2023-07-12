package auth

import (
	"context"

	"github.com/Najwan160/go-experiment-3/code/domain/entity"
)

type LoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	Account AccountResp `json:"account"`
	Token   string      `json:"token"`
}

func (r *LoginResp) SetAccount(acc entity.Account) {
	r.Account = parseAccount(acc)
}

type LoginUsecase interface {
	Login(ctx context.Context, req LoginReq) (LoginResp, error)
}

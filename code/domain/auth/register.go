package auth

import (
	"context"

	"github.com/Najwan160/go-experiment-3/code/domain/entity"
)

type RegisterReq struct {
	Name                 string  `json:"name" validate:"required"`
	Email                *string `json:"email" validate:"omitempty,email"`
	Password             string  `json:"password"`
	PasswordConfirmation string  `json:"password_confirmation"`
}
type RegisterResp struct {
	Account *AccountResp `json:"account"`
}

type RegisterRespWithToken struct {
	RegisterResp
	Token string `json:"token"`
}

func (r *RegisterResp) SetAccount(acc entity.Account) {
	a := parseAccount(acc)
	r.Account = &a
}

type RegisterUsecase interface {
	Register(ctx context.Context, req RegisterReq) (RegisterResp, error)
}

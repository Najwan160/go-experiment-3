package usecase

import (
	"context"
	"errors"

	"github.com/Najwan160/go-experiment-3/code/domain/auth"
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/Najwan160/go-experiment-3/code/domain/entity"
	"github.com/Najwan160/go-experiment-3/code/domain/model"
)

type LoginUsecase struct {
	db          base.DB
	validator   base.Validator
	hasher      base.Hasher
	accountRepo auth.AccountRepository
	token       base.Token
}

func NewLoginUsecase(
	db base.DB,
	validator base.Validator,
	hasher base.Hasher,
	accountRepo auth.AccountRepository,
	token base.Token,
) auth.LoginUsecase {
	return &LoginUsecase{
		db:          db,
		validator:   validator,
		hasher:      hasher,
		accountRepo: accountRepo,
		token:       token,
	}
}

func (uc *LoginUsecase) Login(ctx context.Context, req auth.LoginReq) (resp auth.LoginResp, err error) {
	if err = uc.validator.Validate(req); err != nil {
		return auth.LoginResp{}, err
	}

	acc, err := uc.accountRepo.Find(ctx, entity.Account{Email: &req.Email})
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		return auth.LoginResp{}, err
	}

	if uc.accountNotFound(acc, err) {
		return auth.LoginResp{}, base.ErrInvalidLogin
	}

	if f := uc.hasher.Verify(req.Password, acc.Password); !f {
		return auth.LoginResp{}, base.ErrInvalidLogin
	}

	resp.SetAccount(acc)

	token, err := uc.token.Generate(model.Auth{AccountID: acc.ID})
	if err != nil {
		return auth.LoginResp{}, err
	}

	resp.Token = token

	return resp, nil
}

func (uc *LoginUsecase) accountNotFound(acc entity.Account, err error) bool {
	return errors.Is(err, base.ErrNotFound) || acc.ID == ""
}

package usecase

import (
	"context"
	"errors"

	"github.com/Najwan160/go-experiment-3/code/domain/auth"
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/Najwan160/go-experiment-3/code/domain/entity"
)

type RegisterUsecase struct {
	db          base.DB
	validator   base.Validator
	uuid        base.UUID
	hasher      base.Hasher
	accountRepo auth.AccountRepository
	redis       base.Redis
}

func NewRegisterUsecase(
	db base.DB,
	validator base.Validator,
	uuid base.UUID,
	hasher base.Hasher,
	accountRepo auth.AccountRepository,
	redis base.Redis,
) auth.RegisterUsecase {
	return &RegisterUsecase{
		db:          db,
		validator:   validator,
		uuid:        uuid,
		hasher:      hasher,
		accountRepo: accountRepo,
		redis:       redis,
	}
}

func (uc *RegisterUsecase) Register(ctx context.Context, req auth.RegisterReq) (resp auth.RegisterResp, err error) {
	if err = uc.validator.Validate(req); err != nil {
		return auth.RegisterResp{}, err
	}

	trx := uc.db.Begin()
	ctx = context.WithValue(ctx, base.TrxContextKey, trx)
	defer func() {
		if err != nil {
			trx.Rollback()
		} else {
			trx.Commit()
		}
	}()

	acc, err := uc.createAccount(ctx, req)
	if err != nil {
		return auth.RegisterResp{}, err
	}
	resp.SetAccount(acc)

	uc.redis.SetKey(ctx, acc.ID, acc.Name)

	return resp, nil
}

func (uc *RegisterUsecase) createAccount(ctx context.Context, req auth.RegisterReq) (entity.Account, error) {
	if req.Password != req.PasswordConfirmation {
		return entity.Account{}, base.ErrMissmatchPasswordConfirmation
	}

	_, err := uc.accountRepo.Find(ctx, entity.Account{Email: req.Email})
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		return entity.Account{}, err
	}

	if err == nil {
		return entity.Account{}, base.ErrEmailConflict
	}

	hashedPassword, err := uc.hasher.Hash([]byte(req.Password), 10)
	if err != nil {
		return entity.Account{}, err
	}

	acc := entity.Account{
		ID:       uc.uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	_, err = uc.accountRepo.Create(ctx, acc)

	return acc, err
}

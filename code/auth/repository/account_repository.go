package repository

import (
	"context"
	"errors"

	"github.com/Najwan160/go-experiment-3/code/domain/auth"
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/Najwan160/go-experiment-3/code/domain/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db base.DB
}

func NewAccountRepository(db base.DB) auth.AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (repo *AccountRepository) Find(ctx context.Context, filter entity.Account) (res entity.Account, err error) {
	err = repo.db.FromContext(ctx).(*gorm.DB).
		WithContext(ctx).
		Where(filter).
		Take(&res).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = base.ErrNotFound
	}

	return res, err
}

func (repo *AccountRepository) Get(ctx context.Context, acc entity.Account) (res []entity.Account, err error) {
	err = repo.db.FromContext(ctx).(*gorm.DB).
		WithContext(ctx).
		Where(acc).
		Find(&res).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = base.ErrNotFound
	}

	return res, err
}
func (repo *AccountRepository) Create(ctx context.Context, acc entity.Account) (res entity.Account, err error) {
	err = repo.db.FromContext(ctx).(*gorm.DB).
		WithContext(ctx).
		Create(&acc).
		Error

	return acc, err
}

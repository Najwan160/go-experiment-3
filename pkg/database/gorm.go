package database

import (
	"context"

	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"gorm.io/gorm"
)

type Transactioner struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) base.DB {
	return &Transactioner{db}
}

func (trx *Transactioner) Begin() base.DB {
	return &Transactioner{
		trx.DB.Begin(),
	}
}

func (trx *Transactioner) Commit() {
	trx.DB.Commit()
	trx = nil
}

func (trx *Transactioner) SavePoint(name string) {
	trx.DB.SavePoint(name)
}

func (trx *Transactioner) Rollback() {
	trx.DB.Rollback()
}

func (trx *Transactioner) FromContext(ctx context.Context) interface{} {
	val := ctx.Value(base.TrxContextKey)
	if val == nil {
		return trx.DB
	}

	fromContext, ok := val.(*Transactioner)
	if !ok {
		return trx.DB
	}

	return fromContext.DB
}

func (trx *Transactioner) FromContextWithTrashed(ctx context.Context) interface{} {
	val := ctx.Value(base.TrxContextKey)
	if val == nil {
		return trx.DB.Unscoped()
	}

	fromContext, ok := val.(*Transactioner)
	if !ok {
		return trx.DB.Unscoped()
	}

	return fromContext.DB.Unscoped()
}

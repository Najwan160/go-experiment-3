package base

import "context"

type ctxKey string

var (
	TrxContextKey = ctxKey("transaction")
)

type DB interface {
	Begin() DB
	Commit()
	Rollback()
	FromContext(context.Context) interface{}
	FromContextWithTrashed(context.Context) interface{}
}

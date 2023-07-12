package base

import (
	"context"
)

type Redis interface {
	SetKey(ctx context.Context, key string, value interface{})
	GetKey(ctx context.Context, key string) (value string)
}

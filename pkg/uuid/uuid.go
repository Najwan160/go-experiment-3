package uuid

import (
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/google/uuid"
)

type CustomUUID struct {
}

func NewCustomUUID() base.UUID {
	return &CustomUUID{}
}

func (*CustomUUID) New() string {
	return uuid.NewString()
}

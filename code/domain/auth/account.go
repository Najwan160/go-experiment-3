package auth

import (
	"context"

	"github.com/Najwan160/go-experiment-3/code/domain/entity"
)

type AccountResp struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`
}

func parseAccount(acc entity.Account) AccountResp {
	return AccountResp{
		ID:    acc.ID,
		Name:  acc.Name,
		Email: acc.Email,
	}
}

type AccountRepository interface {
	Find(ctx context.Context, filter entity.Account) (entity.Account, error)
	Create(ctx context.Context, acc entity.Account) (entity.Account, error)
	Get(ctx context.Context, acc entity.Account) ([]entity.Account, error)
}

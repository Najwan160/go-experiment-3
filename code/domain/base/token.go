package base

import "github.com/Najwan160/go-experiment-3/code/domain/model"

var (
	TokenRoleAccount = "account"
)

type Token interface {
	Generate(model.Auth) (string, error)
	Parse(string) (model.Auth, error)
}

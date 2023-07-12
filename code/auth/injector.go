package auth

import (
	"github.com/Najwan160/go-experiment-3/code/auth/delivery/http/handler"
	"github.com/Najwan160/go-experiment-3/code/auth/repository"
	"github.com/Najwan160/go-experiment-3/code/auth/usecase"

	"github.com/Najwan160/go-experiment-3/cmd/config"
	"github.com/Najwan160/go-experiment-3/pkg/database"
	"github.com/Najwan160/go-experiment-3/pkg/hash"
	"github.com/Najwan160/go-experiment-3/pkg/http"
	"github.com/Najwan160/go-experiment-3/pkg/redis"
	"github.com/Najwan160/go-experiment-3/pkg/token"
	"github.com/Najwan160/go-experiment-3/pkg/uuid"
	"github.com/Najwan160/go-experiment-3/pkg/validator"
)

func ProvideRegisterHandler() *handler.RegisterHandler {
	validator := validator.NewCustomValidator()
	db := database.NewDB(config.DB)
	uuid := uuid.NewCustomUUID()
	hasher := hash.NewHasher()
	redis := redis.NewRedis(config.Redis)

	accountRepo := repository.NewAccountRepository(db)

	registerUC := usecase.NewRegisterUsecase(db, validator, uuid, hasher, accountRepo, redis)

	return handler.NewRegisterHandler(handler.RegisterHandlerParam{
		RegisterUC: registerUC,
	})
}

func ProvideLoginHandler() *handler.LoginHandler {
	validator := validator.NewCustomValidator()
	db := database.NewDB(config.DB)
	hasher := hash.NewHasher()
	token := token.NewJWT(config.Env.JWT.TTL, config.Env.JWT.SignatureKey)

	accountRepo := repository.NewAccountRepository(db)

	loginUC := usecase.NewLoginUsecase(db, validator, hasher, accountRepo, token)

	return handler.NewLoginHandler(handler.LoginHandlerParam{
		LoginUC: loginUC,
	})
}

func ProvideAuthMiddleware() *http.AuthMiddleware {
	token := token.NewJWT(config.Env.JWT.TTL, config.Env.JWT.SignatureKey)

	return http.NewAuthMiddleware(token)
}

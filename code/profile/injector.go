package profile

import (
	"github.com/Najwan160/go-experiment-3/code/auth/repository"
	"github.com/Najwan160/go-experiment-3/code/profile/delivery/http/handler"
	"github.com/Najwan160/go-experiment-3/code/profile/usecase"
	"github.com/Najwan160/go-experiment-3/pkg/redis"

	"github.com/Najwan160/go-experiment-3/cmd/config"
	"github.com/Najwan160/go-experiment-3/pkg/database"
	"github.com/Najwan160/go-experiment-3/pkg/validator"
)

func ProvideProfileHandler() *handler.ProfileHandler {
	validator := validator.NewCustomValidator()
	db := database.NewDB(config.DB)
	redis := redis.NewRedis(config.Redis)

	accountRepo := repository.NewAccountRepository(db)

	profileUC := usecase.NewProfileUsecase(db, validator, accountRepo, redis)

	return handler.NewProfileHandler(handler.ProfileHandlerParam{
		ProfileUC: profileUC,
	})
}

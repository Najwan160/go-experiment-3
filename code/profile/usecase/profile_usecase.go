package usecase

import (
	"context"

	"github.com/Najwan160/go-experiment-3/code/domain/auth"
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/Najwan160/go-experiment-3/code/domain/entity"
	"github.com/Najwan160/go-experiment-3/code/domain/profile"
)

type ProfileUsecase struct {
	db          base.DB
	validator   base.Validator
	accountRepo auth.AccountRepository
	redis       base.Redis
}

func NewProfileUsecase(
	db base.DB,
	validator base.Validator,
	accountRepo auth.AccountRepository,
	redis base.Redis,
) profile.ProfileUsecase {
	return &ProfileUsecase{
		db:          db,
		validator:   validator,
		accountRepo: accountRepo,
		redis:       redis,
	}
}

func (uc *ProfileUsecase) GetProfile(ctx context.Context, req profile.GetProfileReq) (resp profile.GetProfileResp, err error) {
	if err = uc.validator.Validate(req); err != nil {
		return profile.GetProfileResp{}, err
	}

	// filter := entity.Account{
	// 	ID: req.ID,
	// }

	value := uc.redis.GetKey(ctx, req.ID)
	if value == "" {
		return profile.GetProfileResp{}, base.ErrNotFound
	}
	response := profile.GetProfileResp{
		Name: value,
	}

	// uc.kafka.SendMessage(fmt.Sprintf("%v mencoba melihat profile", value))

	// acc, err := uc.accountRepo.Find(ctx, filter)
	// if err != nil {
	// 	return profile.GetProfileResp{}, err
	// }

	// response := profile.GetProfileResp{
	// 	Name: acc.Name,
	// }

	return response, nil
}

func (uc *ProfileUsecase) GetProfiles(ctx context.Context, req profile.GetProfilesReq) (resp []profile.GetProfilesResp, err error) {
	var res []profile.GetProfilesResp

	if err = uc.validator.Validate(req); err != nil {
		return []profile.GetProfilesResp{}, err
	}

	filter := entity.Account{
		Name: req.Name,
	}

	acc, err := uc.accountRepo.Get(ctx, filter)
	if err != nil {
		return []profile.GetProfilesResp{}, err
	}

	for _, v := range acc {
		res = append(res, profile.GetProfilesResp{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return res, nil
}

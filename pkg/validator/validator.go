package validator

import (
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var (
	uni *ut.UniversalTranslator
)

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func NewCustomValidator() base.Validator {
	validator := validator.New()
	language := id.New()
	uni = ut.New(language, language)
	trans, _ := uni.GetTranslator("language")
	id_translations.RegisterDefaultTranslations(validator, trans)
	return &CustomValidator{validator, trans}
}

func (v *CustomValidator) Validate(s interface{}) error {
	if err := v.validator.Struct(s); err != nil {
		result := make(map[string]string)
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			result[e.Field()] = e.Translate(v.trans)
			// can translate each error one at a time.
		}

		return &base.ValidationError{
			Err:     base.ErrInvalidRequest,
			ErrData: result,
		}
	}

	return nil
}

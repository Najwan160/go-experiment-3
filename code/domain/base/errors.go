package base

import (
	"errors"
	"net/http"
)

var (
	ErrMissmatchPasswordConfirmation = errors.New("missmatch password and password confirmation")
	ErrInvalidTransaction            = errors.New("invalid db transaction")
	ErrNotFound                      = errors.New("your requested item is not found")
	ErrEmailConflict                 = errors.New("phone number already exists")
	ErrFailedReferalCode             = errors.New("failed to generate referal code")
	ErrMissingRequiredAttributes     = errors.New("missing required attribute")
	ErrInvalidRequest                = errors.New("invalid request")
	ErrInvalidLogin                  = errors.New("phone number or password invalid")
	ErrTokenNotProvided              = errors.New("token not provided")
	ErrWrongPassword                 = errors.New("wrong password")
	ErrInvalidToken                  = errors.New("invalid token")

	//setting
	ErrDuplicateData = errors.New("cannot duplicate data")
)

type ValidationError struct {
	Err     error
	ErrData map[string]string
}

func (e *ValidationError) Error() string {
	return e.Err.Error()
}

func GetRealErr(err error) error {
	var validationErr *ValidationError
	// if error from validator
	if errors.As(err, &validationErr) {
		return validationErr.Err
	}

	return err
}

func GetStatusCode(err error) int {
	realErr := GetRealErr(err)
	switch realErr {
	case ErrMissmatchPasswordConfirmation, ErrMissingRequiredAttributes, ErrDuplicateData:
		return http.StatusBadRequest
	case ErrEmailConflict:
		return http.StatusConflict
	case ErrInvalidLogin:
		return http.StatusUnauthorized
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func GetRespErr(err error) interface{} { //GetErrResp
	var validationErr *ValidationError

	realErr := GetRealErr(err)

	if errors.As(err, &validationErr) {
		return RespErrValidator{
			Message: GetErrorMessage(realErr, localeID),
			ErrData: validationErr.ErrData,
			Data:    nil,
		}
	} else {
		return RespErr{
			Message: GetErrorMessage(realErr, localeID),
			Data:    nil,
		}
	}
}

package base

type locale string

const (
	localeID locale = "id"
	localeEN locale = "en"
)

func GetErrorMessage(err error, locale locale) string {
	if locale == localeID {
		return GetIndonesianErrorMessage(err)
	}
	return err.Error()
}

func GetIndonesianErrorMessage(err error) string {
	switch err {
	case ErrMissmatchPasswordConfirmation:
		return "Password dan konfirmasi password tidak sama"
	case ErrInvalidTransaction:
		return "Terdapat kesalahan pada sistem, silahkan coba lagi [invalid transaction]"
	case ErrNotFound:
		return "Data tidak ditemukan"
	case ErrEmailConflict:
		return "Email anda sudah terdaftar, silahkan langsung login pada aplikasi"
	case ErrMissingRequiredAttributes:
		return "Pastikan semua form sudah terisi dengan lengkap"
	case ErrInvalidRequest:
		return "Pastikan semua form sudah terisi dengan benar"
	case ErrInvalidLogin:
		return "Email atau password tidak cocok"
	case ErrTokenNotProvided:
		return "Token not provided"
	case ErrWrongPassword:
		return "Password tidak cocok"
	case ErrInvalidToken:
		return "Invalid token"
	case ErrDuplicateData:
		return "Data tidak bisa diduplikasi"
	default:
		return "Terdapat kesalahan pada sistem, silahkan coba lagi"
	}
}

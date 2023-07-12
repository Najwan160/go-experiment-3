package base

type Validator interface {
	Validate(s interface{}) error
}

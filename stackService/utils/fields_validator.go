package utils

type FieldsValidator interface {
	CheckRequiredFieldNotEmpty() error
}

func Validate(val FieldsValidator) error {
	if err := val.CheckRequiredFieldNotEmpty(); err != nil {
		return err
	}
	return nil
}

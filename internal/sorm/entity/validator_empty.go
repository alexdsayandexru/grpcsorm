package entity

func NewEmptyValidator() IValidator {
	return &EmptyValidator{}
}

type EmptyValidator struct {
}

func (v *EmptyValidator) Validate(IEntity) (bool, ErrorInfo) {
	return true, ErrorInfo{FieldErrors: nil}
}

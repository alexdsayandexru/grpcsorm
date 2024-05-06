package entity

func NewValidator() IValidator {
	return &Validator{}
}

type Validator struct {
}

func (v *Validator) Validate(entity IEntity) (bool, ErrorInfo) {
	var errors []FieldError

	for fieldName, rule := range entity.GetRules() {
		ok, err := rule()
		if !ok {
			errors = append(errors, FieldError{FieldName: fieldName, ErrorDescription: err.Error()})
		}
	}
	return len(errors) == 0, ErrorInfo{
		ErrorsCode:        "field-errors",
		ErrorsDescription: "Ошибка валидации данных",
		FieldErrors:       errors,
	}
}

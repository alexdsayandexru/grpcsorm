package validator

import (
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
)

type Validator struct {
	target interface{}
	error  error
}

func Validate(target interface{}) *Validator {
	return &Validator{target: target}
}

func (v *Validator) GetResult() (bool, error) {
	return v.error == nil, v.error
}

func (v *Validator) Required() *Validator {
	if v.error != nil {
		return v
	}
	if arr, ok := v.target.([]string); ok && arr == nil {
		v.error = errors.ValidationErrorRequiredField()
	}
	if str, ok := v.target.(string); ok && len(str) == 0 {
		v.error = errors.ValidationErrorRequiredField()
	}
	if num, ok := v.target.(int32); ok && num == 0 {
		v.error = errors.ValidationErrorRequiredField()
	}
	return v
}

func (v *Validator) RequiredIf(predicate bool) *Validator {
	if v.error != nil {
		return v
	}
	if predicate {
		return v.Required()
	}
	return v
}

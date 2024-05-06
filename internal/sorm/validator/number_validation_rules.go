package validator

import (
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
)

func (v *Validator) Maximum(value int32) *Validator {
	if v.error != nil {
		return v
	}
	if target, ok := v.target.(int32); ok {
		if target > value {
			v.error = errors.ValidationErrorMaxValue(value)
		}
	}
	return v
}

func (v *Validator) Equal(value int32) *Validator {
	if v.error != nil {
		return v
	}
	if target, ok := v.target.(int32); ok {
		if target > 0 && target != value {
			v.error = errors.ValidationErrorEqualValue(value)
		}
	}
	return v
}

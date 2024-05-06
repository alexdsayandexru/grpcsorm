package validator

import (
	"github.com/google/uuid"
)

func (v *Validator) Uiid() *Validator {
	if v.error != nil {
		return v
	}
	if value, ok := v.target.(string); ok {
		if _, err := uuid.Parse(value); err != nil {
			v.error = err
		}
	}
	return v
}

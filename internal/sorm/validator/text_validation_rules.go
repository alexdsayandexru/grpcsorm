package validator

import (
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
	"regexp"
	"strings"
)

func (v *Validator) MaxLength(max int) *Validator {
	if v.error != nil {
		return v
	}
	if value, ok := v.target.(string); ok {
		if len(value) > 0 && strings.Count(value, "") > max+1 {
			v.error = errors.ValidationErrorMaxFieldLength(max)
		}
	}
	return v
}

func (v *Validator) Length(count int) *Validator {
	if v.error != nil {
		return v
	}
	if value, ok := v.target.(string); ok {
		if len(value) > 0 && strings.Count(value, "") != count+1 {
			v.error = errors.ValidationErrorEqualFieldLength(count)
		}
	}
	return v
}

func (v *Validator) Regex(expression string) *Validator {
	if v.error != nil {
		return v
	}
	if value, ok := v.target.(string); ok {
		if len(value) > 0 {
			if re := regexp.MustCompile(expression); !re.MatchString(value) {
				v.error = errors.ValidationErrorInvalidValueOrSimbols()
			}
		}
	}
	return v
}

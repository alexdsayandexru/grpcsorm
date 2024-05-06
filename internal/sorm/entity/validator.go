package entity

import (
	"encoding/json"
)

type ValidationRules map[string]func() (bool, error)

type FieldError struct {
	FieldName        string `json:"field_name"`
	ErrorDescription string `json:"error_description"`
}

type ErrorInfo struct {
	ErrorsCode        string       `json:"errors_code"`
	ErrorsDescription string       `json:"errors_description"`
	FieldErrors       []FieldError `json:"field_errors"`
}

func (e *ErrorInfo) ToBytes() []byte {
	bites, _ := json.Marshal(*e)
	return bites
}

func (e *ErrorInfo) ToJson() string {
	return string(e.ToBytes())
}

type IValidator interface {
	Validate(entity IEntity) (bool, ErrorInfo)
}

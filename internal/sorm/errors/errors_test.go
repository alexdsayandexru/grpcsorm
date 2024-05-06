package errors

import (
	"errors"
	"testing"
)

func Test_Errors(t *testing.T) {
	_ = NotFoundFileCACertificateError("")
	_ = NotFoundCACertificateError()
	_ = InvalidApiKeyError()
	_ = NotFoundKeyInEnvironmentError("")
	_ = EmptyContextMetadataError()
	_ = ValidationError(0, "", errors.New(""))
	_ = ValidationErrorMaxValue(0)
	_ = ValidationErrorEqualValue(0)
	_ = ValidationErrorInvalidValueOrSimbols()
	_ = ValidationErrorEqualFieldLength(0)
	_ = ValidationErrorMaxFieldLength(0)
	_ = ValidationErrorRequiredField()
}

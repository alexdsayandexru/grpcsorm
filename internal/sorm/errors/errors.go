package errors

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	InvalidValue                string = "Поле содержит некорректное значение или недопустимые символы"
	RequiredField                      = "Поле обязательно для заполнения"
	MaxFieldLength                     = "Максимальная длина поля - %d символов"
	EqualFieldLength                   = "Длина поля должна быть %d символов"
	MaxValue                           = "Поле должно быть меньше или равно %d"
	EqualValue                         = "Поле должно быть равно %d"
	ValidationIndex                    = "Ошибка валидации: [%d].[%s].[%s]"
	EmptyContextMetadata               = "В методанных контекста не обнаружено ни одного ключа"
	NotFoundApiKeyInEnvironment        = "Не найден [%s] в переменных окружения"
	InvalidApiKey                      = "Недопустимое значение X_API_KEY"
	NotFoundCACertificate              = "Не найден СА сертификат для подключения к Kafka"
	NotFoundFileCACertificate          = "Не возможно загрузить СА сертификат:[%s]"
)

func NotFoundFileCACertificateError(err string) error {
	return status.Error(codes.Aborted, fmt.Sprintf(NotFoundFileCACertificate, err))
}

func NotFoundCACertificateError() error {
	return status.Error(codes.Aborted, NotFoundCACertificate)
}

func InvalidApiKeyError() error {
	return status.Error(codes.Unauthenticated, InvalidApiKey)
}

func NotFoundKeyInEnvironmentError(name string) error {
	return status.Error(codes.Unauthenticated, fmt.Sprintf(NotFoundApiKeyInEnvironment, name))
}

func EmptyContextMetadataError() error {
	return status.Error(codes.Unauthenticated, EmptyContextMetadata)
}

func ValidationError(index int, field string, err error) error {
	return errors.New(fmt.Sprintf(ValidationIndex, index, field, err.Error()))
}

func ValidationErrorMaxValue(value int32) error {
	return errors.New(fmt.Sprintf(MaxValue, value))
}

func ValidationErrorEqualValue(value int32) error {
	return errors.New(fmt.Sprintf(EqualValue, value))
}

func ValidationErrorInvalidValueOrSimbols() error {
	return errors.New(InvalidValue)
}

func ValidationErrorEqualFieldLength(count int) error {
	return errors.New(fmt.Sprintf(EqualFieldLength, count))
}

func ValidationErrorMaxFieldLength(max int) error {
	return errors.New(fmt.Sprintf(MaxFieldLength, max))
}

func ValidationErrorRequiredField() error {
	return errors.New(RequiredField)
}

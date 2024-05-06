package entity

import (
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/validator"
)

func ValidateMsisdnLogin(msisdnLogin string, emailLogin string) (bool, error) {
	return validator.Validate(msisdnLogin).RequiredIf(len(emailLogin) == 0).MaxLength(16).Regex("^[0-9]+$").GetResult()
}

func ValidateEmailLogin(emailLogin string, msisdnLogin string) (bool, error) {
	return validator.Validate(emailLogin).RequiredIf(len(msisdnLogin) == 0).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
}

func ValidateFactor(factor string) (bool, error) {
	return validator.Validate(factor).MaxLength(255).Regex("^[A-Za-z0-9-]+$").GetResult()
}

func ValidateIp(ip string) (bool, error) {
	return validator.Validate(ip).MaxLength(255).GetResult()
}

func ValidatePort(port int32) (bool, error) {
	return validator.Validate(port).Maximum(99999).GetResult()
}

func ValidateUserAgent(userAgent string) (bool, error) {
	return validator.Validate(userAgent).MaxLength(1023).GetResult()
}

func ValidateMessage(message string) (bool, error) {
	return validator.Validate(message).Required().MaxLength(8000).GetResult()
}

func ValidateCorrelationId(correlationId string) (bool, error) {
	return validator.Validate(correlationId).Required().Uiid().GetResult()
}

func ValidateTelcoId(telcoId int32) (bool, error) {
	return validator.Validate(telcoId).Required().Maximum(100).GetResult()
}

func ValidateUserType(userType int32) (bool, error) {
	return validator.Validate(userType).Required().Maximum(100).GetResult()
}

func ValidateUserId(userId string) (bool, error) {
	return validator.Validate(userId).Required().MaxLength(255).GetResult()
}

func ValidateNick(nick string) (bool, error) {
	return validator.Validate(nick).MaxLength(150).GetResult()
}

func ValidateBirthDate(birthDate string) (bool, error) {
	return validator.Validate(birthDate).Length(10).Regex("^[0-9-]+$").GetResult()
}

func ValidateNameFamilyInitial(name string) (bool, error) {
	return validator.Validate(name).MaxLength(30).Regex("^[A-Za-zА-Яа-яЁё -]+$").GetResult()
}

func ValidateAddress(address string) (bool, error) {
	return validator.Validate(address).MaxLength(255).GetResult()
}

func ValidateRequiredDatetime(datetime string) (bool, error) {
	return validator.Validate(datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
}

func ValidateNotRequiredDatetime(datetime string) (bool, error) {
	return validator.Validate(datetime).Length(23).Regex("^[0-9 :.-]+$").GetResult()
}

func ValidateServiceUser(serviceUser int32) (bool, error) {
	return validator.Validate(serviceUser).Required().Equal(1).GetResult()
}

func ValidateServiceId(serviceId int32) (bool, error) {
	return validator.Validate(serviceId).Required().Maximum(2147483647).GetResult()
}

func ValidateAdditional(additional []Additional) (bool, error) {
	if additional == nil {
		return true, nil
	}
	for i, m := range additional {
		if ok, err := validator.Validate(m.Title).MaxLength(255).Regex("^[A-Za-zА-Яа-яЁё ,_-]+$").GetResult(); !ok {
			return false, errors.ValidationError(i, "title", err)
		}
		if ok, err := validator.Validate(m.Content).MaxLength(255).GetResult(); !ok {
			return false, errors.ValidationError(i, "content", err)
		}
	}
	return true, nil
}

func ValidateMsisdns(msisdn []string, emails []string) (bool, error) {
	ok, err := validator.Validate(msisdn).RequiredIf(emails == nil || len(emails) == 0).GetResult()
	if ok {
		for i, m := range msisdn {
			if ok, err = validator.Validate(m).MaxLength(16).Regex("^[0-9]+$").GetResult(); !ok {
				return false, errors.ValidationError(i, "msisdn", err)
			}
		}
	}
	return ok, err
}

func ValidateEmails(emails []string, msisdn []string) (bool, error) {
	ok, err := validator.Validate(emails).RequiredIf(msisdn == nil || len(msisdn) == 0).GetResult()
	if ok {
		for i, m := range emails {
			if ok, err = validator.Validate(m).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult(); !ok {
				return false, errors.ValidationError(i, "email", err)
			}
		}
	}
	return ok, err
}

func ValidateImId(imIds []ImId) (bool, error) {
	if imIds != nil {
		for i, m := range imIds {
			if ok, err := validator.Validate(m.ServiceId).MaxLength(255).GetResult(); !ok {
				return false, errors.ValidationError(i, "service_id", err)
			}
			if ok, err := validator.Validate(m.ServiceName).MaxLength(255).GetResult(); !ok {
				return false, errors.ValidationError(i, "service_name", err)
			}
		}
	}
	return true, nil
}

func ValidateOriServices(oriServices []Service) (bool, error) {
	if oriServices != nil {
		for i, s := range oriServices {
			if ok, err := validator.Validate(s.ServiceId).Required().Maximum(2147483647).GetResult(); !ok {
				return false, errors.ValidationError(i, "service_id", err)
			}
			if ok, err := validator.Validate(s.Description).MaxLength(255).GetResult(); !ok {
				return false, errors.ValidationError(i, "description", err)
			}
			if ok, err := validator.Validate(s.BeginTime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult(); !ok {
				return false, errors.ValidationError(i, "begin_time", err)
			}
			if ok, err := validator.Validate(s.EndTime).Length(23).Regex("^[0-9 :.-]+$").GetResult(); !ok {
				return false, errors.ValidationError(i, "end_time", err)
			}
		}
	}
	return true, nil
}

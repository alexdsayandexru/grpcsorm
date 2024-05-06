package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewRegisterUser(s *sorm.RegisterUserRequest) *entity.RegisterUser {
	return &entity.RegisterUser{
		EventType:     "user_registration",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: entity.AbInfo{
			UserId:       s.UserId,
			Msisdns:      PhonesToStrings(s.Msisdns),
			Emails:       EmailsToStrings(s.Emails),
			DatetimeReg:  s.DatetimeReg,
			ServiceId:    s.ServiceUser,
			Additional:   AddsToAdditionals(s.Additional),
			ContractDate: s.ContractDate,
		},
		EventData: entity.RegisterUserEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Msisdn:    s.MsisdnLogin,
			Email:     s.EmailLogin,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
			Message:   s.Message,
		},
	}
}

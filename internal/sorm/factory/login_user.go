package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewLoginUser(s *sorm.LoginUserRequest) *entity.LoginUser {
	return &entity.LoginUser{
		EventType:     "user_login",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		Event: entity.LoginUserEvent{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Msisdn:    s.MsisdnLogin,
			Email:     s.EmailLogin,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Factor:    s.Factor,
			Datetime:  s.Datetime,
		},
	}
}

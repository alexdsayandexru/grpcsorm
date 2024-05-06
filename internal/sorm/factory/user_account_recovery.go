package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewUserAccountRecovery(s *sorm.UserAccountRecoveryRequest) *entity.UserAccountRecovery {
	return &entity.UserAccountRecovery{
		EventType:     "user_account_recovery",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		EventData: entity.UserAccountRecoveryEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Msisdn:    s.MsisdnLogin,
			Email:     s.EmailLogin,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
		},
	}
}

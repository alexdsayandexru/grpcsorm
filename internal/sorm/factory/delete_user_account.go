package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewDeleteUserAccount(s *sorm.DeleteUserAccountRequest) *entity.DeleteUserAccount {
	return &entity.DeleteUserAccount{
		EventType:     "delete_user_account",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		EventData: entity.DeleteUserAccountEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
			Message:   s.Message,
		},
	}
}

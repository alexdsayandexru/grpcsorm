package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewUpdateUserData(s *sorm.UpdateUserDataRequest) *entity.UpdateUserData {
	return &entity.UpdateUserData{
		EventType:     "user_update_data",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: entity.UpdateUserDataAbInfo{
			UserId:          s.UserId,
			Nick:            s.Nick,
			BirthDate:       s.BirthDate,
			Name:            s.Name,
			Family:          s.Family,
			Initial:         s.Initial,
			Msisdns:         PhonesToStrings(s.Msisdns),
			Emails:          EmailsToStrings(s.Emails),
			Address:         s.Address,
			DatetimeReg:     s.DatetimeReg,
			DatetimeUpdated: s.DatetimeUpdated,
			ServiceId:       s.ServiceUser,
			ImId:            ImIdsToImIds(s.ImIds),
			Additional:      AddsToAdditionals(s.Additional),
			ContractDate:    s.ContractDate,
		},
		EventData: entity.UpdateUserDataEventData{
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

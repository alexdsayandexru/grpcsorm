package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewUpdateUserDataAdmin(s *sorm.UpdateUserDataAdminRequest) *entity.UpdateUserDataAdmin {
	return &entity.UpdateUserDataAdmin{
		EventType:     "user_update_admin",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: entity.UpdateUserDataAdminAbInfo{
			UserId:          s.UserId,
			Msisdns:         PhonesToStrings(s.Msisdns),
			Emails:          EmailsToStrings(s.Emails),
			Nick:            s.Nick,
			BirthDate:       s.BirthDate,
			Name:            s.Name,
			Family:          s.Family,
			Initial:         s.Initial,
			Address:         s.Address,
			DatetimeReg:     s.DatetimeReg,
			DatetimeUpdated: s.DatetimeUpdated,
			ServiceId:       s.ServiceUser,
			ImId:            ImIdsToImIds(s.ImIds),
			Additional:      AddsToAdditionals(s.Additional),
			ContractDate:    s.ContractDate,
		},
		EventData: entity.UpdateUserDataAdminEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Datetime:  s.Datetime,
		},
	}
}

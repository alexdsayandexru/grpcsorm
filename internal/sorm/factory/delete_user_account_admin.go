package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewDeleteUserAccountAdmin(s *sorm.DeleteUserAccountAdminRequest) *entity.DeleteUserAccountAdmin {
	return &entity.DeleteUserAccountAdmin{
		EventType:     "user_delete_admin",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: entity.DeleteUserAccountAdminAbInfo{
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
			DatetimeUnreg:   s.DatetimeUnreg,
			ContractDate:    s.ContractDate,
		},
		EventData: entity.DeleteUserAccountAdminEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Datetime:  s.Datetime,
		},
	}
}

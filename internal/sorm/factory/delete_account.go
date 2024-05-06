package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewDeleteAccount(s *sorm.DeleteAccountRequest) *entity.DeleteAccount {
	return &entity.DeleteAccount{
		EventType:     "delete_account",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: entity.DeleteAccountAbInfo{
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
	}
}

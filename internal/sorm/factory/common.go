package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func PhonesToStrings(msisdns []*sorm.Msisdn) []string {
	var result []string
	for _, m := range msisdns {
		result = append(result, m.Msisdn)
	}
	return result
}

func EmailsToStrings(emails []*sorm.Email) []string {
	var result []string
	for _, m := range emails {
		result = append(result, m.Email)
	}
	return result
}

func AddsToAdditionals(adds []*sorm.Add) []entity.Additional {
	var result []entity.Additional
	for _, m := range adds {
		result = append(result, entity.Additional{
			Title:   m.Title,
			Content: m.Content,
		})
	}
	return result
}

func ImIdsToImIds(adds []*sorm.ImId) []entity.ImId {
	var result []entity.ImId
	for _, m := range adds {
		result = append(result, entity.ImId{
			ServiceId:   m.ServiceId,
			ServiceName: m.ServiceName,
		})
	}
	return result
}

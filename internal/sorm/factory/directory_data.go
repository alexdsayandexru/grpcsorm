package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
)

func NewDirectoryData(s *sorm.DirectoryDataRequest) *entity.DirectoryData {
	return &entity.DirectoryData{
		EventType:     "directory",
		CorrelationId: s.CorrelationId,
		Datetime:      s.Datetime,
		OriServices:   ServicesToServices(s.Services),
	}
}

func ServicesToServices(services []*sorm.Service) []entity.Service {
	var result []entity.Service
	for _, s := range services {
		result = append(result, entity.Service{
			ServiceId:   s.ServiceId,
			Description: s.Decription,
			BeginTime:   s.BeginTime,
			EndTime:     s.EndTime,
		})
	}
	return result
}

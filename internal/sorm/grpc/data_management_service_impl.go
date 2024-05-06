package grpc

import (
	"crypto/tls"
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type HealthCheckCallBackMethod func(code codes.Code)

type DataManagementServiceImpl struct {
	config       *configs.Config
	tls          *tls.Config
	healthCheck  HealthCheckCallBackMethod
	unitTestMode bool
	sorm.UnimplementedUserDataManagementServer
}

func NewDataManagementService(config *configs.Config, tls *tls.Config, healthCheck HealthCheckCallBackMethod, unitTestMode bool) *DataManagementServiceImpl {
	return &DataManagementServiceImpl{
		config:       config,
		tls:          tls,
		healthCheck:  healthCheck,
		unitTestMode: unitTestMode,
	}
}

func (s *DataManagementServiceImpl) RegisterGRPC(server *grpc.Server) {
	sorm.RegisterUserDataManagementServer(server, s)
}

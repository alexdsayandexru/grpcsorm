package grpc

import (
	"context"
	"github.com/google/uuid"
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/factory"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/kafka"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/security"
	"google.golang.org/grpc/codes"
)

func (s *DataManagementServiceImpl) Response(code codes.Code, message string, details string) *sorm.DataManagementResponse {
	s.healthCheck(code)
	return &sorm.DataManagementResponse{Code: int32(code), Message: message, Details: details}
}

func (s *DataManagementServiceImpl) GetProducer(topic string) kafka.IProducer {
	if s.config.EnabledDelivery {
		if s.config.EnabledTls {
			return kafka.NewProducerSaramaSingleton(s.config.KafkaBrokers, topic, s.tls, s.config.SaslUserName, s.config.SaslPassword)
		} else {
			return kafka.NewDefaultProducerSarama(s.config.KafkaBrokers, topic)
		}
	}
	return kafka.NewEmptyProducer()
}

func (s *DataManagementServiceImpl) GetValidator() entity.IValidator {
	if s.config.EnabledValidation {
		return entity.NewValidator()
	}
	return entity.NewEmptyValidator()
}

func (s *DataManagementServiceImpl) GetSecurity() security.ISecurity {
	if s.config.EnabledSecurity {
		return security.NewSecurity(s.config.XApiKey, s.config.XApiKeyValue)
	}
	return security.NewEmptySecurity()
}

func (s *DataManagementServiceImpl) Handle(ctx *context.Context, target entity.IEntity, topic string) (*sorm.DataManagementResponse, error) {
	code, message, details := NewDataManagementServiceEventHandler(
		ctx,
		target,
		s.GetValidator(),
		s.GetProducer(topic),
		s.GetSecurity(),
		uuid.New(),
		s.unitTestMode).Handle()

	return s.Response(code, message, details), nil
}

func (s *DataManagementServiceImpl) RegisterUser(ctx context.Context, request *sorm.RegisterUserRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewRegisterUser(request), s.config.TopicUserRegistration)
}

func (s *DataManagementServiceImpl) LoginUser(ctx context.Context, request *sorm.LoginUserRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewLoginUser(request), s.config.TopicUserLogin)
}

func (s *DataManagementServiceImpl) LogoutUser(ctx context.Context, request *sorm.LogoutUserRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewLogoutUser(request), s.config.TopicUserLogout)
}

func (s *DataManagementServiceImpl) DeleteUserAccount(ctx context.Context, request *sorm.DeleteUserAccountRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewDeleteUserAccount(request), s.config.TopicDeleteUserAccount)
}

func (s *DataManagementServiceImpl) UpdateUserData(ctx context.Context, request *sorm.UpdateUserDataRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewUpdateUserData(request), s.config.TopicUserUpdateData)
}

func (s *DataManagementServiceImpl) DeleteUserAccountAdmin(ctx context.Context, request *sorm.DeleteUserAccountAdminRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewDeleteUserAccountAdmin(request), s.config.TopicUserDeleteAdmin)
}

func (s *DataManagementServiceImpl) UpdateUserDataAdmin(ctx context.Context, request *sorm.UpdateUserDataAdminRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewUpdateUserDataAdmin(request), s.config.TopicUserUpdateAdmin)
}

func (s *DataManagementServiceImpl) DeleteAccount(ctx context.Context, request *sorm.DeleteAccountRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewDeleteAccount(request), s.config.TopicDeleteAccount)
}

func (s *DataManagementServiceImpl) UserAccountRecovery(ctx context.Context, request *sorm.UserAccountRecoveryRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewUserAccountRecovery(request), s.config.TopicUserAccountRecovery)
}

func (s *DataManagementServiceImpl) DirectoryData(ctx context.Context, request *sorm.DirectoryDataRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewDirectoryData(request), s.config.TopicDirectory)
}

func (s *DataManagementServiceImpl) MigrateUser(ctx context.Context, request *sorm.MigrateUserRequest) (*sorm.DataManagementResponse, error) {
	return s.Handle(&ctx, factory.NewMigrateUser(request), s.config.TopicMigrateUser)
}

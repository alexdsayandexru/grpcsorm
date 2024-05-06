package grpc

import (
	"context"
	"crypto/tls"
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"gitlab.gid.team/sso/auth/sorm/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"
)

func Test_Response(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	s.Response(codes.OK, "", "")
}

func Test_RegisterGRPC(t *testing.T) {
	defer func() { recover() }()
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	s.RegisterGRPC(&grpc.Server{})
}

func Test_RegisterUser(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.RegisterUser(context.Background(), &sorm.RegisterUserRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_LoginUser(t *testing.T) {
	config := configs.Config{}
	config.EnabledValidation = true
	s := NewDataManagementService(&config, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.LoginUser(context.Background(), &sorm.LoginUserRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_LogoutUser(t *testing.T) {
	config := configs.Config{}
	config.EnabledSecurity = true
	s := NewDataManagementService(&config, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.LogoutUser(context.Background(), &sorm.LogoutUserRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_DeleteUserAccount(t *testing.T) {
	config := configs.Config{}
	config.EnabledDelivery = true
	s := NewDataManagementService(&config, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.DeleteUserAccount(context.Background(), &sorm.DeleteUserAccountRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_UpdateUserData(t *testing.T) {
	config := configs.Config{}
	config.EnabledDelivery = true
	config.EnabledTls = true
	s := NewDataManagementService(&config, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.UpdateUserData(context.Background(), &sorm.UpdateUserDataRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_DeleteUserAccountAdmin(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.DeleteUserAccountAdmin(context.Background(), &sorm.DeleteUserAccountAdminRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_UpdateUserDataAdmin(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.UpdateUserDataAdmin(context.Background(), &sorm.UpdateUserDataAdminRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_DeleteAccount(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.DeleteAccount(context.Background(), &sorm.DeleteAccountRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_UserAccountRecovery(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.UserAccountRecovery(context.Background(), &sorm.UserAccountRecoveryRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_DirectoryData(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.DirectoryData(context.Background(), &sorm.DirectoryDataRequest{}); err != nil {
		t.Error(err)
	}
}

func Test_MigrateUser(t *testing.T) {
	s := NewDataManagementService(&configs.Config{}, &tls.Config{}, func(code codes.Code) {}, true)
	if _, err := s.MigrateUser(context.Background(), &sorm.MigrateUserRequest{}); err != nil {
		t.Error(err)
	}
}

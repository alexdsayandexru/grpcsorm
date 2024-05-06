package factory

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"testing"
)

func Test_Common(t *testing.T) {
	PhonesToStrings([]*sorm.Msisdn{{}})
	EmailsToStrings([]*sorm.Email{{}})
	AddsToAdditionals([]*sorm.Add{{}})
	ImIdsToImIds([]*sorm.ImId{{}})
	ServicesToServices([]*sorm.Service{{}})
}

func Test_NewEntity(t *testing.T) {
	NewDeleteAccount(&sorm.DeleteAccountRequest{})
	NewDeleteUserAccount(&sorm.DeleteUserAccountRequest{})
	NewDeleteUserAccountAdmin(&sorm.DeleteUserAccountAdminRequest{})
	NewDirectoryData(&sorm.DirectoryDataRequest{})
	NewLoginUser(&sorm.LoginUserRequest{})
	NewLogoutUser(&sorm.LogoutUserRequest{})
	NewMigrateUser(&sorm.MigrateUserRequest{})
	NewRegisterUser(&sorm.RegisterUserRequest{})
	NewUpdateUserData(&sorm.UpdateUserDataRequest{})
	NewUpdateUserDataAdmin(&sorm.UpdateUserDataAdminRequest{})
	NewUserAccountRecovery(&sorm.UserAccountRecoveryRequest{})
}

package main

import (
	"sync"
	"testing"
)

func Test_AllEntitiesAsinc(t *testing.T) {
	if PerformanceTest {
		wg := sync.WaitGroup{}
		grpc := GRPCClient{}
		if err := grpc.Open(); err != nil {
			t.Error()
		} else {
			client, _ := grpc.GetClient()
			ctx := grpc.GetCTX()
			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if resp, _ := client.RegisterUser(ctx, RegisterUserRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.DeleteAccount(ctx, DeleteAccountRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.DeleteUserAccountAdmin(ctx, DeleteUserAccountAdminRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.DeleteUserAccount(ctx, DeleteUserAccountRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.DirectoryData(ctx, DirectoryDataRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.LoginUser(ctx, LoginUserRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.LogoutUser(ctx, LogoutUserRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.UpdateUserDataAdmin(ctx, UpdateUserDataAdminRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.UpdateUserData(ctx, UpdateUserDataRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.UserAccountRecovery(ctx, UserAccountRecoveryRequest()); resp.Code != 0 {
						t.Error(resp)
					}
					if resp, _ := client.MigrateUser(ctx, MigrateUserRequest()); resp.Code != 0 {
						t.Error(resp)
					}
				}()
			}
		}
		wg.Wait()
		grpc.Close()
	}
}

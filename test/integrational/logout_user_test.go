package main

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"testing"
)

func LogoutUserRequest() *sorm.LogoutUserRequest {
	return &sorm.LogoutUserRequest{
		CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
		TelcoId:       99,
		UserType:      100,
		UserId:        "qweRTY-_1234567890",
		ServiceId:     100000000,
		Ip:            "127.0.0.1",
		Port:          99999,
		UserAgent:     "qwerty QWERTY-",
		Datetime:      "05-03-2024:23.23.23.003",
	}
}

func Test_LogoutUserRequest(t *testing.T) {
	grpc := GRPCClient{}
	if err := grpc.Open(); err != nil {
		t.Error(err)
	} else {
		defer grpc.Close()
		if client, err := grpc.GetClient(); err != nil {
			t.Error(err)
		} else if resp, _ := client.LogoutUser(grpc.GetCTX(), LogoutUserRequest()); resp.Code != 0 {
			t.Error(resp)
		}
	}
}

package main

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"testing"
)

func LoginUserRequest() *sorm.LoginUserRequest {
	return &sorm.LoginUserRequest{
		CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
		TelcoId:       99,
		UserType:      100,
		UserId:        "qweRTY-_1234567890",
		ServiceId:     100000000,
		MsisdnLogin:   "+79261234567",
		EmailLogin:    "qwerty@mail.ru",
		Ip:            "127.0.0.1",
		Port:          99999,
		UserAgent:     "qwerty QWERTY-",
		Factor:        "123qwertyQWERTY-",
		Datetime:      "05-03-2024:23.23.23.003",
	}
}

func Test_LoginUserRequest(t *testing.T) {
	grpc := GRPCClient{}
	if err := grpc.Open(); err != nil {
		t.Error(err)
	} else {
		defer grpc.Close()
		if client, err := grpc.GetClient(); err != nil {
			t.Error(err)
		} else if resp, _ := client.LoginUser(grpc.GetCTX(), LoginUserRequest()); resp.Code != 0 {
			t.Error(resp)
		}
	}
}

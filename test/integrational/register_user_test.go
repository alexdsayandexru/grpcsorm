package main

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"testing"
	"time"
)

func timeNow() string {
	return time.Now().Format("2006-01-02:15.04.06.123")[0:23]
}

func RegisterUserRequest() *sorm.RegisterUserRequest {
	return &sorm.RegisterUserRequest{
		CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
		TelcoId:       99,
		UserType:      100,
		UserId:        "qweRTY-_1234567890",
		Msisdns:       []*sorm.Msisdn{{Msisdn: "+79261234567"}},
		Emails:        []*sorm.Email{{Email: "qwerty@mail.ru"}},
		DatetimeReg:   timeNow(),
		ServiceUser:   1,
		ContractDate:  timeNow(),
		ServiceId:     100000000,
		MsisdnLogin:   "+79261234567",
		EmailLogin:    "qwerty@mail.ru",
		Ip:            "127.0.0.1",
		Port:          99999,
		UserAgent:     "qwerty QWERTY-",
		Message:       "qwertyuiop1234567890",
		Datetime:      timeNow(),
		Additional:    []*sorm.Add{{Title: "", Content: ""}},
	}
}

func Test_RegisterUserRequest(t *testing.T) {
	grpc := GRPCClient{}
	if err := grpc.Open(); err != nil {
		t.Error(err)
	} else {
		defer grpc.Close()
		if client, err := grpc.GetClient(); err != nil {
			t.Error(err)
		} else if resp, _ := client.RegisterUser(grpc.GetCTX(), RegisterUserRequest()); resp.Code != 0 {
			t.Error(resp)
		}
	}
}

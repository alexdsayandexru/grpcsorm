package main

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"testing"
)

func DeleteAccountRequest() *sorm.DeleteAccountRequest {
	return &sorm.DeleteAccountRequest{
		CorrelationId:   "01548f2c-dabf-11ee-a506-0242ac120002",
		TelcoId:         99,
		UserType:        100,
		UserId:          "qweRTY-_1234567890",
		Msisdns:         []*sorm.Msisdn{{Msisdn: "+79261234567"}},
		DatetimeReg:     timeNow(),
		Emails:          []*sorm.Email{{Email: "qwerty@mail.ru"}},
		DatetimeUpdated: timeNow(),
		ServiceUser:     1,
		Additional:      []*sorm.Add{{Title: "", Content: ""}},
		ContractDate:    timeNow(),
		Nick:            "qwertyйцукенг",
		BirthDate:       "05-03-2024",
		Family:          "qwertyйцукенг",
		Initial:         "qwertyйцукенг",
		Address:         "qwertx yйцуvd zfvdzvaefcd  gкенг",
		ImIds:           []*sorm.ImId{{ServiceId: "ssss", ServiceName: "sssxsxa"}},
		DatetimeUnreg:   timeNow(),
	}
}

func Test_DeleteAccountRequest(t *testing.T) {
	grpc := GRPCClient{}
	if err := grpc.Open(); err != nil {
		t.Error(err)
	} else {
		defer grpc.Close()
		if client, err := grpc.GetClient(); err != nil {
			t.Error(err)
		} else if resp, _ := client.DeleteAccount(grpc.GetCTX(), DeleteAccountRequest()); resp.Code != 0 {
			t.Error(resp)
		}
	}
}

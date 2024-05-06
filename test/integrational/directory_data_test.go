package main

import (
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"testing"
)

func DirectoryDataRequest() *sorm.DirectoryDataRequest {
	return &sorm.DirectoryDataRequest{
		CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
		Datetime:      "05-03-2024:23.23.23.003",
		Services: []*sorm.Service{
			{ServiceId: 1, Decription: "=hello=вцвцац", BeginTime: "05-03-2024:23.23.23.003", EndTime: "05-03-2024:23.23.23.003"},
			{ServiceId: 1, Decription: "=hello=вцвцац", BeginTime: "05-03-2024:23.23.23.003"},
			{ServiceId: 1, Decription: "=hello=вцвцац", BeginTime: "05-03-2024:23.23.23.003", EndTime: "05-03-2024:23.23.23.003"},
		},
	}
}

func Test_DirectoryDataRequest(t *testing.T) {
	grpc := GRPCClient{}
	if err := grpc.Open(); err != nil {
		t.Error(err)
	} else {
		defer grpc.Close()
		if client, err := grpc.GetClient(); err != nil {
			t.Error(err)
		} else if resp, _ := client.DirectoryData(grpc.GetCTX(), DirectoryDataRequest()); resp.Code != 0 {
			t.Error(resp)
		}
	}
}

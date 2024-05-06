package main

import (
	"context"
	sorm "gitlab.gid.team/sso/auth/sorm/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	PerformanceTest = true
	XApiKey         = "X-API-KEY"
	XApiKeyValue    = "HELLOWORLD"
	//GrpcTargetClientName = ":8181"
	GrpcTargetClientName = "sorm-dev01-grpc.k8s-dev.gid.team:443"
)

type GRPCClient struct {
	conn *grpc.ClientConn
	err  error
}

func (c *GRPCClient) Open() error {
	//c.conn, c.err = grpc.Dial(GrpcTargetClientName, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c.conn, c.err = grpc.Dial(GrpcTargetClientName, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	return c.err
}

func (c *GRPCClient) Close() {
	_ = c.conn.Close()
}

func (c *GRPCClient) GetClient() (sorm.UserDataManagementClient, error) {
	return sorm.NewUserDataManagementClient(c.conn), nil
}

func (c *GRPCClient) GetCTX() context.Context {
	md := metadata.Pairs(XApiKey, XApiKeyValue)
	return metadata.NewOutgoingContext(context.Background(), md)
}

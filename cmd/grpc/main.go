package main

import (
	"context"
	"gitlab.gid.team/research-development/libraries/go-libs/pkg/application"
	"gitlab.gid.team/research-development/libraries/go-libs/pkg/log"
	"gitlab.gid.team/sso/auth/sorm/configs"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"os"
	"sync"
)

func main() {
	ctx := context.Background()
	app, err := application.NewAppWithContext(ctx,
		application.WithWatchers(
			application.Metrics,
			application.Sentry,
			application.Health,
		),
		application.WithServices(application.GRPCServer))

	if err != nil {
		log.Error(ctx, "init app", err, err.Error())
		os.Exit(1)
	} else {
		log.Info(ctx, "init app", "2024-04-23 15:35")
	}

	config := configs.GetConfig()
	tls := configs.GetTLS(config.CaCert, config.EnabledTls)

	dataManagementService := grpc.NewDataManagementService(
		config,
		tls,
		func(code codes.Code) {
			healthCheckResponseEventHandler(app, code)
		}, false)

	app.RegisterGRPCHandlers(dataManagementService)

	app.Run()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	mu      sync.Mutex
	current = healthpb.HealthCheckResponse_UNKNOWN
)

func healthCheckResponseEventHandler(app *application.App, code codes.Code) {
	mu.Lock()
	defer mu.Unlock()

	if code == codes.OK && current != healthpb.HealthCheckResponse_SERVING {
		current = healthpb.HealthCheckResponse_SERVING
		app.GRPCServer.HealthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	} else if code == codes.Aborted && current != healthpb.HealthCheckResponse_NOT_SERVING {
		current = healthpb.HealthCheckResponse_NOT_SERVING
		app.GRPCServer.HealthServer.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
	}
}

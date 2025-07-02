package grpcclients

import (
	"log"
	"time"

	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SsoSvcUser(cfg config.Config) sso.UserServiceClient {
	conn, err := grpc.NewClient(
		cfg.Services.SSO.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("cannot dial SSO service %q: %v", cfg.Services.SSO.Url, err)
	}

	client := sso.NewUserServiceClient(conn)

	return client
}

func SsoSvcAdmin(cfg config.Config) sso.AdminServiceClient {
	conn, err := grpc.NewClient(
		cfg.Services.SSO.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("cannot dial SSO service %q: %v", cfg.Services.SSO.Url, err)
	}

	client := sso.NewAdminServiceClient(conn)

	return client
}

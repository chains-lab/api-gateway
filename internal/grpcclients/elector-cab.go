package grpcclients

import (
	"log"
	"time"

	"github.com/chains-lab/api-gateway/internal/config"
	elcab "github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ElectorCabUserSvc(cfg config.Config) elcab.UserServiceClient {
	conn, err := grpc.NewClient(
		cfg.Services.ElectorCab.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("cannot dial Elector Cab service %q: %v", cfg.Services.ElectorCab.Url, err)
	}

	client := elcab.NewUserServiceClient(conn)

	return client
}

func ElectorCabAdminSvc(cfg config.Config) elcab.AdminServiceClient {
	conn, err := grpc.NewClient(
		cfg.Services.ElectorCab.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("cannot dial Elector Cab service %q: %v", cfg.Services.ElectorCab.Url, err)
	}

	client := elcab.NewAdminServiceClient(conn)

	return client
}

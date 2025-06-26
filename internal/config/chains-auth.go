package config

import (
	"log"
	"time"

	"github.com/chains-lab/proto-storage/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Config) ChainsAuth() auth.ServiceClient {
	conn, err := grpc.NewClient(
		c.Services.SSO.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("cannot dial SSO service %q: %v", c.Services.SSO.Address, err)
	}

	client := auth.NewServiceClient(conn)

	return client
}

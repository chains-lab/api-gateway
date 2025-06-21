package config

import (
	"log"
	"time"

	pb "github.com/chains-lab/proto-storage/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Config) ChainsAuth() pb.SsoServiceClient {
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

	client := pb.NewSsoServiceClient(conn)

	return client
}

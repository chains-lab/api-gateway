package config

import (
	"log"
	"time"

	"github.com/chains-lab/proto-storage/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Config) SsoSvc() sso.ServiceClient {
	conn, err := grpc.NewClient(
		c.Services.SSO.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("cannot dial SSO service %q: %v", c.Services.SSO.Url, err)
	}

	client := sso.NewServiceClient(conn)

	return client
}

package signer

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

// TODO In future will be changed to use a more secure way to store secret key
var secretKey string

const serviceName = "api-gateway"

func ServiceToken(r *http.Request, requestID uuid.UUID, audience []string) (context.Context, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return r.Context(), fmt.Errorf("failed to load config with sk: %w", err)
	}

	secretKey = cfg.JWT.Service.SecretKey

	token, err := tokens.GenerateServiceJWT(tokens.GenerateServiceJwtRequest{
		Issuer:   serviceName,
		Subject:  serviceName,
		Audience: audience,
		Ttl:      15 * time.Second,
	}, secretKey)
	if err != nil {
		return r.Context(), fmt.Errorf("failed to generate service token: %w", err)
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return r.Context(), fmt.Errorf("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return r.Context(), fmt.Errorf("missing or invalid Authorization header")
	}

	return metadata.AppendToOutgoingContext(
		r.Context(),
		"authorization", fmt.Sprintf("Bearer %s", token),
		"x-user-token", fmt.Sprintf("Bearer %s", parts[1]),
		"x-request-id", requestID.String(),
	), nil
}

package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func SsoUserClient(r *http.Request) sso.UserServiceClient {
	return r.Context().Value(middleware.ChainsSsoUserCtxKey).(sso.UserServiceClient)
}

func SsoAdminClient(r *http.Request) sso.AdminServiceClient {
	return r.Context().Value(middleware.ChainsSsoAdminCtxKey).(sso.AdminServiceClient)
}

func Log(r *http.Request, requestID uuid.UUID) *logrus.Entry {
	log, ok := r.Context().Value(middleware.LogCtxKey).(*logrus.Logger)
	if !ok {
		panic("log entry not found in context")
	}
	return log.WithField("request_id", requestID.String())
}

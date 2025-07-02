package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	elcab "github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
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

func ElectorCabUserClient(r *http.Request) elcab.UserServiceClient {
	return r.Context().Value(middleware.ChainsElectorCabUserCtxKey).(elcab.UserServiceClient)
}

func ElectorCabAdminClient(r *http.Request) elcab.AdminServiceClient {
	return r.Context().Value(middleware.ChainsElectorCabAdminCtxKey).(elcab.AdminServiceClient)
}

func Log(r *http.Request, requestID uuid.UUID) *logrus.Entry {
	entry, ok := r.Context().Value(middleware.LogCtxKey).(*logrus.Entry)
	if !ok {
		entry = logrus.NewEntry(logrus.New())
	}
	return entry.WithField("request_id", requestID)
}

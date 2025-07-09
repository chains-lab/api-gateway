package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	elcab "github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func ElectorCabUserClient(r *http.Request) elcab.UserServiceClient {
	return r.Context().Value(middleware.ChainsElectorCabUserCtxKey).(elcab.UserServiceClient)
}

func ElectorCabAdminClient(r *http.Request) elcab.AdminServiceClient {
	return r.Context().Value(middleware.ChainsElectorCabAdminCtxKey).(elcab.AdminServiceClient)
}

func Log(r *http.Request, requestID uuid.UUID) *logrus.Entry {
	log, ok := r.Context().Value(middleware.LogCtxKey).(*logrus.Logger)
	if !ok {
		panic("log entry not found in context")
	}
	return log.WithField("request_id", requestID.String())
}

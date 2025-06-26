package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func AuthClient(r *http.Request) sso.ServiceClient {
	return r.Context().Value(middleware.ChainsAutHCtxKey).(sso.ServiceClient)
}

func Log(r *http.Request, requestID uuid.UUID) *logrus.Entry {
	entry, ok := r.Context().Value(middleware.LogCtxKey).(*logrus.Entry)
	if !ok {
		entry = logrus.NewEntry(logrus.New())
	}
	return entry.WithField("request_id", requestID)
}

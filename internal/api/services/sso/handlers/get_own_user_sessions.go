package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetOwnUserSessions(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for own sessions")
		renderer.InternalError(w, &requestID)

		return
	}

	sessions, err := SsoUserClient(r).GetOwnUserSessions(signature, &emptypb.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving own sessions")
		renderer.RenderGRPCError(w, err)

		return
	}

	renderer.Render(w, responses.SessionCollection(sessions))
}

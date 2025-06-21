package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/auth/responses"
	"github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/google/uuid"
)

func OwnGetSession(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-auth"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for own session")
		renderer.InternalError(w, requestID)

		return
	}

	sessions, err := AuthClient(r).GetOwnUserSessions(signature, &sso.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving own session")
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	renderer.Render(w, responses.SessionCollection(sessions))
}

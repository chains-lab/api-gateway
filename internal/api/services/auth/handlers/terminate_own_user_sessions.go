package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/proto-storage/gen/go/auth"
	"github.com/google/uuid"
)

func OwnTerminateSessions(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-auth"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for own session termination")
		renderer.InternalError(w, requestID)

		return
	}

	_, err = AuthClient(r).TerminateOwnUserSessions(signature, &auth.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error terminating own sessions")
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	renderer.Render(w, http.StatusNoContent)
}

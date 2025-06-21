package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/google/uuid"
)

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-auth"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", initiator.UserID)
		renderer.InternalError(w, requestID)

		return
	}

	_, err = AuthClient(r).DeleteOwnUserSession(signature, &sso.DeleteOwnUserSessionRequest{
		SessionId: initiator.SessionID.String(),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error deleting session for user %s", initiator.UserID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	Log(r, requestID).Infof("session %s deleted successfully for user %s", initiator.SessionID, initiator.UserID)
	renderer.Render(w, http.StatusNoContent)
}

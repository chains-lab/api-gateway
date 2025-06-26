package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/google/uuid"
)

func AdminDeleteUserSession(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing user ID %s", r.URL.Query().Get("user_id"))
		renderer.BadRequest(w, requestID, "The provided user ID is not a valid UUID.")

		return
	}

	sessionID, err := uuid.Parse(r.URL.Query().Get("session_id"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing session ID %s", r.URL.Query().Get("session_id"))
		renderer.BadRequest(w, requestID, "The provided session ID is not a valid UUID.")

		return
	}

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for session %s", sessionID)
		renderer.InternalError(w, requestID)

		return
	}

	_, err = AuthClient(r).AdminDeleteUserSession(signature, &sso.AdminDeleteUserSessionRequest{
		SessionId: sessionID.String(),
		UserId:    userID.String(),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error deleting session %s", sessionID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	renderer.Render(w, http.StatusNoContent)
}

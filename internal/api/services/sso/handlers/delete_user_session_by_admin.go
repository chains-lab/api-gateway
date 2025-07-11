package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func DeleteUserSessionByAdmin(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing user ID %s", chi.URLParam(r, "user_id"))
		renderer.BadRequest(w, requestID, "The provided user ID is not a valid UUID.")

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", userID)
		renderer.InternalError(w, &requestID)

		return
	}

	_, err = SsoAdminClient(r).DeleteUserSessionsByAdmin(signature, &sso.DeleteUserSessionsByAdminRequest{
		UserId: userID.String(),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error delete session for user %s by %s", userID, initiator.UserID)
		renderer.RenderGRPCError(w, err)

		return
	}

	Log(r, requestID).Infof("delete session %s for user %s", userID, initiator.UserID)
	renderer.Render(w, http.StatusAccepted)
}

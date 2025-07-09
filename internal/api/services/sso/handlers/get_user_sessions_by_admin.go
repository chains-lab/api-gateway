package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUserSessionsByAdmin(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing user ID %s", chi.URLParam(r, "user_id"))
		renderer.BadRequest(w, requestID, "The provided user ID is not a valid UUID.")

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", userID)
		renderer.InternalError(w, requestID)

		return
	}

	sessions, err := SsoAdminClient(r).GetUserSessionsByAdmin(signature, &sso.GetUserSessionsByAdminRequest{
		UserId: userID.String(),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving sessions for user %s", userID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	renderer.Render(w, responses.SessionCollection(sessions))
}

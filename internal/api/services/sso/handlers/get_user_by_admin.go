package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/google/uuid"
)

func GetUserByAdmin(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing user ID %s", r.URL.Query().Get("user_id"))
		renderer.BadRequest(w, requestID, "The provided user ID is not a valid UUID.")

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", userID)
		renderer.InternalError(w, &requestID)

		return
	}

	user, err := SsoAdminClient(r).GetUserByAdmin(signature, &sso.GetUserByAdminRequest{
		UserId: userID.String(),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving user %s", userID)
		renderer.RenderGRPCError(w, err)

		return
	}

	renderer.Render(w, responses.User(user))
}

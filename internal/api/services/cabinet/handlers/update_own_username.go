package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/responses"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UpdateOwnUsername(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	signature, err := signer.ServiceToken(r, requestID, []string{"elector-cab-svc"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", initiator.UserID)
		renderer.InternalError(w, requestID)

		return
	}

	username := chi.URLParam(r, "username")

	res, err := ElectorCabUserClient(r).UpdateOwnUsername(signature, &electorcab.UpdateOwnUsernameRequest{
		Username: username,
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error updating own username for user %s", initiator.UserID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	renderer.Render(w, responses.Profile(res))
}

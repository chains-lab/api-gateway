package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/requests"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/google/uuid"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	req, err := requests.NewRefresh(r)
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing refresh request")
		renderer.BadRequestValidate(w, requestID, err)

		return
	}

	signature, err := signer.SignWithoutUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for own session termination")
		renderer.InternalError(w, &requestID)

		return
	}
	resp, err := SsoUserClient(r).RefreshToken(signature, &sso.RefreshTokenRequest{
		Agent:        "agent",
		RefreshToken: req.Data.Attributes.RefreshToken,
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error refreshing token")
		renderer.RenderGRPCError(w, err)

		return
	}

	Log(r, requestID).Infof("token refreshed successfully for user %s", initiator.UserID)
	renderer.Render(w, responses.TokensPair(resp))
}

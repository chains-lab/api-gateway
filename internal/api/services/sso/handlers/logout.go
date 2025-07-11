package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", initiator.UserID)
		renderer.InternalError(w, &requestID)

		return
	}

	_, err = SsoUserClient(r).Logout(signature, &emptypb.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error logging out user %s", initiator.UserID)
		renderer.RenderGRPCError(w, err)

		return
	}

	Log(r, requestID).Infof("user %s logged out successfully", initiator.UserID)
	renderer.Render(w, http.StatusNoContent)
}

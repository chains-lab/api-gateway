package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/responses"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetOwnCabinet(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"elector-cab-svc"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", initiator.UserID)
		renderer.InternalError(w, &requestID)

		return
	}

	cabinet, err := ElectorCabUserClient(r).GetOwnCabinet(signature, &emptypb.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving own cabinet for user %s", initiator.UserID)
		renderer.RenderGRPCError(w, err)

		return
	}

	renderer.Render(w, responses.Cabinet(cabinet))
}

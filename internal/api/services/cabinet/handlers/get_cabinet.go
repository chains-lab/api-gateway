package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/responses"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"github.com/google/uuid"
)

func GetCabinet(w http.ResponseWriter, r *http.Request) {
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

	username := r.URL.Query().Get("username")
	userIDStr := r.URL.Query().Get("user_id")

	grpcReq := &electorcab.GetCabinetRequest{}

	if username != "" {
		grpcReq.Identifier = &electorcab.GetCabinetRequest_Username{
			Username: username,
		}
	} else if userIDStr != "" {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			Log(r, requestID).WithError(err).Errorf("error parsing user ID %s", userIDStr)
			renderer.BadRequest(w, requestID, "The provided user ID is not a valid UUID.")

			return
		}
		grpcReq.Identifier = &electorcab.GetCabinetRequest_UserId{
			UserId: userID.String(),
		}
	} else {
		Log(r, requestID).Error("neither username nor user_id provided")
		renderer.BadRequest(w, requestID, "Either username or user_id must be provided.")

		return
	}

	cabinet, err := ElectorCabUserClient(r).GetCabinet(signature, grpcReq)
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving cabinet for user %s", initiator.UserID)
		renderer.RenderGRPCError(w, err)

		return
	}

	renderer.Render(w, responses.Cabinet(cabinet))
}

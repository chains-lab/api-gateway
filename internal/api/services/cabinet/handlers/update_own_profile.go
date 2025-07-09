package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/requests"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/responses"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"github.com/google/uuid"
)

func UpdateOwnProfile(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	req, err := requests.NewUpdateOwnProfile(r)
	if err != nil {
		Log(r, requestID).WithError(err).Error("error parsing request body for updating own profile")
		renderer.BadRequestValidate(w, requestID, err)

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"elector-cab-svc"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", initiator.UserID)
		renderer.InternalError(w, requestID)

		return
	}

	res, err := ElectorCabUserClient(r).UpdateOwnProfile(signature, &electorcab.UpdateOwnProfileRequest{
		Pseudonym:   req.Data.Attributes.Pseudonym,
		Description: req.Data.Attributes.Description,
		Avatar:      req.Data.Attributes.Avatar,
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error updating own profile for user %s", initiator.UserID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	renderer.Render(w, responses.Profile(res))
}

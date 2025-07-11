package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/requests"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"github.com/google/uuid"
)

func CreateOwnProfile(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	req, err := requests.NewCreateOwnProfile(r)
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error decoding request body for own profile creation")
		renderer.BadRequestValidate(w, requestID, err)

		return
	}

	signature, err := signer.SignWithUser(r, requestID, []string{"elector-cab-svc"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user")
		renderer.InternalError(w, &requestID)

		return
	}

	profile, err := ElectorCabUserClient(r).CreateOwnProfile(signature, &electorcab.CreateProfileRequest{
		Username:    req.Data.Attributes.Username,
		Pseudonym:   req.Data.Attributes.Pseudonym,
		Description: req.Data.Attributes.Description,
		Avatar:      req.Data.Attributes.Avatar,
	})
	if err != nil {
		Log(r, requestID).WithError(err).Error("error creating own profile")
		renderer.RenderGRPCError(w, err)

		return
	}

	Log(r, requestID).Infof("user %s created own profile with username: %s", initiator.UserID, profile.Username)
}

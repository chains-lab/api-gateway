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

func CreateUserByAdmin(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	req, err := requests.NewCreateUserByAdmin(r)
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error decoding request body for user creation by admin")
		renderer.BadRequestValidate(w, requestID, err)

		return
	}

	email := req.Data.Attributes.Email
	role := req.Data.Attributes.Role

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user")
		renderer.InternalError(w, requestID)

		return
	}

	user, err := SsoAdminClient(r).CreateUserByAdmin(signature, &sso.CreateUserByAdminRequest{
		Email: email,
		Role:  role,
	})
	if err != nil {
		Log(r, requestID).WithError(err).Error("error creating user by admin")
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	Log(r, requestID).Infof("user %s create user with email: %s, role: %s", initiator.UserID, user.Email, user.Role)
	renderer.Render(w, responses.User(user))
}

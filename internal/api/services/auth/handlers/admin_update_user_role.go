package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/auth/responses"
	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func AdminUpdateRole(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	initiator, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		Log(r, requestID).WithError(err).Error("error retrieving account data from token")
		renderer.Unauthorized(w, requestID, "Error retrieving account data from token")

		return
	}

	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing user ID %s", chi.URLParam(r, "user_id"))
		renderer.BadRequest(w, requestID, "The provided user ID is not a valid UUID.")

		return
	}

	role, err := roles.ParseRole(chi.URLParam(r, "role"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing role %s", chi.URLParam(r, "role"))
		renderer.BadRequest(w, requestID, "The provided role is not valid.")

		return
	}

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-auth"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", userID)
		renderer.InternalError(w, requestID)

		return
	}

	user, err := AuthClient(r).AdminUpdateUserRole(signature, &auth.AdminUpdateUserRoleRequest{
		UserId: userID.String(),
		Role:   string(role),
	})

	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error updating user role for user %s", userID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	Log(r, requestID).Infof("user %s role updated to %s by %s", userID, role, initiator.UserID)
	renderer.Render(w, responses.User(user))
}

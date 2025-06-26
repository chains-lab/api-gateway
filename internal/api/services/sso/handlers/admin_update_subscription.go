package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func AdminUpdateSubscription(w http.ResponseWriter, r *http.Request) {
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

	subscription, err := uuid.Parse(chi.URLParam(r, "subscription"))
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error parsing subscription ID %s", chi.URLParam(r, "subscription"))
		renderer.BadRequest(w, requestID, "The provided subscription ID is not a valid UUID.")

		return
	}

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for user %s", userID)
		renderer.InternalError(w, requestID)

		return
	}

	user, err := AuthClient(r).AdminUpdateUserSubscription(signature, &sso.AdminUpdateUserSubscriptionRequest{
		UserId:       userID.String(),
		Subscription: subscription.String(),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error updating user subscription for user %s", userID)
		renderer.RenderGRPCError(w, requestID, err)

		return
	}

	Log(r, requestID).WithField("user_id", userID).Infof("user %s subscription updated to %s by %s", userID, subscription, initiator.UserID)
	renderer.Render(w, responses.User(user))
}

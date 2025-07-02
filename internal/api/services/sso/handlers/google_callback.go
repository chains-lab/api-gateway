package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/google/uuid"
)

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	signature, err := signer.ServiceToken(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for new user")
		renderer.InternalError(w, requestID)

		return
	}

	resp, err := SsoUserClient(r).GoogleCallback(signature, &sso.GoogleCallbackRequest{
		Code: r.URL.Query().Get("code"),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving Google callback response")
		renderer.InternalError(w, requestID)

		return
	}

	renderer.Render(w, responses.TokensPair(resp))
}

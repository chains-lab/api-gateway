package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/proto-storage/gen/go/sso"
	"github.com/google/uuid"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	//TODO change
	signature, err := signer.ServiceToken(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for new user")
		renderer.InternalError(w, requestID)

		return
	}

	rep, err := AuthClient(r).GoogleLogin(signature, &sso.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving Google login URL")
		renderer.InternalError(w, requestID)

		return
	}

	http.Redirect(w, r, rep.Url, http.StatusTemporaryRedirect)
}

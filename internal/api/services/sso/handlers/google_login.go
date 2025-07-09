package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	//TODO change
	signature, err := signer.SignWithoutUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for new user")
		renderer.InternalError(w, requestID)

		return
	}

	rep, err := SsoUserClient(r).GoogleLogin(signature, &emptypb.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving Google login URL")
		renderer.InternalError(w, requestID)

		return
	}

	http.Redirect(w, r, rep.Url, http.StatusTemporaryRedirect)
}

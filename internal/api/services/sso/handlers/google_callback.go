package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/responses"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/protoadapt"
)

type Error struct {
	code    codes.Code
	reason  string
	message string
	details []protoadapt.MessageV1
	cause   error
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	signature, err := signer.SignWithoutUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for new user")
		renderer.InternalError(w, &requestID)

		return
	}

	resp, err := SsoUserClient(r).GoogleCallback(signature, &sso.GoogleCallbackRequest{
		Code: r.URL.Query().Get("code"),
	})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error retrieving Google callback response")
		renderer.InternalError(w, &requestID)

		return
	}

	renderer.Render(w, responses.TokensPair(resp))
}

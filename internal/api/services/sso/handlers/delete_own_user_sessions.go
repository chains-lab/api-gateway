package handlers

import (
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/renderer"
	"github.com/chains-lab/api-gateway/internal/api/common/signer"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func DeleteOwnUserSessions(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	signature, err := signer.SignWithUser(r, requestID, []string{"chains-sso"})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error signing service token for own session termination")
		renderer.InternalError(w, &requestID)

		return
	}

	_, err = SsoUserClient(r).DeleteOwnUserSessions(signature, &emptypb.Empty{})
	if err != nil {
		Log(r, requestID).WithError(err).Errorf("error delete own sessions")
		renderer.RenderGRPCError(w, err)

		return
	}

	renderer.Render(w, http.StatusNoContent)
}

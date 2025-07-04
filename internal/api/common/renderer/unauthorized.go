package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func Unauthorized(w http.ResponseWriter, requestID uuid.UUID, message string) {
	httpkit.RenderErr(w, &jsonapi.ErrorObject{
		Title:  "Unauthorized",
		Status: http.StatusText(http.StatusUnauthorized),
		Detail: message,
		Meta: &map[string]any{
			"request_id": requestID.String(),
		},
	})
}

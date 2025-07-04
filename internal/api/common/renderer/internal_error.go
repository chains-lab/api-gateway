package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func InternalError(w http.ResponseWriter, requestID uuid.UUID) {
	httpkit.RenderErr(w, &jsonapi.ErrorObject{
		Title:  "Internal Error",
		Status: http.StatusText(http.StatusInternalServerError),
		Detail: "An unexpected error occurred. Please try again later.",
		Meta: &map[string]any{
			"request_id": requestID.String(),
		},
	})
}

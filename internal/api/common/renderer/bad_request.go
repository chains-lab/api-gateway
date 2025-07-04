package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func BadRequest(w http.ResponseWriter, requestID uuid.UUID, details string) {
	httpkit.RenderErr(w, &jsonapi.ErrorObject{
		Title:  "Bad Request",
		Status: http.StatusText(http.StatusBadRequest),
		Detail: details,
		Meta: &map[string]any{
			"request_id": requestID.String(),
		},
	})
}

package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func BadRequestValidate(w http.ResponseWriter, requestID uuid.UUID, err error) {
	httpkit.RenderErr(w, &jsonapi.ErrorObject{
		Title:  "Bad Request",
		Status: http.StatusText(http.StatusBadRequest),
		Detail: err.Error(),
		Meta: &map[string]any{
			"request_id": requestID.String(),
		},
	})
}

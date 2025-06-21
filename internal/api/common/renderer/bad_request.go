package renderer

import (
	"net/http"

	"github.com/google/uuid"
)

func BadRequest(w http.ResponseWriter, requestID uuid.UUID, details string) {
	writeJSONAPIError(w, http.StatusBadRequest, "Bad Request", details, requestID)
}

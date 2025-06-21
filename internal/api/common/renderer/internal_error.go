package renderer

import (
	"net/http"

	"github.com/google/uuid"
)

func InternalError(w http.ResponseWriter, requestID uuid.UUID) {
	writeJSONAPIError(w, http.StatusInternalServerError, "Internal Error", "Unexpected error", requestID)
}

package renderer

import (
	"net/http"

	"github.com/google/uuid"
)

func Unauthorized(w http.ResponseWriter, requestID uuid.UUID, message string) {
	writeJSONAPIError(w, http.StatusUnauthorized, "Unauthorized", message, requestID)
}

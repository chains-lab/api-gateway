package renderer

import (
	"net/http"

	"github.com/google/uuid"
)

func BadRequestValidate(w http.ResponseWriter, requestID uuid.UUID, err error) {
	writeJSONAPIError(w, http.StatusBadRequest, "Bad Request", err.Error(), requestID)
}

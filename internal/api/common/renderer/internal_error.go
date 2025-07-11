package renderer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func InternalError(w http.ResponseWriter, requestID *uuid.UUID) {
	meta := map[string]interface{}{
		"timestamp": time.Now().UTC().Format(time.RFC3339Nano),
	}

	if requestID != nil {
		meta["request_id"] = requestID.String()
	}

	httpkit.RenderErr(w, &jsonapi.ErrorObject{
		Code:   "INTERNAL_SERVER_ERROR",
		Title:  "Internal Error",
		Status: fmt.Sprintf("%d", http.StatusInternalServerError), // "500"
		Detail: "An unexpected error occurred. Please try again later.",
		Meta:   &meta,
	})
}

package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RenderGRPCError конвертирует gRPC-ошибку в JSON:API HTTP-ответ.
// - st.Code() мапится в http-статус.
// - st.Message() попадает в Detail.
// - Title берётся из набора читаемых заголовков.
func RenderGRPCError(w http.ResponseWriter, requestID uuid.UUID, err error) {
	// Разбираем gRPC-статус
	st, ok := status.FromError(err)
	if !ok {
		// не gRPC-ошибка — Internal Server Error
		writeJSONAPIError(w, http.StatusInternalServerError, "Internal Error", "Unexpected error", requestID)
		return
	}

	// Мапим код
	httpCode, title := mapCodeToHTTP(st.Code())

	// Рендерим JSON:API
	writeJSONAPIError(w, httpCode, title, st.Message(), requestID)
}

// mapCodeToHTTP возвращает HTTP-код и заголовок (Title) для данного gRPC-кода.
func mapCodeToHTTP(code codes.Code) (int, string) {
	switch code {
	case codes.OK:
		return http.StatusOK, "OK"

	case codes.InvalidArgument:
		return http.StatusBadRequest, "Bad Request"
	case codes.NotFound:
		return http.StatusNotFound, "Not Found"
	case codes.AlreadyExists:
		return http.StatusConflict, "Already Exists"
	case codes.PermissionDenied:
		return http.StatusForbidden, "Permission Denied"
	case codes.Unauthenticated:
		return http.StatusUnauthorized, "Unauthenticated"
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests, "Resource Exhausted"

	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, "Deadline Exceeded"
	case codes.Unavailable:
		return http.StatusServiceUnavailable, "Service Unavailable"

	default:
		// codes.Internal, codes.Unknown и проч.
		return http.StatusBadGateway, "Upstream Error"
	}
}

// writeJSONAPIError отсылает один JSON:API ErrorObject с заданными полями.
func writeJSONAPIError(w http.ResponseWriter, httpCode int, title, detail string, requestID uuid.UUID) {
	httpkit.RenderErr(w, &jsonapi.ErrorObject{
		Title:  title,
		Status: http.StatusText(httpCode),
		Detail: detail,
		Meta: &map[string]any{
			"request_id": requestID.String(),
		},
	})
}

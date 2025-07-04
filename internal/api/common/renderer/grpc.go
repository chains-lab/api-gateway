package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RenderGRPCError конвертирует gRPC-ошибку в JSON:API HTTP-ответ.
func RenderGRPCError(w http.ResponseWriter, requestID uuid.UUID, err error) {
	st, ok := status.FromError(err)
	if !ok {
		InternalError(w, requestID)
		return
	}
	httpCode := mapCodeToHTTP(st.Code())

	// Инициализируем Meta как указатель на карту
	meta := map[string]interface{}{
		"request_id": requestID.String(),
	}

	mainErr := &jsonapi.ErrorObject{
		ID:     requestID.String(),
		Status: http.StatusText(httpCode),
		Code:   st.Code().String(),
		Title:  st.Message(),
		Meta:   &meta,
	}

	var errs []*jsonapi.ErrorObject

	for _, d := range st.Details() {
		switch info := d.(type) {

		case *errdetails.ErrorInfo:
			// Добавляем в meta.error_info
			(*mainErr.Meta)["error_info"] = map[string]string{
				"reason": info.Reason,
				"domain": info.Domain,
			}

		case *errdetails.ResourceInfo:
			// Привязываем ресурс
			(*mainErr.Meta)["resource"] = map[string]string{
				"type":        info.ResourceType,
				"name":        info.ResourceName,
				"description": info.Description,
			}
			//mainErr.Source = &jsonapi.ErrorSource{
			//	Pointer: "/" + info.ResourceType,
			//}

		case *errdetails.BadRequest:
			// Для каждого field violation создаём отдельный объект
			for _, fv := range info.FieldViolations {
				errs = append(errs, &jsonapi.ErrorObject{
					Status: mainErr.Status,
					Code:   mainErr.Code,
					Title:  mainErr.Title,
					Detail: fv.Description,
					//Source: &jsonapi.ErrorSource{
					//	Pointer: "/data/attributes/" + fv.Field,
					//},
					Meta: mainErr.Meta,
				})
			}

		case *errdetails.PreconditionFailure:
			// Собираем preconditions в meta.preconditions
			pre := make([]map[string]string, len(info.Violations))
			for i, v := range info.Violations {
				pre[i] = map[string]string{
					"type":        v.Type,
					"subject":     v.Subject,
					"description": v.Description,
				}
			}
			(*mainErr.Meta)["preconditions"] = pre
		}
	}

	// Собираем итоговый список ошибок
	if len(errs) == 0 {
		errs = []*jsonapi.ErrorObject{mainErr}
	} else {
		errs = append([]*jsonapi.ErrorObject{mainErr}, errs...)
	}

	httpkit.RenderErr(w, errs...)
}

// mapCodeToHTTP возвращает HTTP-код для данного gRPC-кода.
func mapCodeToHTTP(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusBadGateway
	}
}

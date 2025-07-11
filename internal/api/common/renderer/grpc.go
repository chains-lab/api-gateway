package renderer

import (
	"net/http"
	"strings"
	"time"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/jsonapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RenderGRPCError(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		InternalError(w, nil)
		return
	}
	httpCode := mapCodeToHTTP(st.Code())

	var (
		requestID    string
		timestamp    string
		reason       string
		domain       string
		ri           *errdetails.ResourceInfo
		descriptions []string
	)

	for _, d := range st.Details() {
		switch info := d.(type) {
		case *errdetails.RequestInfo:
			requestID = info.RequestId

		case *errdetails.ErrorInfo:
			reason = info.Reason
			domain = info.Domain
			if t, ok := info.Metadata["timestamp"]; ok {
				timestamp = t
			}

		case *errdetails.ResourceInfo:
			ri = info
			descriptions = append(descriptions, info.Description)

		case *errdetails.BadRequest:
			for _, fv := range info.FieldViolations {
				descriptions = append(descriptions, fv.Description)
			}

		case *errdetails.PreconditionFailure:
			for _, v := range info.Violations {
				descriptions = append(descriptions, v.Description)
			}
		}
	}

	if reason == "" {
		reason = st.Code().String()
	}
	if timestamp == "" {
		timestamp = time.Now().UTC().Format(time.RFC3339Nano)
	}
	if len(descriptions) == 0 {
		descriptions = append(descriptions, st.Message())
	}
	detail := strings.Join(descriptions, "; ")

	e := &jsonapi.ErrorObject{
		Status: http.StatusText(httpCode),
		Title:  st.Message(),
		Code:   reason,
		Detail: detail,
	}

	meta := map[string]interface{}{
		"request_id": requestID,
		"timestamp":  timestamp,
		"domain":     domain,
	}

	if ri != nil {
		meta["resource_type"] = ri.ResourceType
		meta["resource_name"] = ri.ResourceName
		meta["resource_description"] = ri.Description
	}

	e.Meta = &meta

	httpkit.RenderErr(w, e)
}

func mapCodeToHTTP(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK

	case codes.Canceled:
		return http.StatusRequestTimeout

	case codes.Unknown,
		codes.Internal:
		return http.StatusInternalServerError

	case codes.InvalidArgument,
		codes.OutOfRange:
		return http.StatusBadRequest

	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout

	case codes.NotFound:
		return http.StatusNotFound

	case codes.AlreadyExists:
		return http.StatusConflict

	case codes.PermissionDenied:
		return http.StatusForbidden

	case codes.ResourceExhausted:
		return http.StatusTooManyRequests

	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed

	case codes.Aborted:
		return http.StatusConflict

	case codes.Unimplemented:
		return http.StatusNotImplemented

	case codes.Unavailable:
		return http.StatusServiceUnavailable

	case codes.Unauthenticated:
		return http.StatusUnauthorized

	default:
		return http.StatusBadGateway
	}
}

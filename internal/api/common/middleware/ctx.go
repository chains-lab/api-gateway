package middleware

import (
	"context"
	"net/http"

	"github.com/chains-lab/proto-storage/gen/go/auth"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	LogCtxKey ctxKey = iota

	ChainsAutHCtxKey
)

func CtxMiddleWare(extenders ...func(context.Context) context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for _, extender := range extenders {
				ctx = extender(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CtxLog(log *logrus.Logger) func(context.Context) context.Context {
	return func(context.Context) context.Context {
		return context.WithValue(context.Background(), LogCtxKey, log)
	}
}

func ChainsAuthCtx(service auth.ServiceClient) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ChainsAutHCtxKey, service)
	}
}

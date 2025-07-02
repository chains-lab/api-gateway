package middleware

import (
	"context"
	"net/http"

	elcab "github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
	"github.com/chains-lab/proto-storage/gen/go/svc/sso"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	LogCtxKey ctxKey = iota

	ChainsSsoUserCtxKey
	ChainsSsoAdminCtxKey
	ChainsElectorCabUserCtxKey
	ChainsElectorCabAdminCtxKey
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
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, LogCtxKey, log)
	}
}

func ChainsSsoSvcAdminCtx(service sso.AdminServiceClient) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ChainsSsoAdminCtxKey, service)
	}
}

func ChainsSsoSvcUserCtx(service sso.UserServiceClient) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ChainsSsoUserCtxKey, service)
	}
}

func ChainsElectorCabSvcUserCtx(service elcab.UserServiceClient) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ChainsElectorCabUserCtxKey, service)
	}
}

func ChainsElectorCabSvcAdminCtx(service elcab.AdminServiceClient) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ChainsElectorCabAdminCtxKey, service)
	}
}

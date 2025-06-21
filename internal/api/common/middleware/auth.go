package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
)

func AuthMdl(skUser, skService string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
					Status: http.StatusUnauthorized,
					Code:   "MISSING_AUTHORIZATION_HEADER",
					Detail: "Missing Authorization header",
				})...)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
					Status: http.StatusUnauthorized,
					Code:   "MISSING_AUTHORIZATION_HEADER",
					Detail: "Missing Authorization header",
				})...)
				return
			}

			tokenString := parts[1]

			userData, err := tokens.VerifyUserJWT(r.Context(), tokenString, skUser)
			if err != nil {
				httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
					Status: http.StatusUnauthorized,
					Code:   "TOKEN_VALIDATION_FAILED",
					Detail: "Token validation failed",
				})...)
				return
			}

			ctx = context.WithValue(ctx, tokens.SubjectIDKey, userData.Subject)
			ctx = context.WithValue(ctx, tokens.SessionIDKey, userData.Session)
			ctx = context.WithValue(ctx, tokens.SubscriptionKey, userData.Subscription)
			ctx = context.WithValue(ctx, tokens.RoleKey, userData.Role)
			ctx = context.WithValue(ctx, tokens.VerifiedKey, userData.Verified)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

package auth

import (
	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/api-gateway/internal/api/services/auth/handlers"
	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/gatekit/roles"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, cfg config.Config) chi.Router {
	auth := middleware.AuthMdl(cfg.JWT.User.AccessToken.SecretKey, cfg.JWT.Service.SecretKey)
	admin := middleware.RolesGrant(cfg.JWT.User.AccessToken.SecretKey, roles.Admin, roles.SuperUser)

	r.Route("/auth", func(r chi.Router) {
		r.Route("/own", func(r chi.Router) {
			r.Use(auth)

			r.Get("/", handlers.OwnUserGet)

			r.Get("/login", handlers.GoogleLogin)
			r.Get("/login_callback", handlers.GoogleCallback)
			r.Post("/refresh", handlers.Refresh)
			r.Delete("/logout", handlers.Logout)

			r.Route("/sessions", func(r chi.Router) {
				r.Route("/{session_id}", func(r chi.Router) {
					r.Get("/", handlers.OwnGetSession)
					r.Delete("/", handlers.DeleteSession)
				})

				r.Get("/", handlers.OwnGetSessions)
				r.Delete("/", handlers.OwnTerminateSessions)
			})
		})

		r.Route("/admin", func(r chi.Router) {
			r.Use(admin)

			r.Route("/{user_id}", func(r chi.Router) {
				r.Get("/", handlers.AdminGetUser)

				r.Patch("/role/{role}", handlers.AdminUpdateRole)
				r.Patch("/subscription/{subscription_id}", handlers.AdminUpdateSubscription)
				r.Patch("/verified/{verified}", handlers.AdminUpdateVerified)
				r.Patch("/suspended/{suspended}", handlers.AdminUpdateSuspended)

				r.Route("/sessions", func(r chi.Router) {
					r.Get("/", handlers.AdminGetUserSessions)
					r.Delete("/", handlers.AdminTerminateUserSessions)
					r.Route("/{session_id}", func(r chi.Router) {
						r.Get("/", handlers.AdminGetUserSession)
						r.Delete("/", handlers.AdminDeleteUserSession)
					})
				})
			})
		})
	})

	return r
}

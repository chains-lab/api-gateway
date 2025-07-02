package sso

import (
	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/handlers"
	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/gatekit/roles"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, cfg config.Config) chi.Router {
	auth := middleware.AuthMdl(cfg.JWT.User.AccessToken.SecretKey, cfg.JWT.Service.SecretKey)
	admin := middleware.RolesGrant(cfg.JWT.User.AccessToken.SecretKey, roles.Admin, roles.SuperUser)
	superUser := middleware.RolesGrant(cfg.JWT.User.AccessToken.SecretKey, roles.SuperUser)

	r.Route("/svc", func(r chi.Router) {
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

			r.With(superUser).Post("/create_user", handlers.AdminCreateUser)

			r.Route("/{user_id}", func(r chi.Router) {
				r.Get("/", handlers.AdminGetUser)

				r.Route("/sessions", func(r chi.Router) {
					r.Get("/", handlers.AdminGetUserSessions)
					r.Delete("/", handlers.AdminDeleteUserSessions)
				})
			})
		})
	})

	return r
}

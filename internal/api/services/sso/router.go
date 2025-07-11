package sso

import (
	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/api-gateway/internal/api/services/sso/handlers"
	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/gatekit/roles"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, cfg config.Config) {
	auth := middleware.AuthMdl(cfg.JWT.User.AccessToken.SecretKey, cfg.JWT.Service.SecretKey)
	admin := middleware.RolesGrant(cfg.JWT.User.AccessToken.SecretKey, roles.Admin, roles.SuperUser)
	superUser := middleware.RolesGrant(cfg.JWT.User.AccessToken.SecretKey, roles.SuperUser)

	r.Route("/google", func(r chi.Router) {
		r.Route("/login", func(r chi.Router) {
			r.Get("/", handlers.GoogleLogin)
			r.Get("/callback", handlers.GoogleCallback)
		})
	})

	r.Route("/own", func(r chi.Router) {
		r.With(auth).Get("/", handlers.GetOwnUser)

		r.Post("/refresh", handlers.Refresh)

		r.With(auth).Delete("/logout", handlers.Logout)

		r.Route("/sessions", func(r chi.Router) {
			r.Use(auth)

			r.Route("/{session_id}", func(r chi.Router) {
				r.Get("/", handlers.GetOwnUserSession)
				r.Delete("/", handlers.DeleteOwnUserSession)
			})

			r.Get("/", handlers.GetOwnUserSessions)
			r.Delete("/", handlers.DeleteOwnUserSessions)
		})
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(admin)

		r.With(superUser).Post("/create_user", handlers.CreateUserByAdmin)

		r.Route("/{user_id}", func(r chi.Router) {
			r.Get("/", handlers.GetUserByAdmin)

			r.Route("/sessions", func(r chi.Router) {
				r.Get("/", handlers.GetUserSessionsByAdmin)
				r.Delete("/", handlers.DeleteUserSessionByAdmin)
			})
		})
	})
}

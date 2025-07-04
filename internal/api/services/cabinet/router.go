package cabinet

import (
	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet/handlers"
	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/gatekit/roles"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, cfg config.Config) {
	auth := middleware.AuthMdl(cfg.JWT.User.AccessToken.SecretKey, cfg.JWT.Service.SecretKey)
	admin := middleware.RolesGrant(cfg.JWT.User.AccessToken.SecretKey, roles.Admin, roles.SuperUser)

	r.Route("/cabinets", func(r chi.Router) {
		r.With(auth).Post("/", handlers.CreateOwnProfile)

		r.Route("/own", func(r chi.Router) {
			r.Use(auth)

			r.Get("/", handlers.GetOwnCabinet)
		})

		r.Get("/", handlers.GetCabinet)
	})

	r.Route("/profile", func(r chi.Router) {
		r.Route("/own", func(r chi.Router) {
			r.Use(auth)

			r.Patch("/username/{username}", handlers.UpdateOwnUsername)

			r.Patch("/", handlers.UpdateOwnProfile)
			r.Get("/", handlers.GetOwnProfile)
		})

		//r.Get("/search", handlers.SearchProfile)
		r.Get("/", handlers.GetProfile)
	})

	r.Route("/biographies", func(r chi.Router) {
		r.Use(auth)

		r.Patch("/sex/{sex}", handlers.UpdateOwnSex)
		r.Patch("/birthday/birthday", handlers.UpdateOwnBirthday)
		r.Patch("/nationality/{nationality}", handlers.UpdateOwnNationality)
		r.Patch("/primary_language/{primary_language}", handlers.UpdateOwnPrimaryLanguage)
		r.Patch("/residence/{country}/{region}/{city}", handlers.UpdateOwnResidence)
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(admin)

		r.Route("/{cabinet_id}", func(r chi.Router) {
			r.Route("/profile", func(r chi.Router) {
				r.Patch("/username", handlers.ResetUsernameByAdmin)
				r.Patch("/official/{official}", handlers.UpdateOfficialByAdmin)

				r.Patch("/", handlers.ResetProfileByAdmin)
			})
		})
	})
}

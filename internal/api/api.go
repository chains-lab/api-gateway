package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/chains-lab/api-gateway/internal/api/common/middleware"
	"github.com/chains-lab/api-gateway/internal/api/services/cabinet"
	"github.com/chains-lab/api-gateway/internal/api/services/sso"
	"github.com/chains-lab/api-gateway/internal/config"
	"github.com/chains-lab/api-gateway/internal/grpcclients"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Api struct {
	server *http.Server
	router *chi.Mux
	log    *logrus.Entry
	cfg    config.Config
}

func NewAPI(cfg config.Config) Api {
	logger := cfg.GetLogger().WithField("module", "api")
	router := chi.NewRouter()
	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: router,
	}

	return Api{
		server: server,
		router: router,
		log:    logger,
		cfg:    cfg,
	}
}

func (a Api) Run(ctx context.Context) {
	r := a.router

	r.Use(
		middleware.CtxMiddleWare(
			middleware.CtxLog(a.cfg.GetLogger()),
			middleware.ChainsSsoSvcAdminCtx(grpcclients.SsoSvcAdmin(a.cfg)),
			middleware.ChainsSsoSvcUserCtx(grpcclients.SsoSvcUser(a.cfg)),
			middleware.ChainsElectorCabSvcUserCtx(grpcclients.ElectorCabUserSvc(a.cfg)),
			middleware.ChainsElectorCabSvcAdminCtx(grpcclients.ElectorCabAdminSvc(a.cfg)),
		),
	)

	r.Route("/sso", func(r chi.Router) {
		sso.Router(r, a.cfg)
	})

	r.Route("/elector-cab", func(r chi.Router) {
		cabinet.Router(r, a.cfg)
	})

	a.Start(ctx)

	<-ctx.Done()
	a.Stop(ctx)
}

func (a Api) Start(ctx context.Context) {
	go func() {
		a.log.Info("Starting server...")
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Fatalf("Server failed to start: %v", err)
		}
	}()
}

func (a Api) Stop(ctx context.Context) {
	a.log.Info("Shutting down server...")
	if err := a.server.Shutdown(ctx); err != nil {
		a.log.Errorf("Server shutdown failed: %v", err)
	}
}

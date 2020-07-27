package api

import (
	"context"
	"net/http"
	"time"

	"user_api/api/handler"
	"user_api/config"
	"user_api/repo"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"
)

func GetServer(cfg *config.Cfg, logger *logrus.Entry, ) *api.Server {
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
		log.NewRequestLogger(logger.Logger),
	)

	if cfg.API.EnableCORS {
		mux.Use(getCORS().Handler)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if cfg.API.ApiRequestTimeout > 0 {
		t := time.Duration(cfg.API.ApiRequestTimeout)
		mux.Use(middleware.Timeout(t * time.Second))
	}

	mongoRepo, err := repo.NewMongoRepo(context.TODO(), cfg.Mongo)
	if err != nil {
		logger.WithError(err).Fatal("unable to init mongo repo")
	}

	h := handler.NewHandler(mongoRepo, logger)
	mux.Route("/", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			render.Success(w, config.AppInfo())
		})
		r.Post("/create_account", h.CreateUser)
		r.Post("/authenticate", h.Authenticate)

	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	logger.Info("router initialized; starting API server...")
	return api.NewServer(cfg.API, mux)
}

func getCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "jwt", "X-UID"},
		ExposedHeaders:   []string{"Link", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}

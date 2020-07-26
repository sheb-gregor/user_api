package handler

import (
	"user_api/config"
	"user_api/repo"

	"github.com/sirupsen/logrus"
)

// Handler contains realization of http handlers
type Handler struct {
	couch *repo.CouchRepo
	log   *logrus.Entry
}

func NewHandler(cfg *config.Cfg, entry *logrus.Entry, ) *Handler {

	return &Handler{
		couch: repo.NewCouchRepo(cfg.CouchDB),
		log:   entry.WithField("app_layer", "api.Handler"),
	}

}

package handler

import (
	"user_api/repo"

	"github.com/sirupsen/logrus"
)

// Handler contains realization of http handlers
type Handler struct {
	mongoRepo *repo.MongoRepo
	log       *logrus.Entry
}

func NewHandler(repo *repo.MongoRepo, entry *logrus.Entry, ) *Handler {
	return &Handler{
		mongoRepo: repo,
		log:       entry.WithField("app_layer", "api.Handler"),
	}
}

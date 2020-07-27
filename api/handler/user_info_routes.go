package handler

import (
	"net/http"

	"user_api/api/resources"
	"user_api/repo/models"

	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.log, r)

	data := new(resources.UserInfoRequest)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	logger.Debug("Trying to write data into mongo")
	userInfo, err := h.mongoRepo.UserInfo(r.Context())
	if err != nil {
		logger.WithError(err).Error("Can not establish connection with database")
		render.ServerError(w)
		return
	}

	err = userInfo.AddUserInfo(models.NewUserInfo(data.Email, data.Password))
	if err != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.ResultSuccess.Render(w)
}

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.log, r)

	request := new(resources.UserInfoRequest)
	err := httpx.ParseJSONBody(r, request)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	userInfo, err := h.mongoRepo.UserInfo(r.Context())
	if err != nil {
		logger.WithError(err).Error("unable to create custom doc")
		render.ServerError(w)
		return
	}

	record, err := userInfo.GetUserInfo(request.Email)
	if err != nil {
		logger.WithError(err).Error("can not to get document")
		render.ServerError(w)
		return
	}

	if !models.ValidatePassword(request.Password, record.PasswordHash) {
		render.Unauthorized(w, "invalid password")
		return
	}

	render.Success(w, record)
}

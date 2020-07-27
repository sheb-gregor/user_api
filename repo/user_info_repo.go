package repo

import (
	"context"

	"user_api/repo/models"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInfoRepoI interface {
	AddUserInfo(info *models.UserInfo) error
	GetUserInfo(email string) (*models.UserInfo, error)
}

type UserInfoRepo struct {
	ctx context.Context

	collection *mongo.Collection
}

func (repo *UserInfoRepo) AddUserInfo(info *models.UserInfo) error {
	_, err := repo.collection.InsertOne(repo.ctx, info)
	if err != nil {
		return errors.Wrap(err, "unable to add UserInfo")
	}
	return nil
}

func (repo *UserInfoRepo) GetUserInfo(email string) (*models.UserInfo, error) {
	filter := bson.D{{"email", email}}
	result := new(models.UserInfo)

	err := repo.collection.FindOne(repo.ctx, filter).Decode(result)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find UserInfo")
	}

	return result, nil
}

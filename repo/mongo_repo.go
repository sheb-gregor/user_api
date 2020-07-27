package repo

import (
	"context"
	"fmt"

	"user_api/config"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	ctx      context.Context
	client   *mongo.Client
	database string
}

func NewMongoRepo(ctx context.Context, conf config.MongoConf) (*MongoRepo, error) {
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d/", conf.Host, conf.Port))
	// SetAuth(options.Credential{
	// 	Username: conf.Username.Get(),
	// 	Password: conf.Password.Get(),
	// })

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to mongo")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to ping mongo")
	}
	return &MongoRepo{
		ctx:      ctx,
		client:   client,
		database: conf.Database,
	}, nil
}

func (repo *MongoRepo) UserInfo(ctx context.Context) (UserInfoRepoI, error) {
	collection := repo.client.Database(repo.database).Collection("users")
	return &UserInfoRepo{collection: collection, ctx: ctx}, nil
}

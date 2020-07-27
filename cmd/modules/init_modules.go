package modules

import (
	"context"
	"fmt"
	"time"

	"user_api/config"

	"github.com/lancer-kit/armory/initialization"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(c *cli.Context) (*config.Cfg, error) {
	cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config")
	}

	err = getModules(cfg).InitAll()
	if err != nil {
		return nil, errors.Wrap(err, "modules initialization failed")
	}

	return &cfg, nil
}

func getModules(cfg config.Cfg) initialization.Modules {
	return initialization.Modules{
		initialization.Module{
			Name:         "mongo",
			DependsOn:    "",
			Timeout:      time.Duration(cfg.ServicesInitTimeout) * time.Second,
			InitInterval: 500 * time.Millisecond,
			Init: func(entry *logrus.Entry) error {
				var ctx = context.TODO()

				clientOptions := options.Client().
					ApplyURI(fmt.Sprintf("mongodb://%s:%d/", cfg.Mongo.Host, cfg.Mongo.Port))
					// SetAuth(options.Credential{
						// AuthMechanism: "PLAIN",
						// Username: cfg.Mongo.Username.Get(),
						// Password: cfg.Mongo.Password.Get(),
					// })

				client, err := mongo.Connect(ctx, clientOptions)
				if err != nil {
					return errors.Wrap(err, "unable to connect to mongo")
				}

				err = client.Ping(ctx, nil)
				if err != nil {
					return errors.Wrap(err, "unable to ping mongo")
				}

				return nil

			},
		},
	}

}

package modules

import (
	"time"

	"user_api/config"

	"github.com/lancer-kit/armory/initialization"
	cdb "github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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
			Name:         "couchdb",
			DependsOn:    "",
			Timeout:      time.Duration(cfg.ServicesInitTimeout) * time.Second,
			InitInterval: 500 * time.Millisecond,
			Init: func(entry *logrus.Entry) error {
				_, err := cdb.NewDatabase(cfg.CouchDB)
				if err != nil {
					return errors.Wrap(err, "unable to connect to couchdb")
				}
				return nil
			},
		},
	}

}

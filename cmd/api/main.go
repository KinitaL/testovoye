package main

import (
	"context"
	"github.com/KinitaL/testovoye/config"
	"github.com/KinitaL/testovoye/internal/app/api"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log, err := newLogger(c.Service.Development)
	if err != nil {
		return
	}
	defer log.Sync() //nolint:errcheck

	app := api.NewApp(c, log)

	if err := app.Configure(ctx); err != nil {
		log.Error("cannot configure app", zap.Error(err))
		return
	}

	if err := app.Run(); err != nil {
		log.Error("cannot run app", zap.Error(err))
		return
	}
}

func newLogger(development bool) (*zap.Logger, error) {
	var lConf zap.Config
	if development {
		lConf = zap.NewDevelopmentConfig()
	} else {
		lConf = zap.NewProductionConfig()
	}

	lConf.Encoding = "json"

	zLogger, err := lConf.Build()
	if err != nil {
		return nil, err
	}

	return zLogger, nil
}

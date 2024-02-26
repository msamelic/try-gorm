package main

import (
	"try-gorm/internal/repository"
	"try-gorm/internal/service/ginrouter"
	"try-gorm/internal/util/envconf"
	"try-gorm/internal/util/zaplog"
)

func main() {
	env := envconf.New()
	log := zaplog.New()

	repo, err := repository.New(env.SqliteDsn, log)
	if err != nil {
		log.Sugar().Panic(err)
	}

	r := ginrouter.New(repo, env, log)
	log.Info(env.AppName + " is starting")

	if err := r.Run(":" + env.Port); err != nil {
		log.Sugar().Panic(err)
	}
}

/*
https://gorm.io/docs/
*/

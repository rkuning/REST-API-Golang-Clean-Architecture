package main

import (
	"context"
	"rest-api-golang-clean-code/bootstrap"
	"rest-api-golang-clean-code/internal/config"
	"time"
)

func main() {
	env, _ := config.NewEnv()
	logger := config.NewLogger(env.LogLevel)
	db := config.NewDatabase(env, logger)
	validator := config.NewValidator()
	router := config.NewRouter(env.BasePrefixUrl)

	boostrapConfig := &bootstrap.BootstrapConfig{
		DB:       db,
		Router:   router,
		Log:      logger,
		Validate: validator,
		Env:      env,
		Ctx:      context.Background(),
		Timeout:  time.Duration(env.ContextTimeout) * time.Second,
	}

	bootstrap.Bootstrap(boostrapConfig)
	bootstrap.ExecuteApp(env, logger, router)
}

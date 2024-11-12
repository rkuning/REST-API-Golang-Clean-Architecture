package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"rest-api-golang-clean-code/internal/config"
	"rest-api-golang-clean-code/internal/delivery/router"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	Router   *mux.Router
	Log      *logrus.Logger
	Validate *validator.Validate
	Env      *config.Env
	Ctx      context.Context
	Timeout  time.Duration
}

func Bootstrap(config *BootstrapConfig) {
	routerConfig := router.RouterConfig{
		Router: config.Router,
		Env:    config.Env,
		Logger: config.Log,
	}

	routerConfig.SetupRouter()
}

func ExecuteApp(env *config.Env, log *logrus.Logger, r *mux.Router) {
	configureLogger(log)

	server := createServer(env, r)

	fmt.Printf("server running on port %s\n", env.ServerPort)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func configureLogger(log *logrus.Logger) {
	config.ConfigureLogFileOutput(log)
}

func createServer(env *config.Env, r *mux.Router) *http.Server {
	return &http.Server{
		Handler:      r,
		WriteTimeout: time.Duration(env.ContextTimeout) * time.Second,
		ReadTimeout:  time.Duration(env.ContextTimeout) * time.Second,
		IdleTimeout:  time.Duration(env.ContextTimeout) * time.Second,
		Addr:         fmt.Sprintf("%s:%s", env.ServerAddress, env.ServerPort),
	}
}

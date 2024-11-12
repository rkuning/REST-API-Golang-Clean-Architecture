package router

import (
	"net/http"
	"rest-api-golang-clean-code/internal/config"
	"rest-api-golang-clean-code/internal/delivery/middleware"
	"rest-api-golang-clean-code/internal/model"
	"rest-api-golang-clean-code/internal/util"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type RouterConfig struct {
	Router *mux.Router
	Env    *config.Env
	Logger *logrus.Logger
}

func (rc *RouterConfig) SetupRouter() {
	rc.Router.Use(middleware.CorsMiddleware, middleware.ErrorMiddleware(rc.Logger, rc.Env.ContextTimeout))

	rc.Router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		util.WriteJSON(w, http.StatusOK, model.WebResponse[any]{
			Code:    http.StatusOK,
			Message: "success",
		})
	}).Methods(http.MethodGet).Name("Base Rest API Golang Clean Architecture")
}

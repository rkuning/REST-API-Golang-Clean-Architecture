package config

import "github.com/gorilla/mux"

func NewRouter(prefix string) *mux.Router {
	return mux.NewRouter().PathPrefix(prefix).Subrouter()
}

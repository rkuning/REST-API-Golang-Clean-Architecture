package middleware

import (
	"fmt"
	"net/http"
	"rest-api-golang-clean-code/internal/util"
	"time"

	"github.com/sirupsen/logrus"
)

func ErrorMiddleware(logger *logrus.Logger, timeout int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			done := make(chan struct{})
			panicChanel := make(chan any)

			go func() {
				defer func() {
					if err := recover(); err != nil {
						panicChanel <- err
					}
				}()

				next.ServeHTTP(w, r)
				close(done)
			}()

			select {
			case <-done:
				return
			case err := <-panicChanel:
				handlePanic(logger, w, err)
			case <-time.After(time.Duration(timeout) * time.Second):
				util.WriteJSON(w, http.StatusGatewayTimeout, util.TimeoutError())
			}
		})
	}
}

func handlePanic(log *logrus.Logger, w http.ResponseWriter, err any) {
	if apiErr, ok := err.(util.ApiError); ok {
		log.WithError(fmt.Errorf("%v", err)).Error(apiErr.Message)
		util.WriteJSON(w, apiErr.Code, apiErr)
	} else {
		log.WithError(fmt.Errorf("%v", err)).Error("internal server error")
		util.WriteJSON(w, http.StatusInternalServerError, util.InternalServerError())
	}
}

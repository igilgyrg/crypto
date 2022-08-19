package error

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, request *http.Request) (interface{}, error)

func MiddleWare(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var appError *Error
		_, err := h(w, r)
		if err != nil {
			if errors.As(err, &appError) {
				w.WriteHeader(appError.HttpCode)
				w.Write(appError.Marshal())
				return
			}

			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err).Marshal())
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

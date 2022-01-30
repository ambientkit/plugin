package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	handlerFunc     func(w http.ResponseWriter, r *http.Request) (status int, err error)
	customServeHTTP func(w http.ResponseWriter, r *http.Request, status int, err error)
	Handle          func(http.ResponseWriter, *http.Request, httprouter.Params)
}

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn.handlerFunc(w, r)

	if fn.customServeHTTP == nil {
		return
	}

	fn.customServeHTTP(w, r, status, err)
}

package ambhandler

import (
	"net/http"
)

// Handler represents an Ambient handler.
type Handler struct {
	HandlerFunc     func(w http.ResponseWriter, r *http.Request) (err error)
	CustomServeHTTP func(w http.ResponseWriter, r *http.Request, err error)
}

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn.HandlerFunc(w, r)

	if fn.CustomServeHTTP == nil {
		return
	}

	fn.CustomServeHTTP(w, r, err)
}

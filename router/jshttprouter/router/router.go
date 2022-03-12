// Package router provides request handling capabilities.
package router

import (
	"net/http"

	"github.com/ambientkit/ambient"
	"github.com/julienschmidt/httprouter"
)

// Mux contains the router.
type Mux struct {
	router *httprouter.Router

	// customServeHTTP is the serve function.
	customServeHTTP func(w http.ResponseWriter, r *http.Request, err error)
}

// New returns an instance of the router.
func New() *Mux {
	r := httprouter.New()

	return &Mux{
		router: r,
	}
}

// SetServeHTTP sets the ServeHTTP function.
func (m *Mux) SetServeHTTP(csh func(w http.ResponseWriter, r *http.Request, err error)) {
	m.customServeHTTP = csh
}

// SetNotFound sets the NotFound function.
func (m *Mux) SetNotFound(notFound http.Handler) {
	m.router.NotFound = notFound
}

// ServeHTTP routes the incoming http.Request based on method and path
// extracting path parameters as it goes.
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

// StatusError returns error with a status code.
func (m *Mux) StatusError(status int, err error) error {
	return ambient.StatusError{Code: status, Err: err}
}

// Error shows error page based on the status code.
func (m *Mux) Error(status int, w http.ResponseWriter, r *http.Request) {
	if m.customServeHTTP != nil {
		m.customServeHTTP(w, r, ambient.StatusError{Code: status, Err: nil})
		return
	}

	http.Error(w, http.StatusText(status), status)
}

// Param returns a URL parameter.
func (m *Mux) Param(r *http.Request, param string) string {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(param)
}

// Wrap a standard http handler so it can be used easily.
func (m *Mux) Wrap(handler http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) (err error) {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		handler.ServeHTTP(w, r)
		return
	}
}

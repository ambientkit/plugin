// Package router provides request handling capabilities.
package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Mux contains the router.
type Mux struct {
	router *chi.Mux

	// customServeHTTP is the serve function.
	customServeHTTP func(w http.ResponseWriter, r *http.Request, status int, err error)
	notFound        http.Handler
}

// New returns an instance of the router.
func New() *Mux {
	r := chi.NewRouter()

	return &Mux{
		router:   r,
		notFound: r.NotFoundHandler(),
	}
}

// SetServeHTTP sets the ServeHTTP function.
func (m *Mux) SetServeHTTP(csh func(w http.ResponseWriter, r *http.Request, status int, err error)) {
	m.customServeHTTP = csh
}

// SetNotFound sets the NotFound function.
func (m *Mux) SetNotFound(notFound http.Handler) {
	m.notFound = notFound
	m.router.NotFound(notFound.ServeHTTP)
}

// Clear will remove a method and path from the router.
func (m *Mux) Clear(method string, path string) {
	// Overwrite instead of delete.
	m.router.Method(method, path, m.notFound)
}

// ServeHTTP routes the incoming http.Request based on method and path
// extracting path parameters as it goes.
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

// Error shows error page based on the status code.
func (m *Mux) Error(status int, w http.ResponseWriter, r *http.Request) {
	m.customServeHTTP(w, r, status, nil)
}

// Param returns a URL parameter.
func (m *Mux) Param(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}

// Wrap a standard http handler so it can be used easily.
func (m *Mux) Wrap(handler http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) (status int, err error) {
	return func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		handler.ServeHTTP(w, r)
		return
	}
}

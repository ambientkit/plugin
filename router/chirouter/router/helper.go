package router

import (
	"net/http"

	"github.com/ambientkit/plugin/pkg/ambhandler"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (m *Mux) Delete(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("DELETE", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (m *Mux) Get(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("GET", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (m *Mux) Head(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("HEAD", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (m *Mux) Options(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("OPTIONS", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (m *Mux) Patch(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("PATCH", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (m *Mux) Post(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("POST", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (m *Mux) Put(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Method("PUT", path, ambhandler.Handler{
		HandlerFunc:     fn,
		CustomServeHTTP: m.customServeHTTP,
	})
}

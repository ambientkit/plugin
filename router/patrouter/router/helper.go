package router

import (
	"net/http"

	"github.com/ambientkit/plugin/pkg/paramconvert"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (m *Mux) Delete(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("DELETE", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (m *Mux) Get(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("GET", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (m *Mux) Head(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("HEAD", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (m *Mux) Options(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("OPTIONS", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (m *Mux) Patch(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("PATCH", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (m *Mux) Post(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("POST", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (m *Mux) Put(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Add("PUT", paramconvert.BraceToColon(path), handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	})
}

package router

import (
	"net/http"
)

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (m *Mux) Delete(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("DELETE", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (m *Mux) Get(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("GET", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (m *Mux) Head(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("HEAD", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (m *Mux) Options(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("OPTIONS", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (m *Mux) Patch(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("PATCH", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (m *Mux) Post(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("POST", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (m *Mux) Put(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.HandlerFunc("PUT", path, func(w http.ResponseWriter, req *http.Request) {
		handler{
			handlerFunc:     fn,
			customServeHTTP: m.customServeHTTP,
		}.ServeHTTP(w, req)
	})
}

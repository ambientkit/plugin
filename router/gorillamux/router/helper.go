package router

import "net/http"

// Delete is a shortcut for router.Handle("DELETE", path, handle)
func (m *Mux) Delete(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("DELETE")
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (m *Mux) Get(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("GET")
}

// Head is a shortcut for router.Handle("HEAD", path, handle)
func (m *Mux) Head(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("HEAD")
}

// Options is a shortcut for router.Handle("OPTIONS", path, handle)
func (m *Mux) Options(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("OPTIONS")
}

// Patch is a shortcut for router.Handle("PATCH", path, handle)
func (m *Mux) Patch(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("PATCH")
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (m *Mux) Post(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("POST")
}

// Put is a shortcut for router.Handle("PUT", path, handle)
func (m *Mux) Put(path string, fn func(http.ResponseWriter, *http.Request) (int, error)) {
	m.router.Handle(path, handler{
		handlerFunc:     fn,
		customServeHTTP: m.customServeHTTP,
	}).Methods("PUT")
}

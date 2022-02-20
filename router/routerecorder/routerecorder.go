// Package routerecorder keeps track of each of the routes a plugin adds to the
// router. It is not a functioning router.
package routerecorder

import (
	"net/http"

	"github.com/ambientkit/ambient"
)

// Route represents a route.
type Route struct {
	Method string
	Path   string
}

// Plugin represents an Ambient plugin.
type Plugin struct {
	routeList []Route
}

// New returns an Ambient plugin for a router that records routes.
func New() *Plugin {
	return &Plugin{}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "routerecorder"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Routes returns list of routes.
func (p *Plugin) Routes() []Route {
	return p.routeList
}

// handleRoute records each of the routes.
func (p *Plugin) handleRoute(path string, method string) {
	p.routeList = append(p.routeList, Route{
		Method: method,
		Path:   path,
	})
}

// Router returns a router.
func (p *Plugin) Router(logger ambient.Logger, te ambient.Renderer) (ambient.AppRouter, error) {
	return p, nil
}

// Clear will remove a method and path from the router.
func (p *Plugin) Clear(method string, path string) {}

// SetServeHTTP sets the ServeHTTP function.
func (p *Plugin) SetServeHTTP(h func(w http.ResponseWriter, r *http.Request, status int, err error)) {
}

// ServeHTTP routes the incoming http.Request based on method and path
// extracting path parameters as it goes.
func (p *Plugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// SetNotFound sets the NotFound function.
func (p *Plugin) SetNotFound(notFound http.Handler) {}

// Get registers a pattern with the router.
func (p *Plugin) Get(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodGet)
}

// Post registers a pattern with the router.
func (p *Plugin) Post(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodPost)
}

// Patch registers a pattern with the router.
func (p *Plugin) Patch(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodPatch)
}

// Put registers a pattern with the router.
func (p *Plugin) Put(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodPut)
}

// Head registers a pattern with the router.
func (p *Plugin) Head(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodHead)
}

// Options registers a pattern with the router.
func (p *Plugin) Options(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodOptions)
}

// Delete registers a pattern with the router.
func (p *Plugin) Delete(path string, fn func(http.ResponseWriter, *http.Request) (status int, err error)) {
	p.handleRoute(path, http.MethodDelete)
}

// Param returns a URL parameter.
func (p *Plugin) Param(r *http.Request, name string) string {
	return ""
}

// Error shows error page based on the status code.
func (p *Plugin) Error(status int, w http.ResponseWriter, r *http.Request) {}

// Wrap a standard http handler so it can be used easily.
func (p *Plugin) Wrap(handler http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) (status int, err error) {
	return nil
}

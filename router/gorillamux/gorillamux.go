// Package gorillamux is an Ambient plugin for a router using gorilla/mux.
package gorillamux

import (
	"context"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/router/gorillamux/router"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	serveHTTP ambient.CustomServeHTTP
}

// New returns an Ambient plugin for a router using a variation of the way router.
// A custom CustomServeHTTP can be passed in to override how errors are handled.
func New(serveHTTP ambient.CustomServeHTTP) *Plugin {
	return &Plugin{
		serveHTTP: serveHTTP,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "gorillamux"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// Router returns a router.
func (p *Plugin) Router(_ context.Context, logger ambient.Logger, te ambient.Renderer) (ambient.AppRouter, error) {
	// Set up the default router.
	mux := router.New()

	// Set the NotFound and custom ServeHTTP handlers.
	ambient.SetupRouter(logger, mux, te, p.serveHTTP)

	return mux, nil
}

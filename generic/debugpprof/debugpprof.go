// Package debugpprof is an Ambient plugin that provides pprof functionality.
package debugpprof

import (
	"context"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that provides pprof functionality.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "debugpprof"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterRouteWrite, Description: "Access to create routes for accessing debug information."},
	}
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes() {
	p.Mux.Get("/debug/pprof", p.Mux.Wrap(p.index))
	p.Mux.Get("/debug/pprof/{pprof}", p.Mux.Wrap(p.profile))
}

// Index shows the profile index.
func (p *Plugin) index(w http.ResponseWriter, r *http.Request) {
	// pprof requires a trailing slash to work properly.
	if !strings.HasSuffix(r.URL.Path, "/") {
		p.Redirect(w, r, "/debug/pprof/", http.StatusFound)
		return
	}
	pprof.Index(w, r)
}

// Profile shows the individual profiles.
func (p *Plugin) profile(w http.ResponseWriter, r *http.Request) {
	switch p.Mux.Param(r, "pprof") {
	case "cmdline":
		pprof.Cmdline(w, r)
	case "profile":
		pprof.Profile(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	case "trace":
		pprof.Trace(w, r)
	default:
		pprof.Index(w, r)
	}
}

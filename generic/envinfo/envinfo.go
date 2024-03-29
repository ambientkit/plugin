// Package envinfo is an Ambient plugin that provides a dashboard page showing env variables.
package envinfo

import (
	"embed"

	"github.com/ambientkit/ambient"
)

//go:embed template/*.tmpl
var assets embed.FS

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that provides a dashboard page showing env variables.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "envinfo"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterRouteWrite, Description: "Access to make env variable page available at: /dashboard/env"},
	}
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes() {
	p.Mux.Get("/dashboard/env", p.showEnv)
}

// Package robots is an Ambient plugin that serves a robots.txt file.
package robots

import "github.com/ambientkit/ambient"

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that serves a robots.txt file.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "robots"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes() {
	p.Mux.Get("/robots.txt", p.index)
}

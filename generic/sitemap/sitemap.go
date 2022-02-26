// Package sitemap is an Ambient plugin that provides a sitemap.
package sitemap

import "github.com/ambientkit/ambient"

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that provides a sitemap.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "sitemap"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantSiteURLRead, Description: "Access to read the site URL."},
		{Grant: ambient.GrantSiteSchemeRead, Description: "Access to read the site scheme."},
		{Grant: ambient.GrantSiteUpdatedRead, Description: "Access to read the last updated date."},
		{Grant: ambient.GrantSitePostRead, Description: "Access to read all the posts."},
	}
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes() {
	p.Mux.Get("/sitemap.xml", p.index)
}

// Package pluginmanager is an Ambient plugin that provides a plugin management system.
package pluginmanager

import (
	"context"
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/ambientkit/ambient"
)

//go:embed template/*.tmpl
var assets embed.FS

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that provides a plugin management system.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "pluginmanager"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantSitePluginRead, Description: "Access to read the plugins."},
		{Grant: ambient.GrantSitePluginEnable, Description: "Access to enable plugins."},
		{Grant: ambient.GrantSitePluginDisable, Description: "Access to disable plugins."},
		{Grant: ambient.GrantSitePluginDelete, Description: "Access to delete plugin storage."},
		{Grant: ambient.GrantSiteFuncMapWrite, Description: "Access add FuncMap for template helpers."},
		{Grant: ambient.GrantPluginNeighborSettingRead, Description: "Access to read other plugin settings."},
		{Grant: ambient.GrantPluginNeighborSettingWrite, Description: "Access to write to other plugin settings"},
		{Grant: ambient.GrantPluginNeighborGrantRead, Description: "Access to read grant requests for plugins"},
		{Grant: ambient.GrantPluginNeighborGrantWrite, Description: "Access to approve grants for plugins."},
		{Grant: ambient.GrantPluginNeighborRouteRead, Description: "Access to read routes for plugins."},
		{Grant: ambient.GrantRouterRouteWrite, Description: "Access to create routes for editing the plugins."},
		{Grant: ambient.GrantPluginTrustedRead, Description: "Access to read if a plugin is trusted or not."},
	}
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes(context.Context) {
	p.Mux.Get("/dashboard/plugins", p.edit)
	p.Mux.Post("/dashboard/plugins", p.update)
	p.Mux.Get("/dashboard/plugins/{id}/delete", p.destroy)
	p.Mux.Get("/dashboard/plugins/{id}/settings", p.settingsEdit)
	p.Mux.Post("/dashboard/plugins/{id}/settings", p.settingsUpdate)
	p.Mux.Get("/dashboard/plugins/{id}/grants", p.grantsEdit)
	p.Mux.Post("/dashboard/plugins/{id}/grants", p.grantsUpdate)
	p.Mux.Get("/dashboard/plugins/{id}/routes", p.routesView)
}

// FuncMap returns a callable function that accepts a request.
func (p *Plugin) FuncMap(context.Context) func(r *http.Request) template.FuncMap {
	return func(r *http.Request) template.FuncMap {
		fm := make(template.FuncMap)
		fm["pluginmanager_URLHasParam"] = func(s string) bool {
			return strings.Contains(s, "{")
		}

		return fm
	}
}

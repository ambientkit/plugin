// Package redirecttourl is an Ambient plugin with middleware that redirects to the correct site URL.
package redirecttourl

import (
	"context"
	"net/http"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin with middleware that redirects to the correct site URL.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "redirecttourl"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to redirect to the correct URL if the user request URL doesn't match."},
		{Grant: ambient.GrantSitePluginRead, Description: "Access to read the scheme and URL settings to redirect to."},
	}
}

const (
	// SiteScheme allows user to set scheme to redirect to.
	SiteScheme = "Site Scheme"
	// SiteURL allows user to set the URL to redirect to.
	SiteURL = "Site URL"
)

// Settings returns a list of user settable fields.
func (p *Plugin) Settings() []ambient.Setting {
	return []ambient.Setting{
		{
			Name: SiteScheme,
			Description: ambient.SettingDescription{
				Text: "http or https",
			},
		},
		{
			Name: SiteURL,
			Description: ambient.SettingDescription{
				Text: "example: domain.com",
			},
		},
	}
}

// Middleware returns router middleware.
func (p *Plugin) Middleware() []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.stripSlash,
	}
}

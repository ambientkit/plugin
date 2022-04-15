// Package etagcache is an Ambient plugin that provides caching using etag.
package etagcache

import (
	"context"
	"net/http"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that provides gzip content compression midddleware.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "etagcache"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to read and write ETag headers on responses."},
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read MaxAge setting."},
	}
}

const (
	// MaxAge allows user to set MaxAge in seconds.
	MaxAge = "MaxAge"
)

// Settings returns a list of user settable fields.
func (p *Plugin) Settings(context.Context) []ambient.Setting {
	return []ambient.Setting{
		{
			Name: MaxAge,
			Description: ambient.SettingDescription{
				Text: "MaxAge in seconds before Etag is checked. 30 days is 2592000.",
			},
		},
	}
}

// Middleware returns router middleware.
func (p *Plugin) Middleware(context.Context) []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.Handler,
	}
}

// Package jwt is an Ambient plugin that enables jwt.
package jwt

import (
	"context"
	"net/http"
	"time"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/jwtoken"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase

	secret         []byte
	sessionTimeout time.Duration
	whitelist      []string
}

// New an Ambient plugin that provides request logging middleware.
func New(secret []byte, sessionTimeout time.Duration, whitelist []string) *Plugin {
	return &Plugin{
		PluginBase:     &ambient.PluginBase{},
		secret:         secret,
		sessionTimeout: sessionTimeout,
		whitelist:      whitelist,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "jwt"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to force authentication on routes using JWTs."},
	}
}

// Middleware returns router middleware.
func (p *Plugin) Middleware(context.Context) []func(next http.Handler) http.Handler {
	jwt := NewJWT(p.whitelist, jwtoken.New(p.secret, p.sessionTimeout), p.Toolkit.Site)
	return []func(next http.Handler) http.Handler{
		jwt.Handler,
	}
}

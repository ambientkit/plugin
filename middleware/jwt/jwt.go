// Package jwt is an Ambient plugin that enables jwt.
package jwt

import (
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
func (p *Plugin) PluginName() string {
	return "jwt"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Middleware returns router middleware.
func (p *Plugin) Middleware() []func(next http.Handler) http.Handler {
	jwt := NewJWT(p.whitelist, jwtoken.New(p.secret, p.sessionTimeout), p.Toolkit.Site)
	return []func(next http.Handler) http.Handler{
		jwt.Handler,
	}
}

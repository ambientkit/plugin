// Package notrailingslash is an Ambient plugin with middleware that removes trailing slashes from requests.
package notrailingslash

import (
	"context"
	"net/http"
	"strings"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new notrailingslash plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName an Ambient plugin with middleware that removes trailing slashes from requests.
func (p *Plugin) PluginName(context.Context) string {
	return "notrailingslash"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to redirect or respond with a 404 based on request URL."},
	}
}

// Middleware returns router middleware.
func (p *Plugin) Middleware(context.Context) []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.stripSlash,
	}
}

// StripSlash will strip trailing slashes from requests.
func (p *Plugin) stripSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Don't allow access to files with a slash at the end.
		if strings.Contains(r.URL.Path, ".") && strings.HasSuffix(r.URL.Path, "/") {
			p.Mux.Error(http.StatusNotFound, w, r)
			return
		}

		// Allow access to debug/pprof and force a trailing slash.
		if strings.HasPrefix(r.URL.Path, "/debug") {
			if r.URL.Path == p.Path("/debug/pprof") {
				http.Redirect(w, r, r.URL.Path+"/", http.StatusPermanentRedirect)
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		// Strip trailing slash.
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, strings.TrimRight(r.URL.Path, "/"), http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

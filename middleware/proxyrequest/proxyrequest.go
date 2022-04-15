// Package proxyrequest is an Ambient plugin with middleware that proxies requests.
package proxyrequest

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase

	urlForProxy  *url.URL
	prefixForAPI string

	handlerUI http.Handler
}

// New returns an Ambient plugin with middleware that proxies requests.
func New(urlForProxy *url.URL, prefixForAPI string) *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},

		urlForProxy:  urlForProxy,
		prefixForAPI: prefixForAPI,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "proxyrequest"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to proxy requests based on request URL."},
	}
}

// Enable accepts the toolkit.
func (p *Plugin) Enable(ctx context.Context, toolkit *ambient.Toolkit) error {
	err := p.PluginBase.Enable(ctx, toolkit)
	if err != nil {
		return err
	}

	uiProxy := httputil.NewSingleHostReverseProxy(p.urlForProxy)
	uiProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusBadGateway)
	}
	p.handlerUI = uiProxy

	return nil
}

// Middleware returns router middleware.
func (p *Plugin) Middleware(context.Context) []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.ProxyRequest,
	}
}

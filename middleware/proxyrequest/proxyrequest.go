// Package proxyrequest is an Ambient plugin with middleware that proxies requests.
package proxyrequest

import (
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
func (p *Plugin) PluginName() string {
	return "proxyrequest"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to proxy requests based on request URL."},
	}
}

// Enable accepts the toolkit.
func (p *Plugin) Enable(toolkit *ambient.Toolkit) error {
	err := p.PluginBase.Enable(toolkit)
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
func (p *Plugin) Middleware() []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.ProxyRequest,
	}
}

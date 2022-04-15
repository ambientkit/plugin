// Package htmlengine is an Ambient plugin that provides a HTML template engine.
package htmlengine

import (
	"context"
	"embed"

	"github.com/ambientkit/ambient"
)

//go:embed layout/*.tmpl
var assets embed.FS

// Plugin represents an Ambient plugin.
type Plugin struct{}

// New returns an Ambient plugin that provides a HTML template engine.
func New() *Plugin {
	return &Plugin{}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "htmlengine"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// TemplateEngine returns a template engine.
func (p *Plugin) TemplateEngine(_ context.Context, logger ambient.Logger, injector ambient.AssetInjector) (ambient.Renderer, error) {
	tmpl := NewTemplateEngine(logger, injector)
	return tmpl, nil
}

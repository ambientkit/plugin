// Package zaplogger is an Ambient plugin that provides logging using zap.
package zaplogger

import (
	"context"
	"io"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	log *Logger
}

// New returns an Ambient plugin that provides logging using zap.
func New() *Plugin {
	return &Plugin{}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "zaplogger"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// Logger returns a logger.
func (p *Plugin) Logger(_ context.Context, appName string, appVersion string, optionalWriter io.Writer) (ambient.AppLogger, error) {
	// Create the logger.
	p.log = NewLogger(appName, appVersion, optionalWriter)

	return p.log, nil
}

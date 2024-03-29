// Package logruslogger is an Ambient plugin that provides log functionality using logrus.
package logruslogger

import (
	"io"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	log *Logger
}

// New returns an Ambient plugin that provides log functionality using logrus.
func New() *Plugin {
	return &Plugin{}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "logruslogger"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Logger returns a logger.
func (p *Plugin) Logger(appName string, appVersion string, optionalWriter io.Writer) (ambient.AppLogger, error) {
	// Create the logger.
	p.log = NewLogger(appName, appVersion, optionalWriter)

	return p.log, nil
}

// Package localstorage is an Ambient plugin that provides local storage.
package localstorage

import (
	"context"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/filestore"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	sitePath    string
	sessionPath string
}

// New returns an Ambient plugin that provides local storage.
func New(sitePath string, sessionPath string) *Plugin {
	return &Plugin{
		sitePath:    sitePath,
		sessionPath: sessionPath,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "localstorage"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Storage returns data and session storage.
func (p *Plugin) Storage(logger ambient.Logger) (ambient.DataStorer, ambient.SessionStorer, error) {
	// Use local filesytem for site and session information.
	ds := filestore.New(p.sitePath)
	ss := filestore.New(p.sessionPath)

	return ds, ss, nil
}

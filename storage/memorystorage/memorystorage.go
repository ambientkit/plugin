// Package memorystorage is an Ambient plugin that provides storage in memory.
package memorystorage

import (
	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/memorystore"
)

// Plugin represents an Ambient plugin.
type Plugin struct{}

// New returns an Ambient plugin that provides local storage.
func New() *Plugin {
	return &Plugin{}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "memorystorage"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Storage returns data and session storage.
func (p *Plugin) Storage(logger ambient.Logger) (ambient.DataStorer, ambient.SessionStorer, error) {
	// Use local filesytem for site and session information.
	ds := memorystore.New()
	ss := memorystore.New()

	return ds, ss, nil
}

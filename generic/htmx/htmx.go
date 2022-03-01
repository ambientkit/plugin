// Package htmx is an Ambient plugin that adds the htmx JavaScript library to all pages: https://htmx.org/.
package htmx

import (
	"embed"
	"fmt"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new htmx plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "htmx"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read the version setting."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to add the htmx JavaScript tag to every page."},
	}
}

const (
	// Version allows user to set the library version.
	Version = "Version"
)

// Settings returns a list of plugin settings.
func (p *Plugin) Settings() []ambient.Setting {
	return []ambient.Setting{
		{
			Name:    Version,
			Default: "1.7.0",
		},
	}
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets() ([]ambient.Asset, *embed.FS) {
	version, err := p.Site.PluginSettingString(Version)
	if err != nil || len(version) == 0 {
		// Otherwise don't set the assets.
		return nil, nil
	}

	return []ambient.Asset{
		{
			Filetype: ambient.AssetJavaScript,
			Location: ambient.LocationHead,
			External: true,
			Path:     fmt.Sprintf("https://unpkg.com/htmx.org@%v", version),
		},
	}, nil
}

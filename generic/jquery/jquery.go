// Package jquery is an Ambient plugin that adds the jQuery library to all pages: https://jquery.com/.
package jquery

import (
	"embed"
	"fmt"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new jquery plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "jquery"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read the version setting."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to add the jquery JavaScript tag to every page."},
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
			Default: "3.6.0",
			Description: ambient.SettingDescription{
				Text: "Version cannot be left blank. Ex: 3.6.0",
				URL:  "https://releases.jquery.com/jquery/",
			},
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
			Location: ambient.LocationBody,
			External: true,
			Path:     fmt.Sprintf("https://code.jquery.com/jquery-%v.min.js", version),
			Attributes: []ambient.Attribute{
				{
					Name:  "crossorigin",
					Value: "anonymous",
				},
			},
		},
	}, nil
}

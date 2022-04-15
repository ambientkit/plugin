// Package tailwindcss is an Ambient plugin that adds the Tailwind CSS library to all pages: https://tailwindcsscss.com/.
package tailwindcss

import (
	"context"
	"fmt"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new tailwindcss plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "tailwindcss"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read the version setting."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to add the Tailwind CSS JavaScript tag to every page."},
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
			Name: Version,
			Description: ambient.SettingDescription{
				Text: "When blank, will use the latest version. Ex: 3.0.23",
				URL:  "https://github.com/tailwindlabs/tailwindcss/releases",
			},
		},
	}
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets() ([]ambient.Asset, ambient.FileSystemReader) {
	version, err := p.Site.PluginSettingString(Version)
	if err != nil {
		// Otherwise don't set the assets.
		return nil, nil
	}

	return []ambient.Asset{
		{
			Filetype: ambient.AssetJavaScript,
			Location: ambient.LocationHead,
			External: true,
			Path:     fmt.Sprintf("https://cdn.tailwindcss.com/%v", version),
		},
	}, nil
}

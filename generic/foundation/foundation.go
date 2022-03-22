// Package foundation is an Ambient plugin that adds the Foundation library to all pages: https://get.foundation/. It requires jQuery.
package foundation

import (
	"fmt"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new foundation plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "foundation"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read the version setting."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to add the Foundation CSS and JavaScript tags to every page."},
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
			Default: "6.7.4",
			Description: ambient.SettingDescription{
				Text: "Version cannot be left blank. Ex: 6.7.4",
				URL:  "https://github.com/foundation/foundation-sites/releases",
			},
		},
	}
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets() ([]ambient.Asset, ambient.FileSystemReader) {
	version, err := p.Site.PluginSettingString(Version)
	if err != nil || len(version) == 0 {
		// Otherwise don't set the assets.
		return nil, nil
	}

	return []ambient.Asset{
		{
			Filetype: ambient.AssetStylesheet,
			Location: ambient.LocationHead,
			External: true,
			Path:     fmt.Sprintf("https://cdn.jsdelivr.net/npm/foundation-sites@%v/dist/css/foundation.min.css", version),
			Attributes: []ambient.Attribute{
				{
					Name:  "crossorigin",
					Value: "anonymous",
				},
			},
		},
		{
			Filetype: ambient.AssetJavaScript,
			Location: ambient.LocationBody,
			External: true,
			Path:     fmt.Sprintf("https://cdn.jsdelivr.net/npm/foundation-sites@%v/dist/js/foundation.min.js", version),
			Attributes: []ambient.Attribute{
				{
					Name:  "crossorigin",
					Value: "anonymous",
				},
			},
		},
		{
			Filetype: ambient.AssetJavaScript,
			Location: ambient.LocationBody,
			Inline:   true,
			Content:  "$(document).foundation();",
		},
	}, nil
}

// Package bootstrap is an Ambient plugin that adds the Bootstrap library to all pages: https://getbootstrap.com/.
package bootstrap

import (
	"context"
	"fmt"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new bootstrap plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "bootstrap"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read the version setting."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to add the Bootstrap CSS and JavaScript tags to every page."},
	}
}

const (
	// Version allows user to set the library version.
	Version = "Version"
)

// Settings returns a list of plugin settings.
func (p *Plugin) Settings(context.Context) []ambient.Setting {
	return []ambient.Setting{
		{
			Name:    Version,
			Default: "5.1.3",
			Description: ambient.SettingDescription{
				Text: "Version cannot be left blank. Ex: 5.1.3",
				URL:  "https://github.com/twbs/bootstrap/releases",
			},
		},
	}
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets(ctx context.Context) ([]ambient.Asset, ambient.FileSystemReader) {
	version, err := p.Site.PluginSettingString(ctx, Version)
	if err != nil || len(version) == 0 {
		// Otherwise don't set the assets.
		return nil, nil
	}

	return []ambient.Asset{
		{
			Filetype: ambient.AssetStylesheet,
			Location: ambient.LocationHead,
			External: true,
			Path:     fmt.Sprintf("https://cdn.jsdelivr.net/npm/bootstrap@%v/dist/css/bootstrap.min.css", version),
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
			Path:     fmt.Sprintf("https://cdn.jsdelivr.net/npm/bootstrap@%v/dist/js/bootstrap.bundle.min.js", version),
			Attributes: []ambient.Attribute{
				{
					Name:  "crossorigin",
					Value: "anonymous",
				},
			},
		},
	}, nil
}

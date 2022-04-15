// Package viewport is an Ambient plugin that sets a viewport meta tag in the HTML header.
package viewport

import (
	"context"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns an Ambient plugin that sets a viewport meta tag in the HTML header.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "viewport"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

const (
	// Viewport allows user to set the viewport.
	Viewport = "Viewport"
)

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantPluginSettingRead, Description: "Access to read the plugin settings."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to write a meta tag to the header."},
	}
}

// Settings returns a list of user settable fields.
func (p *Plugin) Settings(context.Context) []ambient.Setting {
	return []ambient.Setting{
		{
			Name:    Viewport,
			Default: "width=device-width, initial-scale=1.0",
		},
	}
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets(ctx context.Context) ([]ambient.Asset, ambient.FileSystemReader) {
	vp, err := p.Site.PluginSettingString(ctx, Viewport)
	if err != nil || len(vp) == 0 {
		// Otherwise don't set the assets.
		return nil, nil
	}

	return []ambient.Asset{
		{
			Filetype:   ambient.AssetGeneric,
			Location:   ambient.LocationHead,
			TagName:    "meta",
			ClosingTag: false,
			Attributes: []ambient.Attribute{
				{
					Name:  "name",
					Value: "viewport",
				},
				{
					Name:  "content",
					Value: vp,
				},
			},
		},
	}, nil
}

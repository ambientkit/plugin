package pluginmanager

import (
	"net/http"
	"sort"

	"github.com/ambientkit/ambient"
)

type pluginWithSettings struct {
	PluginData ambient.PluginData     `json:"plugindata"`
	Name       string                 `json:"name"`
	Settings   []ambient.Setting      `json:"settings"`
	Grants     []ambient.GrantRequest `json:"grants"`
	Trusted    bool                   `json:"trusted"`
	Routes     []ambient.Route        `json:"routes"`
}

func (p *Plugin) edit(w http.ResponseWriter, r *http.Request) (err error) {
	vars := make(map[string]interface{})
	vars["title"] = "Plugin Manager"
	vars["token"] = p.Site.SetCSRF(r)

	plugins, err := p.Site.Plugins()
	if err != nil {
		return p.Site.Error(err)
	}

	pluginNames, err := p.Site.PluginNames()
	if err != nil {
		return p.Site.Error(err)
	}
	sort.Strings(pluginNames)

	arr := make([]pluginWithSettings, 0)
	for _, pluginName := range pluginNames {
		// Get the list of grants.
		grantList, err := p.Site.NeighborPluginGrantList(pluginName)
		if err != nil {
			return p.Site.Error(err)
		}

		// Get the list of settings.
		settingsList, err := p.Site.PluginNeighborSettingsList(pluginName)
		if err != nil {
			return p.Site.Error(err)
		}

		trusted, err := p.Site.PluginTrusted(pluginName)
		if err != nil {
			return p.Site.Error(err)
		}

		routes := make([]ambient.Route, 0)
		// Only enabled plugins have routes.
		if plugins[pluginName].Enabled {
			routes, err = p.Site.PluginNeighborRoutesList(pluginName)
			if err != nil {
				return p.Site.Error(err)
			}
		}

		arr = append(arr, pluginWithSettings{
			Name:       pluginName,
			PluginData: plugins[pluginName],
			Grants:     grantList,
			Settings:   settingsList,
			Trusted:    trusted,
			Routes:     routes,
		})
	}

	vars["plugins"] = arr

	return p.Render.Page(w, r, assets, "template/plugins_edit.tmpl", p.FuncMap(), vars)
}

func (p *Plugin) update(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()

	// CSRF protection.
	ok := p.Site.CSRF(r, r.FormValue("token"))
	if !ok {
		return p.Mux.StatusError(http.StatusBadRequest, nil)
	}

	// Get list of plugin names.
	names, err := p.Site.PluginNames()
	if err != nil {
		return p.Site.Error(err)
	}

	// Get list of plugins.
	plugins, err := p.Site.Plugins()
	if err != nil {
		return p.Site.Error(err)
	}

	// Disable plugins: loop through each plugin to get the settings then save.
	// Disable plugins first so they don't collide with enabling plugins that
	// have the same routes, etc.
	for _, name := range names {
		info, ok := plugins[name]
		if !ok {
			continue
		}

		trusted, err := p.Site.PluginTrusted(name)
		if err != nil {
			return p.Site.Error(err)
		}

		enable := (r.FormValue(name) == "on")
		// Only disable plugins that are enabled and not trusted since trusted
		// plugins can't be disabled.
		if !enable && info.Enabled && !trusted {
			// Disable the plugin.
			err = p.Site.DisablePlugin(name, true)
			if err != nil {
				return p.Site.Error(err)
			}
		}
	}

	// Enable plugins: loop through each plugin to get the settings then save.
	for _, name := range names {
		info, ok := plugins[name]
		if !ok {
			continue
		}

		enable := (r.FormValue(name) == "on")
		if enable && !info.Enabled {
			err = p.Site.EnablePlugin(name, true)
			if err != nil {
				return p.Site.Error(err)
			}
		}
	}

	p.Redirect(w, r, "/dashboard/plugins", http.StatusFound)
	return
}

func (p *Plugin) destroy(w http.ResponseWriter, r *http.Request) (err error) {
	ID := p.Mux.Param(r, "id")

	plugins, err := p.Site.Plugins()
	if err != nil {
		return p.Site.Error(err)
	}

	if _, ok := plugins[ID]; !ok {
		return p.Mux.StatusError(http.StatusNotFound, nil)
	}

	err = p.Site.DeletePlugin(r.Context(), ID)
	if err != nil {
		return p.Site.Error(err)
	}

	p.Redirect(w, r, "/dashboard/plugins", http.StatusFound)
	return
}

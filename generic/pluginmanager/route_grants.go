package pluginmanager

import (
	"fmt"
	"net/http"

	"github.com/ambientkit/ambient"
)

type pluginGrant struct {
	Index       int
	Name        ambient.Grant
	Granted     bool
	Description string
}

func (p *Plugin) grantsEdit(w http.ResponseWriter, r *http.Request) (err error) {
	pluginName := p.Mux.Param(r, "id")

	vars := make(map[string]interface{})
	vars["title"] = "Edit grants for: " + pluginName
	vars["token"] = p.Site.SetCSRF(r)

	grantList, err := p.Site.NeighborPluginGrantList(r.Context(), pluginName)
	if err != nil {
		return p.Site.Error(err)
	}

	grants, err := p.Site.NeighborPluginGrants(r.Context(), pluginName)
	if err != nil {
		return p.Site.Error(err)
	}

	trusted, err := p.Site.PluginTrusted(r.Context(), pluginName)
	if err != nil {
		return p.Site.Error(err)
	}

	arr := make([]pluginGrant, 0)
	for index, request := range grantList {
		arr = append(arr, pluginGrant{
			Index:       index,
			Name:        request.Grant,
			Granted:     grants[request.Grant],
			Description: request.Description,
		})
	}

	vars["trusted"] = trusted
	vars["grants"] = arr

	return p.Render.Page(w, r, assets, "template/grants_edit.tmpl", p.FuncMap(r.Context()), vars)
}

func (p *Plugin) grantsUpdate(w http.ResponseWriter, r *http.Request) (err error) {
	pluginName := p.Mux.Param(r, "id")
	r.ParseForm()

	// CSRF protection.
	ok := p.Site.CSRF(r, r.FormValue("token"))
	if !ok {
		return p.Mux.StatusError(http.StatusBadRequest, nil)
	}

	grantList, err := p.Site.NeighborPluginGrantList(r.Context(), pluginName)
	if err != nil {
		return p.Site.Error(err)
	}

	// Loop through each plugin to get the grants then save.
	for index, request := range grantList {
		val := r.FormValue(fmt.Sprintf("field%v", index))
		err := p.Site.SetNeighborPluginGrant(r.Context(), pluginName, request.Grant, val == "true")
		if err != nil {
			return p.Site.Error(err)
		}
	}

	p.Redirect(w, r, fmt.Sprintf("/dashboard/plugins/%v/grants", pluginName), http.StatusFound)
	return
}

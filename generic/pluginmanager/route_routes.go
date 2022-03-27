package pluginmanager

import (
	"net/http"
)

func (p *Plugin) routesView(w http.ResponseWriter, r *http.Request) (err error) {
	pluginName := p.Mux.Param(r, "id")

	vars := make(map[string]interface{})
	vars["title"] = "View routes for: " + pluginName
	vars["token"] = p.Site.SetCSRF(r)

	routes, err := p.Site.PluginNeighborRoutesList(pluginName)
	if err != nil {
		return p.Site.Error(err)
	}

	vars["routes"] = routes

	return p.Render.Page(w, r, assets, "template/routes_view.tmpl", p.FuncMap(), vars)
}

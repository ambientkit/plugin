package bearblog

import (
	"net/http"
)

func (p *Plugin) index(w http.ResponseWriter, r *http.Request) (err error) {
	content, err := p.Site.Content(r.Context())
	if err != nil {
		return p.Site.Error(err)
	}

	if content == "" {
		content = "*No content yet.*"
	}

	vars := make(map[string]interface{})
	vars["title"] = ""
	vars["tags"] = ""
	vars["postcontent"] = p.sanitized(r.Context(), content)
	return p.Render.Page(w, r, assets, "template/content/home.tmpl", p.FuncMap(r.Context()), vars)
}

func (p *Plugin) edit(w http.ResponseWriter, r *http.Request) (err error) {
	siteContent, err := p.Site.Content(r.Context())
	if err != nil {
		return p.Site.Error(err)
	}

	siteTitle, err := p.Site.Title(r.Context())
	if err != nil {
		return p.Site.Error(err)
	}

	siteSubtitle, err := p.Site.PluginSetting(r.Context(), Subtitle)
	if err != nil {
		return p.Site.Error(err)
	}

	baseURL, err := p.Site.URL(r.Context())
	if err != nil {
		return p.Site.Error(err)
	}

	siteScheme, err := p.Site.Scheme(r.Context())
	if err != nil {
		return p.Site.Error(err)
	}

	loginURL, err := p.Site.PluginSetting(r.Context(), LoginURL)
	if err != nil {
		return p.Site.Error(err)
	}

	footer, err := p.Site.PluginSetting(r.Context(), Footer)
	if err != nil {
		return p.Site.Error(err)
	}

	vars := make(map[string]interface{})
	vars["title"] = "Edit site"
	vars["homeContent"] = siteContent
	vars["ptitle"] = siteTitle
	vars["subtitle"] = siteSubtitle
	vars["token"] = p.Site.SetCSRF(r)

	// Help the user set the domain based off the current URL.
	if baseURL == "" {
		vars["domain"] = r.Host
	} else {
		vars["domain"] = baseURL
	}

	vars["scheme"] = siteScheme
	vars["loginurl"] = loginURL
	vars["footer"] = footer

	return p.Render.Page(w, r, assets, "template/content/home_edit.tmpl", p.FuncMap(r.Context()), vars)
}

func (p *Plugin) update(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()

	// CSRF protection.
	success := p.Site.CSRF(r, r.FormValue("token"))
	if !success {
		return p.Mux.StatusError(http.StatusBadRequest, nil)
	}

	err = p.Site.SetTitle(r.Context(), r.FormValue("title"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetContent(r.Context(), r.FormValue("content"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetScheme(r.Context(), r.FormValue("scheme"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetURL(r.Context(), r.FormValue("domain"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetPluginSetting(r.Context(), Subtitle, r.FormValue("subtitle"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetPluginSetting(r.Context(), LoginURL, r.FormValue("loginurl"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetPluginSetting(r.Context(), Footer, r.FormValue("footer"))
	if err != nil {
		return p.Site.Error(err)
	}

	p.Redirect(w, r, "/dashboard", http.StatusFound)
	return
}

func (p *Plugin) reload(w http.ResponseWriter, r *http.Request) (err error) {
	err = p.Site.Load(r.Context())
	if err != nil {
		p.Site.Error(err)
	}

	p.Redirect(w, r, "/dashboard", http.StatusFound)
	return
}

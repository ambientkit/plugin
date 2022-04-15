package simplelogin

import (
	"context"
	"net/http"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func (p *Plugin) index(w http.ResponseWriter, r *http.Request) (err error) {
	content, err := p.Site.Content()
	if err != nil {
		return p.Site.Error(err)
	}

	if content == "" {
		content = "*No content yet.*"
	}

	vars := make(map[string]interface{})
	vars["postcontent"] = p.sanitized(r.Context(), content)
	return p.Render.Page(w, r, assets, "template/content/home.tmpl", p.FuncMap(r.Context()), vars)
}

func (p *Plugin) edit(w http.ResponseWriter, r *http.Request) (err error) {
	siteContent, err := p.Site.Content()
	if err != nil {
		return p.Site.Error(err)
	}

	siteTitle, err := p.Site.Title()
	if err != nil {
		return p.Site.Error(err)
	}

	siteSubtitle, err := p.Site.PluginSetting(r.Context(), Subtitle)
	if err != nil {
		return p.Site.Error(err)
	}

	baseURL, err := p.Site.URL()
	if err != nil {
		return p.Site.Error(err)
	}

	siteScheme, err := p.Site.Scheme()
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

	err = p.Site.SetTitle(r.FormValue("title"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetContent(r.FormValue("content"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetScheme(r.FormValue("scheme"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetURL(r.FormValue("domain"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetPluginSetting(Subtitle, r.FormValue("subtitle"))
	if err != nil {
		return p.Site.Error(err)
	}

	err = p.Site.SetPluginSetting(Footer, r.FormValue("footer"))
	if err != nil {
		return p.Site.Error(err)
	}

	p.Redirect(w, r, "/dashboard", http.StatusFound)
	return
}

func (p *Plugin) reload(w http.ResponseWriter, r *http.Request) (err error) {
	err = p.Site.Load()
	if err != nil {
		p.Site.Error(err)
	}

	p.Redirect(w, r, "/dashboard", http.StatusFound)
	return
}

// sanitized returns a sanitized content block or an error is one occurs.
func (p *Plugin) sanitized(ctx context.Context, content string) string {
	b := []byte(content)
	// Ensure unit line endings are used when pulling out of JSON.
	markdownWithUnixLineEndings := strings.Replace(string(b), "\r\n", "\n", -1)
	htmlCode := blackfriday.Run([]byte(markdownWithUnixLineEndings))

	// Determine if raw HTML is allowed.
	allowed, err := p.Site.PluginSettingBool(ctx, AllowHTMLinMarkdown)
	if err != nil {
		p.Log.Debug("plugins: error in sanitized() getting plugin field: %v", err)
	}

	// Sanitize by removing HTML if allowed.
	if !allowed {
		htmlCode = bluemonday.UGCPolicy().SanitizeBytes(htmlCode)
	}

	return string(htmlCode)
}

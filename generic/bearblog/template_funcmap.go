package bearblog

import (
	"context"
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/ambientkit/ambient"
)

// FuncMap returns a callable function that accepts a request.
func (p *Plugin) FuncMap(context.Context) func(r *http.Request) template.FuncMap {
	return func(r *http.Request) template.FuncMap {
		fm := make(template.FuncMap)
		fm["bearblog_Stamp"] = func(t string) string {
			tt, err := time.Parse(time.RFC3339, t)
			if err != nil {
				return t
			}
			return tt.Format("2006-01-02")
		}
		fm["bearblog_StampFriendly"] = func(t string) string {
			tt, err := time.Parse(time.RFC3339, t)
			if err != nil {
				return t
			}
			return tt.Format("02 Jan, 2006")
		}
		fm["bearblog_PublishedPages"] = func() []ambient.Post {
			arr, err := p.Site.PublishedPages(r.Context())
			if err != nil {
				p.Log.Warn("bearblog: error getting published pages: %v", err.Error())
			}
			return arr
		}
		fm["bearblog_SiteSubtitle"] = func() string {
			subtitle, err := p.Site.PluginSettingString(r.Context(), Subtitle)
			if err != nil {
				p.Log.Warn("bearblog: error getting subtitle: %v", err.Error())
			}
			return subtitle
		}
		fm["bearblog_Authenticated"] = func() bool {
			// If user is not authenticated, don't allow them to access the page.
			_, err := p.Site.AuthenticatedUser(r)
			return err == nil
		}
		fm["bearblog_SiteFooter"] = func() string {
			f, err := p.Site.PluginSettingString(r.Context(), Footer)
			if err != nil {
				p.Log.Warn("bearblog: error getting footer: %v", err.Error())
			}
			return f
		}
		fm["bearblog_PageURL"] = func() string {
			siteURL, err := p.Site.FullURL(r.Context())
			if err != nil {
				p.Log.Warn("bearblog: error getting site URL: %v", err.Error())
			}

			return path.Join(siteURL, r.URL.Path)
		}
		fm["bearblog_MFAEnabled"] = func() bool {
			mfakey, err := p.Site.PluginSettingString(r.Context(), MFAKey)
			if err != nil {
				p.Log.Warn("bearblog: error getting MFA key: %v", err.Error())
			}
			return len(mfakey) > 0
		}

		return fm
	}
}

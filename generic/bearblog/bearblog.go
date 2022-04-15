// Package bearblog is an Ambient plugin that provides basic blog functionality.
package bearblog

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/ambientkit/ambient"
)

//go:embed template/partial/*.tmpl template/content/*.tmpl
var assets embed.FS

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase

	passwordHash string
}

// New returns an Ambient plugin that provides basic blog functionality.
func New(passwordHash string) *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},

		passwordHash: passwordHash,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "bearblog"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests(context.Context) []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantUserAuthenticatedRead, Description: "Show different menus to authenticated vs unauthenticated users."},
		{Grant: ambient.GrantUserAuthenticatedWrite, Description: "Access to login and logout the user."},
		{Grant: ambient.GrantUserPersistWrite, Description: "Access to set session as persistent."},
		{Grant: ambient.GrantPluginSettingRead, Description: "Read own plugin settings."},
		{Grant: ambient.GrantPluginSettingWrite, Description: "Write own plugin settings."},
		{Grant: ambient.GrantSitePostRead, Description: "Read all site posts."},
		{Grant: ambient.GrantSitePostWrite, Description: "Create and edit site posts."},
		{Grant: ambient.GrantSiteSchemeRead, Description: "Read site scheme."},
		{Grant: ambient.GrantSiteSchemeWrite, Description: "Update the site scheme."},
		{Grant: ambient.GrantSiteURLRead, Description: "Read the site URL."},
		{Grant: ambient.GrantSiteURLWrite, Description: "Update the site URL."},
		{Grant: ambient.GrantSiteTitleRead, Description: "Read the site title."},
		{Grant: ambient.GrantSiteTitleWrite, Description: "Update the site title."},
		{Grant: ambient.GrantSiteContentRead, Description: "Read home page content."},
		{Grant: ambient.GrantSiteContentWrite, Description: "Update home page content."},
		{Grant: ambient.GrantSiteAssetWrite, Description: "Access to write blog meta tags to the header and add a nav and footer."},
		{Grant: ambient.GrantSiteFuncMapWrite, Description: "Access to create global FuncMaps for templates."},
		{Grant: ambient.GrantRouterRouteWrite, Description: "Access to create routes for editing the blog posts."},
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to create global middleware to protect /dashboard/* routes from anonymous users."},
	}
}

const (
	// LoginURL allows user to set the login URL.
	LoginURL = "Login URL"
	// Author allows user to set the author.
	Author = "Author"
	// Subtitle allows user to set the Subtitle.
	Subtitle = "Subtitle"
	// Description allows user to set the description.
	Description = "Description"
	// Footer allows user to set the footer.
	Footer = "Footer"
	// AllowHTMLinMarkdown allows user to set if they allow HTML in markdown.
	AllowHTMLinMarkdown = "Allow HTML in Markdown"

	// Username allows user to set the login username.
	Username = "Username"
	// Password allows user to set the login password.
	Password = "Password"
	// MFAKey allows user to set the MFA key.
	MFAKey = "MFA Key"
)

// Settings returns a list of user settable fields.
func (p *Plugin) Settings(context.Context) []ambient.Setting {
	return []ambient.Setting{
		{
			Name:    Username,
			Default: "admin",
		},
		{
			Name:    Password,
			Default: p.passwordHash,
			Type:    ambient.InputPassword,
			Hide:    true,
		},
		{
			Name: MFAKey,
			Type: ambient.InputPassword,
			Description: ambient.SettingDescription{
				Text: "Generate an MFA key. Plugin must be enabled first.",
				URL:  "/dashboard/mfa",
			},
		},
		{
			Name:    LoginURL,
			Default: "admin",
			Hide:    true,
		},
		{
			Name: Author,
		},
		{
			Name: Subtitle,
			Hide: true,
		},
		{
			Name: Description,
			Type: ambient.Textarea,
		},
		{
			Name: Footer,
			Type: ambient.Textarea,
			Hide: true,
		},
		{
			Name: AllowHTMLinMarkdown,
			Type: ambient.Checkbox,
		},
	}
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes(context.Context) {
	p.Mux.Get("/blog", p.postIndex)
	p.Mux.Get("/{slug}", p.postShow)

	p.Mux.Get("/login/{slug}", p.login)
	p.Mux.Post("/login/{slug}", p.loginPost)
	p.Mux.Get("/dashboard/logout", p.logout)

	p.Mux.Get("/", p.index)
	p.Mux.Get("/dashboard", p.edit)
	p.Mux.Post("/dashboard", p.update)
	p.Mux.Get("/dashboard/reload", p.reload)

	p.Mux.Get("/dashboard/mfa", p.mfa)
	p.Mux.Post("/dashboard/mfa", p.mfaPost)

	p.Mux.Get("/dashboard/posts", p.postAdminIndex)
	p.Mux.Get("/dashboard/posts/new", p.postAdminCreate)
	p.Mux.Post("/dashboard/posts/new", p.postAdminStore)
	p.Mux.Get("/dashboard/posts/{id}", p.postAdminEdit)
	p.Mux.Post("/dashboard/posts/{id}", p.postAdminUpdate)
	p.Mux.Get("/dashboard/posts/{id}/delete", p.postAdminDestroy)
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets(ctx context.Context) ([]ambient.Asset, ambient.FileSystemReader) {
	arr := make([]ambient.Asset, 0)

	siteTitle, err := p.Site.Title(ctx)
	if err == nil && len(siteTitle) > 0 {
		arr = append(arr, ambient.Asset{
			Filetype: ambient.AssetGeneric,
			Location: ambient.LocationHead,
			TagName:  "title",
			Inline:   true,
			Content:  fmt.Sprintf(`{{if .pagetitle}}{{.pagetitle}} | %v{{else}}%v{{end}}`, siteTitle, siteTitle),
		})
	}

	siteDescription, err := p.Site.PluginSettingString(ctx, Description)
	if err == nil && len(siteDescription) > 0 {
		arr = append(arr, ambient.Asset{
			Filetype:   ambient.AssetGeneric,
			Location:   ambient.LocationHead,
			TagName:    "meta",
			ClosingTag: false,
			Attributes: []ambient.Attribute{
				{
					Name:  "name",
					Value: "description",
				},
				{
					Name:  "content",
					Value: fmt.Sprintf("{{if .pagedescription}}{{.pagedescription}}{{else}}%v{{end}}", siteDescription),
				},
			},
		})
	}

	arr = append(arr, ambient.Asset{
		Filetype:   ambient.AssetGeneric,
		Location:   ambient.LocationHead,
		TagName:    "link",
		ClosingTag: false,
		Attributes: []ambient.Attribute{
			{
				Name:  "rel",
				Value: "canonical",
			},
			{
				Name:  "href",
				Value: `{{if .canonical}}{{.canonical}}{{else}}{{bearblog_PageURL}}{{end}}`,
			},
		},
	})

	siteAuthor, err := p.Site.PluginSettingString(ctx, Author)
	if err == nil && len(siteAuthor) > 0 {
		arr = append(arr, ambient.Asset{
			Filetype:   ambient.AssetGeneric,
			Location:   ambient.LocationHead,
			TagName:    "meta",
			ClosingTag: false,
			Attributes: []ambient.Attribute{
				{
					Name:  "name",
					Value: "author",
				},
				{
					Name:  "content",
					Value: siteAuthor,
				},
			},
		})
	}

	arr = append(arr, ambient.Asset{
		Path:     "template/partial/nav.tmpl",
		Filetype: ambient.AssetGeneric,
		Location: ambient.LocationHeader,
		Inline:   true,
	})

	arr = append(arr, ambient.Asset{
		Path:     "template/partial/footer.tmpl",
		Filetype: ambient.AssetGeneric,
		Location: ambient.LocationFooter,
		Inline:   true,
	})

	return arr, &assets
}

// Middleware returns router middleware.
func (p *Plugin) Middleware(context.Context) []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.DisallowAnon,
	}
}

// DisallowAnon does not allow anonymous users to access the page.
func (p *Plugin) DisallowAnon(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Don't allow anon users to access the dashboard.
		if strings.HasPrefix(r.URL.Path, p.Path("/dashboard")) {
			// If user is not authenticated, don't allow them to access the page.
			username, err := p.Site.AuthenticatedUser(r)
			if err != nil || len(username) == 0 {
				p.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}

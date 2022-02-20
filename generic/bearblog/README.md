# bearblog

Package bearblog is an Ambient plugin that provides basic blog functionality.

**Import:** github.com/ambientkit/plugin/generic/bearblog

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the follow core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (18):

- **Name**: user.authenticated:read
  - **Description**: Show different menus to authenticated vs unauthenticated users.
- **Name**: user.authenticated:write
  - **Description**: Access to login and logout the user.
- **Name**: user.persist:write
  - **Description**: Access to set session as persistent.
- **Name**: plugin.setting:read
  - **Description**: Read own plugin settings.
- **Name**: plugin.setting:write
  - **Description**: Write own plugin settings.
- **Name**: site.post:read
  - **Description**: Read all site posts.
- **Name**: site.post:write
  - **Description**: Create and edit site posts.
- **Name**: site.scheme:read
  - **Description**: Read site scheme.
- **Name**: site.scheme:write
  - **Description**: Update the site scheme.
- **Name**: site.url:read
  - **Description**: Read the site URL.
- **Name**: site.url:write
  - **Description**: Update the site URL.
- **Name**: site.title:read
  - **Description**: Read the site title.
- **Name**: site.title:write
  - **Description**: Update the site title.
- **Name**: site.content:read
  - **Description**: Read home page content.
- **Name**: site.content:write
  - **Description**: Update home page content.
- **Name**: site.asset:write
  - **Description**: Access to write blog meta tags to the header and add a nav and footer.
- **Name**: site.funcmap:write
  - **Description**: Access to create global FuncMaps for templates.
- **Name**: router.route:write
  - **Description**: Access to create routes for editing the blog posts.

## Settings

The plugin has the follow settings (9):

- **Name**: Username
  - **Type**: input
  - **Hidden**: false
  - **Default**: admin
- **Name**: Password
  - **Type**: password
  - **Hidden**: true
  - **Default**: JDJhJDEwJE1tS0tMUDlzb1YudW5VZ0RkYzFzZ2VMOEJseWpzQ05kUjVUMFMwRE0udlhPaGp5aHRIQWdT
- **Name**: MFA Key
  - **Type**: password
  - **Description**: Generate an MFA key. Plugin must be enabled first.
    - **URL**: /dashboard/mfa
  - **Hidden**: false
- **Name**: Login URL
  - **Type**: input
  - **Hidden**: true
  - **Default**: admin
- **Name**: Author
  - **Type**: input
  - **Hidden**: false
- **Name**: Subtitle
  - **Type**: input
  - **Hidden**: true
- **Name**: Description
  - **Type**: textarea
  - **Hidden**: false
- **Name**: Footer
  - **Type**: textarea
  - **Hidden**: true
- **Name**: Allow HTML in Markdown
  - **Type**: checkbox
  - **Hidden**: false

## Routes

The plugin has the following routes (17):
  - **Method:** GET | **Path:** /blog
  - **Method:** GET | **Path:** /{slug}
  - **Method:** GET | **Path:** /login/{slug}
  - **Method:** POST | **Path:** /login/{slug}
  - **Method:** GET | **Path:** /dashboard/logout
  - **Method:** GET | **Path:** /
  - **Method:** GET | **Path:** /dashboard
  - **Method:** POST | **Path:** /dashboard
  - **Method:** GET | **Path:** /dashboard/reload
  - **Method:** GET | **Path:** /dashboard/mfa
  - **Method:** POST | **Path:** /dashboard/mfa
  - **Method:** GET | **Path:** /dashboard/posts
  - **Method:** GET | **Path:** /dashboard/posts/new
  - **Method:** POST | **Path:** /dashboard/posts/new
  - **Method:** GET | **Path:** /dashboard/posts/{id}
  - **Method:** POST | **Path:** /dashboard/posts/{id}
  - **Method:** GET | **Path:** /dashboard/posts/{id}/delete

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin has the follow FuncMap items (8):

  - {{bearblog_SiteSubtitle}}
  - {{bearblog_Authenticated}}
  - {{bearblog_SiteFooter}}
  - {{bearblog_PageURL}}
  - {{bearblog_MFAEnabled}}
  - {{bearblog_Stamp}}
  - {{bearblog_StampFriendly}}
  - {{bearblog_PublishedPages}}

## Assets

The plugin injects the following assets (3):

  - **Type:** generic
    - **Location:** head
    - **Tag Name:** link
    - **Attributes (2):** 
      - **Name:** rel | **Value:** canonical
      - **Name:** href | **Value:** {{if .canonical}}{{.canonical}}{{else}}{{bearblog_PageURL}}{{end}}
  - **Type:** generic
    - **Location:** header
    - **Inline:** true
    - **Path:** template/partial/nav.tmpl
  - **Type:** generic
    - **Location:** footer
    - **Inline:** true
    - **Path:** template/partial/footer.tmpl

## Embedded Files

The plugin has embedded files.

## Example Usage

```go
package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/generic/bearblog"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/passhash"
	"github.com/ambientkit/plugin/storage/memorystorage"
)

func main() {
	// Generate a password hash.
	s, err := passhash.HashString(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}

	plugins := &ambient.PluginLoader{
		// Core plugins are implicitly trusted.
		Router:         nil,
		TemplateEngine: nil,
		SessionManager: nil,
		// Trusted plugins are those that are typically needed to boot so they
		// will be enabled and given full access.
		TrustedPlugins: map[string]bool{},
		Plugins: []ambient.Plugin{
			bearblog.New(base64.StdEncoding.EncodeToString([]byte(s))),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
		},
	}
	_, _, err = ambient.NewApp("myapp", "1.0",
		zaplogger.New(),
		ambient.StoragePluginGroup{
			Storage: memorystorage.New(),
		},
		plugins)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
```
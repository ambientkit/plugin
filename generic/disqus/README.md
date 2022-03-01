# disqus

Package disqus is an Ambient plugin that provides Disqus commenting.

**Import:** github.com/ambientkit/plugin/generic/disqus

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (6):

- **Name**: plugin.setting:read
  - **Description**: Access to the Disqus ID.
- **Name**: site.url:read
  - **Description**: Access to read the site URL.
- **Name**: site.scheme:read
  - **Description**: Access to read the site scheme.
- **Name**: site.funcmap:write
  - **Description**: Access to create global FuncMaps for templates.
- **Name**: site.asset:write
  - **Description**: Access to write meta tags to the header and add a nav and footer.
- **Name**: router.route:write
  - **Description**: Access to create routes for serving javascript and stylesheets.

## Settings

The plugin has the follow settings (1):

- **Name**: Disqus ID
  - **Type**: input
  - **Hidden**: false

## Routes

The plugin does not have any routes.

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin has the follow FuncMap items (1):

  - {{disqus_PageURL}}

## Assets

The plugin does not inject any assets.

## Embedded Files

The plugin does not have any embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/disqus"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/storage/memorystorage"
)

func main() {
	plugins := &ambient.PluginLoader{
		// Core plugins are implicitly trusted.
		Router:         nil,
		TemplateEngine: nil,
		SessionManager: nil,
		// Trusted plugins are those that are typically needed to boot so they
		// will be enabled and given full access.
		TrustedPlugins: map[string]bool{},
		Plugins: []ambient.Plugin{
			disqus.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
		},
	}
	_, _, err := ambientapp.NewApp("myapp", "1.0",
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

---

Docgen by [Ambient](https://ambientkit.github.io)
# sitemap

Package sitemap is an Ambient plugin that provides a sitemap.

**Import:** github.com/ambientkit/plugin/generic/sitemap

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (4):

- **Name**: site.url:read
  - **Description**: Access to read the site URL.
- **Name**: site.scheme:read
  - **Description**: Access to read the site scheme.
- **Name**: site.updated:read
  - **Description**: Access to read the last updated date.
- **Name**: site.post:read
  - **Description**: Access to read all the posts.

## Settings

The plugin does not have any settings.

## Routes

The plugin has the following routes (1):
  - **Method:** GET | **Path:** /sitemap.xml

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

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
	"github.com/ambientkit/plugin/generic/sitemap"
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
			sitemap.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes top to bottom.
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
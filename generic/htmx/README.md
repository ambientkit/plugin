# htmx

Package htmx is an Ambient plugin that adds the htmx JavaScript library to all pages: https://htmx.org/.

**Import:** github.com/ambientkit/plugin/generic/htmx

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (2):

- **Name**: plugin.setting:read
  - **Description**: Access to read the version setting.
- **Name**: site.asset:write
  - **Description**: Access to add the htmx JavaScript tag to every page.

## Settings

The plugin has the follow settings (1):

- **Name**: Version
  - **Type**: input
  - **Hidden**: false
  - **Default**: 1.7.0

## Routes

The plugin does not have any routes.

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (1):

  - **Type:** javascript
    - **Location:** head
    - **External:** true
    - **Path:** https://unpkg.com/htmx.org@1.7.0

## Embedded Files

The plugin does not have any embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/htmx"
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
			htmx.New(),
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
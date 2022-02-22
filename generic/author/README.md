# author

Package author is an Ambient plugin that sets an author meta tag in the HTML header.

**Import:** github.com/ambientkit/plugin/generic/author

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
  - **Description**: Access to the author name.
- **Name**: site.asset:write
  - **Description**: Access to write a meta tag to the header.

## Settings

The plugin has the follow settings (1):

- **Name**: Author
  - **Type**: input
  - **Hidden**: false

## Routes

The plugin does not have any routes.

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
	"github.com/ambientkit/plugin/generic/author"
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
			author.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
		},
	}
	_, _, err := ambient.NewApp("myapp", "1.0",
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

Docgen by [Ambient](https://ambientkit.github.io/)
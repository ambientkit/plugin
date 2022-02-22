# charset

Package charset is an Ambient plugin that sets a charset meta tag in the HTML header.

**Import:** github.com/ambientkit/plugin/generic/charset

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
  - **Description**: Access to the charset.
- **Name**: site.asset:write
  - **Description**: Access to write the charset meta tag.

## Settings

The plugin has the follow settings (1):

- **Name**: Charset
  - **Type**: input
  - **Hidden**: false
  - **Default**: utf-8

## Routes

The plugin does not have any routes.

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (1):

  - **Type:** generic
    - **Location:** head
    - **Tag Name:** meta
    - **Attributes (1):** 
      - **Name:** charset | **Value:** utf-8

## Embedded Files

The plugin does not have any embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/generic/charset"
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
			charset.New(),
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
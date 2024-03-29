# styles

Package styles is an Ambient plugin that provides a page to edit styles.

**Import:** github.com/ambientkit/plugin/generic/styles

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (3):

- **Name**: plugin.setting:read
  - **Description**: Access to read the plugin settings.
- **Name**: site.asset:write
  - **Description**: Access to add favicon.
- **Name**: router.route:write
  - **Description**: Access to add a route for styles.

## Settings

The plugin has the follow settings (2):

- **Name**: Favicon
  - **Type**: input
  - **Description**: Emoji cheatsheet
    - **URL**: https://github.com/ikatyang/emoji-cheat-sheet/blob/master/README.md
  - **Hidden**: false
- **Name**: Styles
  - **Type**: textarea
  - **Description**: No-class css themes. You can also paste a link like this: @import &#39;https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/dark.css&#39;
    - **URL**: https://www.cssbed.com/
  - **Hidden**: false

## Routes

The plugin has the following routes (1):
  - **Method:** GET | **Path:** /plugins/styles/css/style.css

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
	"github.com/ambientkit/plugin/generic/styles"
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
			styles.New(),
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
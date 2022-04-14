# jquery

Package jquery is an Ambient plugin that adds the jQuery library to all pages: https://jquery.com/.

**Import:** github.com/ambientkit/plugin/generic/jquery

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
  - **Description**: Access to add the jquery JavaScript tag to every page.

## Settings

The plugin has the follow settings (1):

- **Name**: Version
  - **Type**: input
  - **Description**: Version cannot be left blank. Ex: 3.6.0
    - **URL**: https://releases.jquery.com/jquery/
  - **Hidden**: false
  - **Default**: 3.6.0

## Routes

The plugin does not have any routes.

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (1):

  - **Type:** javascript
    - **Location:** body
    - **External:** true
    - **Path:** https://code.jquery.com/jquery-3.6.0.min.js
    - **Attributes (1):** 
      - **Name:** crossorigin | **Value:** anonymous

## Embedded Files

The plugin does not have any embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/jquery"
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
			jquery.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes top to bottom.
		},
	}
	_, _, err := ambientapp.NewApp(context.Background(), "myapp", "1.0",
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
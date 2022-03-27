# bootstrap

Package bootstrap is an Ambient plugin that adds the Bootstrap library to all pages: https://getbootstrap.com/.

**Import:** github.com/ambientkit/plugin/generic/bootstrap

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
  - **Description**: Access to add the Bootstrap CSS and JavaScript tags to every page.

## Settings

The plugin has the follow settings (1):

- **Name**: Version
  - **Type**: input
  - **Description**: Version cannot be left blank. Ex: 5.1.3
    - **URL**: https://github.com/twbs/bootstrap/releases
  - **Hidden**: false
  - **Default**: 5.1.3

## Routes

The plugin does not have any routes.

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (2):

  - **Type:** stylesheet
    - **Location:** head
    - **External:** true
    - **Path:** https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css
    - **Attributes (1):** 
      - **Name:** crossorigin | **Value:** anonymous
  - **Type:** javascript
    - **Location:** body
    - **External:** true
    - **Path:** https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js
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
	"github.com/ambientkit/plugin/generic/bootstrap"
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
			bootstrap.New(),
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
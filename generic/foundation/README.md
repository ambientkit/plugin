# foundation

Package foundation is an Ambient plugin that adds the Foundation library to all pages: https://get.foundation/. It requires jQuery.

**Import:** github.com/ambientkit/plugin/generic/foundation

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
  - **Description**: Access to add the Foundation CSS and JavaScript tags to every page.

## Settings

The plugin has the follow settings (1):

- **Name**: Version
  - **Type**: input
  - **Description**: Version cannot be left blank. Ex: 6.7.4
    - **URL**: https://github.com/foundation/foundation-sites/releases
  - **Hidden**: false
  - **Default**: 6.7.4

## Routes

The plugin does not have any routes.

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (3):

  - **Type:** stylesheet
    - **Location:** head
    - **External:** true
    - **Path:** https://cdn.jsdelivr.net/npm/foundation-sites@6.7.4/dist/css/foundation.min.css
    - **Attributes (1):** 
      - **Name:** crossorigin | **Value:** anonymous
  - **Type:** javascript
    - **Location:** body
    - **External:** true
    - **Path:** https://cdn.jsdelivr.net/npm/foundation-sites@6.7.4/dist/js/foundation.min.js
    - **Attributes (1):** 
      - **Name:** crossorigin | **Value:** anonymous
  - **Type:** javascript
    - **Location:** body
    - **Inline:** true
    - **Has Content:** true

## Embedded Files

The plugin does not have any embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/foundation"
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
			foundation.New(),
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
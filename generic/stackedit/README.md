# stackedit

Package stackedit is an Ambient plugin that provides a markdown editor using StackEdit.

**Import:** github.com/ambientkit/plugin/generic/stackedit

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

- **Name**: site.asset:write
  - **Description**: Access to add StackEdit JavaScript on pages.
- **Name**: router.route:write
  - **Description**: Access to add StackEdit JavaScript to URL.

## Settings

The plugin does not have any settings.

## Routes

The plugin has the following routes (1):
  - **Method:** GET | **Path:** /plugins/stackedit/js/stackedit.js

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (2):

  - **Type:** javascript
    - **Location:** body
    - **Auth Type:** authenticated
    - **External:** true
    - **Path:** https://unpkg.com/stackedit-js@1.0.7/docs/lib/stackedit.min.js
  - **Type:** javascript
    - **Location:** body
    - **Auth Type:** authenticated
    - **Path:** js/stackedit.js

## Embedded Files

The plugin has embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/stackedit"
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
			stackedit.New(),
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
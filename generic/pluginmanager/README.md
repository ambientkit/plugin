# pluginmanager

Package pluginmanager is an Ambient plugin that provides a plugin management system.

**Import:** github.com/ambientkit/plugin/generic/pluginmanager

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (12):

- **Name**: site.plugin:read
  - **Description**: Access to read the plugins.
- **Name**: site.plugin:enable
  - **Description**: Access to enable plugins.
- **Name**: site.plugin:disable
  - **Description**: Access to disable plugins.
- **Name**: site.plugin:delete
  - **Description**: Access to delete plugin storage.
- **Name**: site.funcmap:write
  - **Description**: Access add FuncMap for template helpers.
- **Name**: plugin.neighborsetting:read
  - **Description**: Access to read other plugin settings.
- **Name**: plugin.neighborsetting:write
  - **Description**: Access to write to other plugin settings
- **Name**: plugin.neighborgrant:read
  - **Description**: Access to read grant requests for plugins
- **Name**: plugin.neighborgrant:write
  - **Description**: Access to approve grants for plugins.
- **Name**: plugin.neighborroute:read
  - **Description**: Access to read routes for plugins.
- **Name**: router.route:write
  - **Description**: Access to create routes for editing the plugins.
- **Name**: plugin.trusted:read
  - **Description**: Access to read if a plugin is trusted or not.

## Settings

The plugin does not have any settings.

## Routes

The plugin has the following routes (8):
  - **Method:** GET | **Path:** /dashboard/plugins
  - **Method:** POST | **Path:** /dashboard/plugins
  - **Method:** GET | **Path:** /dashboard/plugins/{id}/delete
  - **Method:** GET | **Path:** /dashboard/plugins/{id}/settings
  - **Method:** POST | **Path:** /dashboard/plugins/{id}/settings
  - **Method:** GET | **Path:** /dashboard/plugins/{id}/grants
  - **Method:** POST | **Path:** /dashboard/plugins/{id}/grants
  - **Method:** GET | **Path:** /dashboard/plugins/{id}/routes

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin has the follow FuncMap items (1):

  - {{pluginmanager_URLHasParam}}

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
	"github.com/ambientkit/plugin/generic/pluginmanager"
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
			pluginmanager.New(),
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
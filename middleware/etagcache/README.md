# etagcache

Package etagcache is an Ambient plugin that provides caching using etag.

**Import:** github.com/ambientkit/plugin/middleware/etagcache

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the follow core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (2):

- **Name**: router.middleware:write
  - **Description**: Access to read and write ETag headers on responses.
- **Name**: plugin.setting:read
  - **Description**: Access to read MaxAge setting.

## Settings

The plugin has the follow settings (1):

- **Name**: MaxAge
  - **Type**: input
  - **Description**: MaxAge in seconds before Etag is checked. 30 days is 2592000.
  - **Hidden**: false

## Routes

The plugin does not have any routes.

## Middleware

The plugin has middleware (1).

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
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/middleware/etagcache"
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
		Plugins:        []ambient.Plugin{},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
			etagcache.New(),
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
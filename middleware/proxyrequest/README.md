# proxyrequest

Package proxyrequest is an Ambient plugin with middleware that proxies requests.

**Import:** github.com/ambientkit/plugin/middleware/proxyrequest

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (1):

- **Name**: router.middleware:write
  - **Description**: Access to proxy requests based on request URL.

## Settings

The plugin does not have any settings.

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
	"net/url"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/middleware/proxyrequest"
	"github.com/ambientkit/plugin/storage/memorystorage"
)

func main() {
	URL, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatalln(err.Error())
	}

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
			// Middleware - executes top to bottom.
			proxyrequest.New(URL, "/api"),
		},
	}
	_, _, err = ambientapp.NewApp("myapp", "1.0",
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
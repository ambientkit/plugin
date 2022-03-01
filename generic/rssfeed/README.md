# rssfeed

Package rssfeed is an Ambient plugin that provides an RSS feed.

**Import:** github.com/ambientkit/plugin/generic/rssfeed

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the following core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (6):

- **Name**: site.title:read
  - **Description**: Access to read the site title.
- **Name**: site.scheme:read
  - **Description**: Access to read the site scheme.
- **Name**: site.url:read
  - **Description**: Access to read the site URL.
- **Name**: site.post:read
  - **Description**: Access to read all the site posts.
- **Name**: plugin.setting:read
  - **Description**: Access to read the plugin settings.
- **Name**: plugin.setting:write
  - **Description**: Access to write to the plugin settings.

## Settings

The plugin has the follow settings (2):

- **Name**: Feed URL
  - **Type**: input
  - **Description**: Must start with a slash like this: /rss.xml
  - **Hidden**: false
  - **Default**: /rss.xml
- **Name**: Description
  - **Type**: textarea
  - **Hidden**: false

## Routes

The plugin has the following routes (1):
  - **Method:** GET | **Path:** /rss.xml

## Middleware

The plugin does not have any middleware.

## FuncMap

The plugin does not have a FuncMap.

## Assets

The plugin injects the following assets (1):

  - **Type:** generic
    - **Location:** head
    - **Tag Name:** link
    - **Attributes (4):** 
      - **Name:** rel | **Value:** alternative
      - **Name:** href | **Value:** /rss.xml
      - **Name:** application | **Value:** rss&#43;xml
      - **Name:** title | **Value:** 

## Embedded Files

The plugin does not have any embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/rssfeed"
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
			rssfeed.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
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
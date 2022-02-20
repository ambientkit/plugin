# prism

Package prism is an Ambient plugin that provides syntax highlighting using Prism (https://prismjs.com/).

**Import:** github.com/ambientkit/plugin/pkg/docgen/testdata/prism

**Version:** 1.0.0

## Plugin Type

The plugin can be used as the follow core types:

- **Logger:** false
- **Storage System:** false
- **Router:** false
- **Template Engine:** false
- **Session Manager:** false

## Grants

The plugin request the following grants (3):

- **Name**: site.asset:write
  - **Description**: Access to add stylesheets and javascript to each page.
- **Name**: router.route:write
  - **Description**: Access to create routes for accessing stylesheets.
- **Name**: plugin.setting:read
  - **Description**: Read own plugin settings.

## Settings

The plugin has the follow settings (2):

- **Name**: Version
  - **Type**: input
  - **Description**: View releases (ex: 1.23.0)
    - **URL**: https://github.com/PrismJS/prism/releases
  - **Hidden**: false
  - **Default**: 1.23.0
- **Name**: Styles
  - **Type**: textarea
  - **Description**: You can paste a theme from https://github.com/PrismJS/prism-themes/tree/master/themes or an import like this using https://gitcdn.link/: @import &#39;https://gitcdn.link/cdn/PrismJS/prism-themes/d00360c3b3cfe495f45cc06865969c7731a94763/themes/prism-vsc-dark-plus.css&#39;
    - **URL**: https://github.com/PrismJS/prism-themes/tree/master/themes
  - **Hidden**: false

## Routes

The plugin has the following routes (2):
  - **Method:** GET | **Path:** /plugins/prism/css/style.css
  - **Method:** GET | **Path:** /plugins/prism/css/clean.css

## Middleware

The plugin has middleware (1).

## FuncMap

The plugin has the follow FuncMap items (1):

  - {{prism_PageURL}}

## Assets

The plugin injects the following assets (6):

  - **Type:** stylesheet
    - **Location:** head
    - **Path:** css/clean.css
  - **Type:** javascript
    - **Location:** body
    - **External:** true
    - **Path:** https://unpkg.com/prismjs@1.23.0/components/prism-core.min.js
  - **Type:** javascript
    - **Location:** body
    - **External:** true
    - **Path:** https://unpkg.com/prismjs@1.23.0/plugins/autoloader/prism-autoloader.min.js
  - **Type:** stylesheet
    - **Location:** head
    - **Path:** css/disqus.css
    - **Has Content:** true
    - **Layout Only:** post
  - **Type:** javascript
    - **Location:** body
    - **Inline:** true
    - **Path:** js/disqus.js
    - **Layout Only:** post
    - **Replace (2):**
      - **Find:** {{.DisqusID}} | **Replace:** 123
      - **Find:** {{.SiteURL}} | **Replace:** 456
  - **Type:** generic
    - **Location:** main
    - **Tag Name:** div
    - **Closing Tag:** true
    - **Layout Only:** post
    - **Attributes (1):** 
      - **Name:** id | **Value:** disqus_thread

## Embedded Files

The plugin has embedded files.

## Example Usage

```go
package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/docgen/testdata/prism"
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
			prism.New(),
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
package docgen

import (
	"log"
	"net/http"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/uuid"
	"github.com/ambientkit/plugin/router/routerecorder"
	"github.com/ambientkit/plugin/sessionmanager/scssession"
	"github.com/ambientkit/plugin/storage/memorystorage"
	"github.com/ambientkit/plugin/templateengine/htmlengine"
)

// App contains all the plugins.
type App struct {
	Handler http.Handler
	Mux     *routerecorder.Plugin
}

// LighweightAppSetup setups up a lighweight Ambient system.
func LighweightAppSetup(appName string, p ambient.Plugin, trust bool) *App {
	trusted := map[string]bool{}
	if trust {
		// Automatically trust the plugin.
		trusted[p.PluginName()] = true
	}

	pluginList := []ambient.Plugin{
		p,
	}

	defaultSessionManager := scssession.New(uuid.EncodedString(32))
	var sessionManager ambient.SessionManagerPlugin

	// Determine if the plugin satisfies the session manager interface.
	sm, ok := p.(ambient.SessionManagerPlugin)
	if ok {
		sessionManager = sm
		// Remove plugin from plugin list.
		pluginList = []ambient.Plugin{}
	} else {
		sessionManager = defaultSessionManager
	}

	var middleware ambient.MiddlewarePlugin

	// Determine if the plugin satisfies the middleware interface.
	mw, ok := p.(ambient.MiddlewarePlugin)
	if ok {
		middleware = mw
		// Remove plugin from plugin list.
		pluginList = []ambient.Plugin{}
	} else {
		middleware = defaultSessionManager
	}

	mux := routerecorder.New()
	plugins := &ambient.PluginLoader{
		// Core plugins are implicitly trusted.
		Router:         mux,
		TemplateEngine: htmlengine.New(),
		SessionManager: sessionManager,
		// Trusted plugins are those that are typically needed to boot so they
		// will be enabled and given full access.
		TrustedPlugins: trusted,
		Plugins:        pluginList,
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
			middleware, // Session manager middleware.
		},
	}
	ambientApp, logger, err := ambientapp.NewApp(appName, "1.0",
		zaplogger.New(),
		ambient.StoragePluginGroup{
			Storage: memorystorage.New(),
		},
		plugins)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Load the plugins and return the handler.
	handler, err := ambientApp.Handler()
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &App{
		Handler: handler,
		Mux:     mux,
	}
}

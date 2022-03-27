package grpctestutil

import (
	"net/http"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/debugpprof"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/grpctestutil/testingdata/plugin/neighbor"
	trustPlugin "github.com/ambientkit/plugin/pkg/grpctestutil/testingdata/plugin/trust"
	"github.com/ambientkit/plugin/router/awayrouter"
	"github.com/ambientkit/plugin/sessionmanager/scssession"
	"github.com/ambientkit/plugin/storage/localstorage"
	"github.com/ambientkit/plugin/storage/memorystorage"
	"github.com/ambientkit/plugin/templateengine/htmlengine"
)

// Setup sets up a test gRPC server.
func Setup(trust bool) (*ambientapp.App, error) {
	h := func(log ambient.Logger, renderer ambient.Renderer, w http.ResponseWriter, r *http.Request, err error) {
		if err != nil {
			switch e := err.(type) {
			case ambient.Error:
				errText := e.Error()
				if len(errText) == 0 {
					errText = http.StatusText(e.Status())
				}
				http.Error(w, errText, e.Status())
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}
	}

	trusted := make(map[string]bool)
	trusted["trust"] = true
	if trust {
		trusted["hello"] = true
	}

	sessPlugin := scssession.New("5ba3ad678ee1fd9c4fddcef0d45454904422479ed762b3b0ddc990e743cb65e0")
	plugins := &ambient.PluginLoader{
		// Core plugins are implicitly trusted.
		Router:         awayrouter.New(h),
		TemplateEngine: htmlengine.New(),
		SessionManager: sessPlugin,
		// Trusted plugins are those that are typically needed to boot so they
		// will be enabled and given full access.
		TrustedPlugins: trusted,
		Plugins: []ambient.Plugin{
			neighbor.New(),
			trustPlugin.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes top to bottom.
			sessPlugin,
			ambient.NewGRPCPlugin("hello", "./pkg/grpctestutil/testingdata/plugin/hello/cmd/plugin/ambplugin"),
		},
	}
	app, _, err := ambientapp.NewApp("myapp", "1.0",
		zaplogger.New(),
		ambient.StoragePluginGroup{
			Storage: memorystorage.New(),
		},
		plugins)
	return app, err
}

// Setup2 sets up a test gRPC server.
func Setup2(trust bool) (*ambientapp.App, error) {
	// h := func(log ambient.Logger, renderer ambient.Renderer, w http.ResponseWriter, r *http.Request, err error) {
	// 	if err != nil {
	// 		switch e := err.(type) {
	// 		case ambient.Error:
	// 			errText := e.Error()
	// 			if len(errText) == 0 {
	// 				errText = http.StatusText(e.Status())
	// 			}
	// 			http.Error(w, errText, e.Status())
	// 		default:
	// 			http.Error(w, http.StatusText(http.StatusInternalServerError),
	// 				http.StatusInternalServerError)
	// 		}
	// 	}
	// }

	trusted := make(map[string]bool)
	//trusted["trust"] = true
	if trust {
		//trusted["hello"] = true
		trusted["bearblog"] = true
		trusted["bearcss"] = true
		trusted["pluginmanager"] = true
		//trusted["debugpprof"] = true
	}

	sessPlugin := scssession.New("5ba3ad678ee1fd9c4fddcef0d45454904422479ed762b3b0ddc990e743cb65e0")
	plugins := &ambient.PluginLoader{
		// Core plugins are implicitly trusted.
		Router:         awayrouter.New(nil),
		TemplateEngine: htmlengine.New(),
		SessionManager: sessPlugin,
		// Trusted plugins are those that are typically needed to boot so they
		// will be enabled and given full access.
		TrustedPlugins: trusted,
		Plugins: []ambient.Plugin{
			//neighbor.New(),
			//trustPlugin.New(),
			//bearblog.New(os.Getenv("AMB_PASSWORD_HASH")),
			ambient.NewGRPCPlugin("bearcss", "./generic/bearcss/cmd/plugin/ambplugin"),
			ambient.NewGRPCPlugin("pluginmanager", "./generic/pluginmanager/cmd/plugin/ambplugin"),
			debugpprof.New(),
			//bearcss.New(),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes top to bottom.
			sessPlugin,
			//ambient.NewGRPCPlugin("hello", "./pkg/grpcp/testingdata/plugin/hello/cmd/plugin/hello"),
			//bearblog.New(os.Getenv("AMB_PASSWORD_HASH")),
			ambient.NewGRPCPlugin("bearblog", "./generic/bearblog/cmd/plugin/ambplugin"),
		},
	}
	app, _, err := ambientapp.NewApp("myapp", "1.0",
		zaplogger.New(),
		ambient.StoragePluginGroup{
			//Storage: memorystorage.New(),
			Storage: localstorage.New("teststorage/site.json", "teststorage/session.bin"),
		},
		plugins)

	// app.SetLogLevel(ambient.LogLevelDebug)

	return app, err
}

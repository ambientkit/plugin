package zaplogger_test

import (
	"bufio"
	"log"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/docgen"
	"github.com/ambientkit/plugin/pkg/loggertestsuite"
	"github.com/ambientkit/plugin/storage/memorystorage"
)

// Run the standard logger test suite.
func TestMain(t *testing.T) {
	rt := loggertestsuite.New()
	rt.Run(t, func(writer *bufio.Writer) ambient.AppLogger {
		return zaplogger.NewLogger("app", "1.0", writer)
	})
}

func ExampleNew() {
	plugins := &ambient.PluginLoader{
		// Core plugins are implicitly trusted.
		Router:         nil,
		TemplateEngine: nil,
		SessionManager: nil,
		// Trusted plugins are those that are typically needed to boot so they
		// will be enabled and given full access.
		TrustedPlugins: map[string]bool{},
		Plugins:        []ambient.Plugin{},
		Middleware:     []ambient.MiddlewarePlugin{
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

func TestGenerateDocs(t *testing.T) {
	docgen.Generate(t, zaplogger.New(), "")
}

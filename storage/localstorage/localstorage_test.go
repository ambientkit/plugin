package localstorage_test

import (
	"context"
	"log"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/docgen"
	"github.com/ambientkit/plugin/storage/localstorage"
)

func ExampleNew() {
	var (
		// StorageSitePath is the location of the site file.
		StorageSitePath = "storage/site.bin"
		// StorageSessionPath is the location of the session file.
		StorageSessionPath = "storage/session.bin"
	)

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
			// Middleware - executes top to bottom.
		},
	}
	_, _, err := ambientapp.NewApp(context.Background(), "myapp", "1.0",
		zaplogger.New(),
		ambient.StoragePluginGroup{
			Storage: localstorage.New(StorageSitePath, StorageSessionPath),
		},
		plugins)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func TestGenerateDocs(t *testing.T) {
	docgen.Generate(t, localstorage.New("test1.bin", "test2.bin"), "")
}

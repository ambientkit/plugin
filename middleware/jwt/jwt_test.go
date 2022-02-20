package jwt_test

import (
	"log"
	"testing"
	"time"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/middleware/jwt"
	"github.com/ambientkit/plugin/pkg/docgen"
	"github.com/ambientkit/plugin/pkg/uuid"
	"github.com/ambientkit/plugin/storage/memorystorage"
)

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
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes bottom to top.
			jwt.New([]byte(uuid.EncodedString(32)), time.Hour*1, []string{}),
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

func TestGenerateDocs(t *testing.T) {
	docgen.Generate(t, jwt.New([]byte(uuid.EncodedString(32)), time.Hour*1, []string{}), "")
}

package rove_test

import (
	"log"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/generic/rove"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/docgen"
	"github.com/ambientkit/plugin/storage/memorystorage"
	"github.com/stretchr/testify/assert"
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
		Plugins: []ambient.Plugin{
			rove.New(nil),
		},
		Middleware: []ambient.MiddlewarePlugin{
			// Middleware - executes top to bottom.
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
	docgen.Generate(t, rove.New(nil), "")
}

func TestMain(t *testing.T) {
	var err error
	// docker run --name=mysql57 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:5.7
	// docker rm mysql57 -f
	// os.Setenv("DB_USERNAME", "root")
	// os.Setenv("DB_PASSWORD", "password")
	// os.Setenv("DB_HOSTNAME", "localhost")
	// os.Setenv("DB_PORT", "3306")
	// os.Setenv("DB_NAME", "main")
	// p := rove.New(nil)
	// err = p.Enable(nil)
	assert.NoError(t, err)
}
